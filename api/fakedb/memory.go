package fakedb

import "sync"

type Memory interface {
	Set(key, val interface{})
	Get(key interface{}) (item interface{})
	Has(key interface{}) (has bool)
	Len() (l int)
	Del(key interface{})
	ForEach(fn func(item interface{}) bool)
}

type memory struct {
	data map[interface{}]interface{}
	lk   *sync.Mutex
}

func New() Memory {
	return &memory{
		data: make(map[interface{}]interface{}),
		lk:   &sync.Mutex{},
	}
}

func (m *memory) onLock(rw func()) {
	m.lk.Lock()
	rw()
	m.lk.Unlock()
}

func (m *memory) Set(key, val interface{}) {
	m.onLock(func() {
		m.data[key] = val
	})
}

func (m *memory) Get(key interface{}) (item interface{}) {
	m.onLock(func() {
		item = m.data[key]
	})
	return
}

func (m *memory) Has(key interface{}) (has bool) {
	m.onLock(func() {
		_, has = m.data[key]
	})
	return
}

func (m *memory) Del(key interface{}) {
	m.onLock(func() {
		delete(m.data, key)
	})
}

func (m *memory) Len() (l int) {
	m.onLock(func() {
		l = len(m.data)
	})
	return l
}

func (m *memory) ForEach(fn func(item interface{}) bool) {
	m.onLock(func() {
		for k := range m.data {
			if !fn(m.data[k]) {
				break
			}
		}
	})
}
