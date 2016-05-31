package ssmap

import "sync"

type SimpleSafeMap struct {
	m map[interface{}]interface{}	
	l *sync.RWMutex
}

func (ssm *SimpleSafeMap) ROnly(cbfunc func(map[interface{}]interface{})) {
	ssm.l.RLock()
	defer ssm.l.RUnlock()
	cbfunc(ssm.m)
}

func (ssm *SimpleSafeMap) RW(cbfunc func(map[interface{}]interface{})) {
	ssm.l.Lock()
	defer ssm.l.Unlock()
	cbfunc(ssm.m)
}

func NewSafeMap() *SimpleSafeMap {
	return &SimpleSafeMap {
		m : make(map[interface{}]interface{}),
		l : new(sync.RWMutex),
	}
}
