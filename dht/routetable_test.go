package dht

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAddNode(t *testing.T) {
	lid := generateNodeId("local.test")
	rtable := NewRouteTable(NewNodeInfoFromHost(lid, "127.0.0.1:9090"), 2, 160)

	hnode := rtable.hnode
	id := generateNodeId("test1")
	node1 := NewNodeInfoFromHost(id, "127.0.0.1:9000")
	node2 := NewNodeInfoFromHost(generateNodeId("test2"), "127.0.0.1:9001")
	node3 := NewNodeInfoFromHost(generateNodeId("test3"), "127.0.0.1:9002")
	assert.Equal(t, nil, rtable.AddNode(node1))
	assert.Equal(t, nil, rtable.AddNode(node2))
	assert.Equal(t, nil, rtable.AddNode(node3))
	assert.Equal(t, 4, rtable.size)
	assert.Equal(t, node1, rtable.Find(id))
	assert.Equal(t, []*NodeInfo{hnode, node1, node2, node3}, rtable.FindClosest(id, 4))
}

func TestSplitNode(t *testing.T) {
	lid := generateNodeId("local.test")
	hnode := NewNodeInfoFromHost(lid, "127.0.0.1:9090")
	rtable := NewRouteTable(hnode, 2, 10)

	for i := 9000; i < 10000; i++ {
		host := fmt.Sprintf("127.0.0.1:%d", i)
		rtable.AddNode(NewNodeInfoFromHost(generateNodeId(host), host))
		//assert.Equal(t, nil, rtable.AddNode(NewNodeInfoFromHost(generateNodeId(name), host)))
	}
	t.Logf("routr table size:%d, cap:%d", rtable.size, rtable.lastBucketIndex)
	assert.Equal(t, true, rtable.nodeMap[rtable.lastBucketIndex].Exist(hnode))
	assert.Equal(t, rtable.size, len(rtable.FindClosest(hnode.Id, 60)))
}
