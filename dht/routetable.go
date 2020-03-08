package dht

import (
	"errors"
	"github.com/GalaIO/P2Pcrawler/misc"
	"math/bits"
	"strings"
	"sync"
)

var rtableLogger = misc.GetLogger().SetPrefix("rtable")

type Bucket struct {
	sync.RWMutex
	nodes []*NodeInfo
	size  int
	cap   int
}

func NewBucket(capacity int) *Bucket {
	return &Bucket{
		size:  0,
		cap:   capacity,
		nodes: make([]*NodeInfo, capacity),
	}
}

func (b *Bucket) Exist(node *NodeInfo) bool {
	defer b.RUnlock()
	b.RLock()
	for i := 0; i < b.size; i++ {
		if b.nodes[i].Equals(node) {
			return true
		}
	}
	return false
}

func (b *Bucket) Iterator(handler func(*NodeInfo)) {
	defer b.RUnlock()
	b.RLock()
	for _, v := range b.nodes {
		handler(v)
	}
}

func (b *Bucket) Add(node *NodeInfo) bool {
	if b.Full() {
		return false
	}
	defer b.Unlock()
	b.Lock()
	if b.full() {
		return false
	}
	b.nodes[b.size] = node
	b.size++
	return true
}

func (b *Bucket) All() []*NodeInfo {
	defer b.RUnlock()
	b.RLock()
	return b.nodes[:b.size]
}

func (b *Bucket) Size() int {
	defer b.RUnlock()
	b.RLock()
	return b.size
}

func (b *Bucket) Full() bool {
	defer b.RUnlock()
	b.RLock()
	return b.full()
}

func (b *Bucket) full() bool {
	if b.size >= b.cap {
		return true
	}
	return false
}

type RouteTable struct {
	sync.RWMutex
	hnode           *NodeInfo
	k               int
	nodeMap         map[int]*Bucket
	size            int
	capacity        int
	lastBucketIndex int
}

func NewRouteTable(hostNode *NodeInfo, k, capacity int) *RouteTable {
	rtable := &RouteTable{
		hnode:           hostNode,
		k:               k,
		nodeMap:         make(map[int]*Bucket, 16),
		size:            1,
		capacity:        capacity,
		lastBucketIndex: 0,
	}
	bucket := NewBucket(k)
	bucket.Add(hostNode)
	rtable.nodeMap[0] = bucket
	rtableLogger.Info("init route table", misc.Dict{"k": k})
	return rtable
}

func (t *RouteTable) Find(id string) *NodeInfo {
	defer t.RUnlock()
	t.RLock()
	preLen := t.solveMaxPreLen(id)
	bucket := t.nodeMap[preLen]
	nodes := bucket.All()
	for _, n := range nodes {
		if strings.EqualFold(n.Id, id) {
			return n
		}
	}
	return nil
}

func (t *RouteTable) FindClosest(id string, k int) []*NodeInfo {
	defer t.RUnlock()
	t.RLock()
	pLen := t.solveMaxPreLen(id)
	nodes := make([]*NodeInfo, 0, k)
	for len(nodes) <= k && pLen >= 0 {
		bucket := t.nodeMap[pLen]
		nodes = append(nodes, bucket.All()...)
		pLen--
	}
	return nodes
}

func (t *RouteTable) solveMaxPreLen(id string) int {
	pLen := solvePrefixLen(id, t.hnode.Id)
	if pLen <= t.lastBucketIndex {
		return pLen
	}
	return t.lastBucketIndex
}

func (t *RouteTable) AddNode(node *NodeInfo) error {
	defer t.Unlock()
	t.Lock()
	pLen := t.solveMaxPreLen(node.Id)
	return t.addNode(node, pLen)
}

func (t *RouteTable) addNode(node *NodeInfo, preLen int) error {
	// hit last buket or other buchet
	bucket := t.nodeMap[preLen]
	if preLen >= t.lastBucketIndex && !bucket.Exist(node) && bucket.Full() {
		// hit last bucket, and split new bucket
		return t.addNodeWithsplitBucket(node)
	}

	if bucket.Full() {
		// full bucket just ignore
		rtableLogger.Info("the bucket has full", misc.Dict{"tableSize": t.size, "preLen": preLen})
		return errors.New("the bucket is full")
	}

	if !bucket.Exist(node) && !bucket.Add(node) {
		// full bucket just ignore
		rtableLogger.Info("the bucket has full", misc.Dict{"tableSize": t.size, "preLen": preLen})
		return errors.New("the bucket is full")
	}
	t.size++
	return nil
}

// split left and right bucket, and add node to left bucket
func (t *RouteTable) addNodeWithsplitBucket(node *NodeInfo) error {
	if t.lastBucketIndex >= t.capacity {
		rtableLogger.Error("the route table has full", misc.Dict{"size": t.size})
		return errors.New("full route table")
	}
	sourceBucket := t.nodeMap[t.lastBucketIndex]
	preSize := t.size
	t.size -= sourceBucket.Size()
	newLeftBucket := NewBucket(t.k)
	newRightBucket := NewBucket(t.k)
	t.nodeMap[t.lastBucketIndex] = newRightBucket
	t.lastBucketIndex++
	t.nodeMap[t.lastBucketIndex] = newLeftBucket

	rtableLogger.Trace("the route table has split new bucket", misc.Dict{"lastIndex": t.lastBucketIndex, "size": t.size})

	snodes := sourceBucket.All()
	for _, n := range snodes {
		t.addNode(n, t.solveMaxPreLen(n.Id))
	}
	t.addNode(node, t.solveMaxPreLen(node.Id))
	if preSize+1 != t.nodeCount() {
		rtableLogger.Error("the route table add fail", misc.Dict{"lastIndex": t.lastBucketIndex, "size": t.size})
	}
	return nil
}

func (t *RouteTable) nodeCount() int {
	size := 0
	for i := 0; i <= t.lastBucketIndex; i++ {
		size += t.nodeMap[i].Size()
	}
	return size
}

func (t *RouteTable) Size() int {
	defer t.RUnlock()
	t.RLock()
	return t.size
}

func (t *RouteTable) Full() bool {
	defer t.RUnlock()
	t.RLock()
	return t.lastBucketIndex >= t.capacity && t.nodeMap[t.lastBucketIndex].Full()
}

func solvePrefixLen(a string, b string) int {
	bytes := misc.Xor([]byte(a), []byte(b))
	for i, v := range bytes {
		if v != 0 {
			return i*8 + bits.LeadingZeros8(v)
		}
	}
	return len(bytes) * 8
}
