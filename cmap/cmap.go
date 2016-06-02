package cmap

import "sync"

type KeyValue struct {
	Key interface{}
	Value interface{}
}

type CMap struct {
	m map[interface{}]interface{}	
	l *sync.RWMutex
}

func (sm *CMap) Len() int {
	sm.l.RLock()
	defer sm.l.RUnlock()
	return len(sm.m)
}

func (sm *CMap) IsEmpty() bool {
	sm.l.RLock()
	defer sm.l.RUnlock()
	return len(sm.m) == 0
}

func (sm *CMap) Get(k interface{}) (interface{}, bool) {
	sm.l.RLock()
	defer sm.l.RUnlock()
	v, ok := sm.m[k] 
	return v, ok
}

func (sm *CMap) Set(k interface{}, v interface{}) {
	sm.l.Lock()
	defer sm.l.Unlock()
	sm.m[k] = v
}

func (sm *CMap) SetIfAbsent(k interface{}, v interface{}) bool {
	sm.l.Lock()
	defer sm.l.Unlock()
	_, ok := sm.m[k]
	if !ok {
		sm.m[k] = v
	}
	return !ok
}

func (sm *CMap) replace(k interface{}, v interface{}) bool {
	sm.l.Lock()
	defer sm.l.Unlock()
	_, ok := sm.m[k];
	if ok {
		sm.m[k] = v
	}
	return ok
}

func (sm *CMap) SetDefault(k interface{}, d interface{}) interface{} {
	sm.l.Lock()
	defer sm.l.Unlock()
	if v, ok := sm.m[k]; ok {
		 return v
	} 
	sm.m[k] = d
	return d
}

func (sm *CMap) Delete(k interface{}) {
	sm.l.Lock()
	defer sm.l.Unlock()
	delete(sm.m, k)
}


func (sm *CMap) IsExistsKey(k interface{}) bool {
	sm.l.RLock()
	defer sm.l.RUnlock()
	_, ok := sm.m[k]
	return ok
}

func (sm *CMap) IsExistsValue(v interface{}) bool {
	sm.l.RLock()
	defer sm.l.RUnlock()
	for _, iv := range sm.m {
		if iv == v {
			return true
		}
	}
	return false
}

func (sm *CMap) Clear() {
        sm.l.Lock()
        defer sm.l.Unlock()
	sm.m = make(map[interface{}]interface{})
}

func (sm *CMap) Merge(src *CMap) {
	src.rLock()
	sm.l.Lock()
	defer func() {
		src.rUnlock()
		sm.l.Unlock()
	}()
	for ik, iv := range src.m {
		if _, ok := sm.m[ik]; !ok {
			sm.m[ik] = iv
		}
	}
}

func (sm *CMap) Update(src *CMap) {
	src.rLock()
	sm.l.Lock()
	defer func() {
		src.rUnlock()
		sm.l.Unlock()
	}()
	for ik, iv := range src.m {
		sm.m[ik] = iv
	}
}

func (sm *CMap) Copy() *CMap { 
	nm := NewCMap()
	sm.l.RLock()
	defer sm.l.RUnlock()
	for ik, iv := range sm.m {
		nm.unsafeSet(ik, iv)
	}
	return nm
}
 
func (sm *CMap) Keys() []interface{} { 
	sm.l.RLock()
	defer sm.l.RUnlock()
	l := len(sm.m)
	ks := make([]interface{}, 0, l)
	for ik, _ := range sm.m {
		ks = append(ks, ik)
	}
	return ks
}

func (sm *CMap) Values() []interface{} { 
	sm.l.RLock()
	defer sm.l.RUnlock()
	l := len(sm.m)
	vs := make([]interface{}, 0, l)
	for _, iv := range sm.m {
		vs = append(vs, iv)
	}
	return vs
}

func (sm *CMap) Items() []KeyValue { 
	sm.l.RLock()
	defer sm.l.RUnlock()
	l := len(sm.m)
	kvs := make([]KeyValue, 0, l)
	for ik, iv := range sm.m {
		kv := KeyValue{
			Key: ik,
			Value: iv,
		}
		kvs = append(kvs, kv)
	}
	return kvs
}

func (sm *CMap) Pop(k interface{}) (interface{}, bool) { 
	sm.l.Lock()
	sm.l.Unlock()
	v, ok := sm.m[k]
	if !ok {
		return nil, ok
	}
	delete(sm.m, k)
	return v, ok
}
	
func (sm *CMap) PopItem() (interface{}, interface{}, bool) { 
	sm.l.Lock()
	defer sm.l.Unlock()
	var rk interface{} = nil
	var rv interface{} = nil
	var ok bool = false
	for ik, iv := range sm.m {
		rk = ik 
		rv = iv
		ok = true
		break
	}
	if ok {
		delete(sm.m, rk)
	}
	return rk, rv, ok
}

func (sm *CMap) ForeachKey(cbfunc func(k interface{})) {
	sm.l.Lock()
	defer sm.l.Unlock()
	for ik, _ := range sm.m {
		cbfunc(ik)
	}
}
 
func (sm *CMap) ForeachValue(cbfunc func(v interface{})) {
	sm.l.Lock()
	defer sm.l.Unlock()
	for _, iv := range sm.m {
		cbfunc(iv)
	}
} 

func (sm *CMap) ForeachItem(cbfunc func(k interface{}, v interface{})) {
	sm.l.Lock()
	defer sm.l.Unlock()
	for ik, iv := range sm.m {
		cbfunc(ik, iv)
	}
} 

func (sm *CMap) rLock() {
        sm.l.RLock()
}

func (sm *CMap) rUnlock() {
        sm.l.RUnlock()
}

func (sm *CMap) unsafeSet(k interface{}, v interface{}) {
        sm.m[k] = v 
}

func NewCMap() *CMap {
	return &CMap {
		m : make(map[interface{}]interface{}),
		l : new(sync.RWMutex),
	}
}
