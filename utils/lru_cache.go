package utils

import (
	"container/list"
	"errors"
	"sync"
)

type Lrucache struct {
	cap  int32
	dict map[string]*list.Element
	l    *list.List
	mu   *sync.RWMutex
}

type elem struct {
	k      string
	v      interface{}
	expire int64
}

func NewLrucache(size int32) *Lrucache {
	//if size < 0 {
	//	return nil, errors.New("size must be positive")
	//}
	lc := &Lrucache{
		cap:  size,
		l:    list.New(),
		dict: make(map[string]*list.Element),
		mu:   &sync.RWMutex{},
	}
	return lc
}

func (lc *Lrucache) Set(key string, value interface{}) {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	if v, ok := lc.dict[key]; ok {
		lc.l.MoveToFront(v)
		v.Value.(*elem).v = value
		v.Value.(*elem).expire = 0
		return
	}

	if int32(lc.l.Len()) >= lc.cap {
		lc.DeleteOldest()
	}

	e := &elem{
		k:      key,
		v:      value,
		expire: 0,
	}
	node := lc.l.PushFront(e)
	lc.dict[key] = node

	return
}

func (lc *Lrucache) Get(key string) (value interface{}, ok bool) {
	lc.mu.RLock()
	defer lc.mu.RUnlock()
	if v, ok := lc.dict[key]; ok {
		lc.l.MoveToFront(v)
		return v.Value.(*elem).v, true
	}
	return nil, false
}

func (lc *Lrucache) Delete(key string) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()
	if v, ok := lc.dict[key]; ok {
		lc.l.Remove(v)
		delete(lc.dict, key)
		return nil
	}
	return errors.New("not found")
}

func (lc *Lrucache) DeleteOldest() {
	oldest := lc.l.Back()
	if oldest != nil {
		oel := oldest.Value.(*elem)
		delete(lc.dict, oel.k)
		lc.l.Remove(oldest)
	}
}

func (lc *Lrucache) Keys() []string {
	lc.mu.RLock()
	defer lc.mu.RUnlock()
	var keys []string
	for k := range lc.dict {
		keys = append(keys, k)
	}
	return keys
}
