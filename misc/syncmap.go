package misc

import "sync"

type SyncMap struct {
	sync.Mutex
	mapping map[string]interface{}
}

func NewSyncMap(captity int) *SyncMap {
	return &SyncMap{mapping: make(map[string]interface{}, captity)}
}

func (m *SyncMap) Put(key string, handler interface{}) {
	defer m.Unlock()
	m.Lock()
	m.mapping[key] = handler
}

func (m *SyncMap) Get(key string) (interface{}, bool) {
	defer m.Unlock()
	m.Lock()
	handlerFunc, ok := m.mapping[key]
	return handlerFunc, ok
}

// double check delete
func (m *SyncMap) Del(txId string) {
	if _, ok := m.mapping[txId]; !ok {
		return
	}
	defer m.Unlock()
	m.Lock()
	if _, ok := m.mapping[txId]; !ok {
		return
	}
	delete(m.mapping, txId)
}
