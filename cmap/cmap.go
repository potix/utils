package cmap

import (
	"sync"
)

type KeyValue[K comparable, V any] struct {
	Key   K
	Value V
}

type CMap[K comparable, V any] struct {
	m map[K] V
	l sync.RWMutex
}

func (cm *CMap[K, V]) Len() (length int) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	return len(cm.m)
}

func (cm *CMap[K, V]) Get(key K) (value V, ok bool) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	value, ok = cm.m[key]
	return value, ok
}

func (cm *CMap[K, V]) Set(key K, value V) {
	cm.l.Lock()
	defer cm.l.Unlock()
	cm.m[key] = value
}

func (cm *CMap[K, V]) Update(key K, updateCb func(V) V) (ok bool) {
	cm.l.Lock()
	defer cm.l.Unlock()
	value, ok := cm.m[key]
	if ok {
		cm.m[key] = updateCb(value)
	}
	return ok
}

func (cm *CMap[K, V]) SetIfAbsent(key K, value V) (ok bool) {
	cm.l.Lock()
	defer cm.l.Unlock()
	_, ok = cm.m[key]
	if !ok {
		cm.m[key] = value
	}
	return !ok
}

func (cm *CMap[K, V]) SetIfExist(key K, value V) (ok bool) {
	cm.l.Lock()
	defer cm.l.Unlock()
	_, ok = cm.m[key]
	if ok {
		cm.m[key] = value
	}
	return ok
}

func (cm *CMap[K, V]) GetWithDefault(key K, defValue V) (value V) {
	cm.l.Lock()
	defer cm.l.Unlock()
	value, ok := cm.m[key]
	if !ok {
		cm.m[key] = defValue
		return defValue
	}
	return value
}

func (cm *CMap[K, V]) Delete(key K) (ok bool) {
	cm.l.Lock()
	defer cm.l.Unlock()
	_, ok = cm.m[key]
	if ok {
		delete(cm.m, key)
	}
	return ok
}

func (cm *CMap[K, V]) Clear() {
	cm.l.Lock()
	defer cm.l.Unlock()
	cm.m = make(map[K]V)
}

func (cm *CMap[K, V]) OverwriteMerge(src *CMap[K, V]) {
	src.l.RLock()
	cm.l.Lock()
	defer func() {
		src.l.RUnlock()
		cm.l.Unlock()
	}()
	for iKey, iValue := range src.m {
		cm.m[iKey] = iValue
	}
}

func (cm *CMap[K, V]) KeepMerge(src *CMap[K, V]) {
	src.l.RLock()
	cm.l.Lock()
	defer func() {
		cm.l.Unlock()
		src.l.RUnlock()
	}()
	for iKey, iValue := range src.m {
		if _, ok := cm.m[iKey]; !ok {
			cm.m[iKey] = iValue
		}
	}
}

func (cm *CMap[K, V]) Copy() (newCMap *CMap[K, V]) {
	newCMap = NewCMap[K, V]()
	cm.l.RLock()
	defer cm.l.RUnlock()
	for iKey, iValue := range cm.m {
		newCMap.Set(iKey, iValue)
	}
	return newCMap
}

func (cm *CMap[K, V]) Keys() (keys []K) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	keys = make([]K, 0, len(cm.m))
	for iKey, _ := range cm.m {
		keys = append(keys, iKey)
	}
	return keys
}

func (cm *CMap[K, V]) Values() (values []V) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	values = make([]V, 0, len(cm.m))
	for _, iValue := range cm.m {
		values = append(values, iValue)
	}
	return values
}

func (cm *CMap[K, V]) Items() (items []*KeyValue[K,V]) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	items = make([]*KeyValue[K,V], 0, len(cm.m))
	for iKey, iValue := range cm.m {
		items = append(items, &KeyValue[K,V]{
			Key:   iKey,
			Value: iValue,
		})
	}
	return items
}

func (cm *CMap[K, V]) Pop(key K) (value V, ok bool) {
	cm.l.Lock()
	cm.l.Unlock()
	value, ok = cm.m[key]
	if !ok {
		var result V
		return result, ok
	}
	delete(cm.m, key)
	return value, ok
}

func (cm *CMap[K, V]) Foreach(cbfunc func(key K, value V)) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	for iKey, iValue := range cm.m {
		cbfunc(iKey, iValue)
	}
}

func (cm *CMap[K, V]) ForeachUpdate(cbfunc func(key K, value V) (nvalue V)) {
	cm.l.Lock()
	defer cm.l.Unlock()
	for iKey, iValue := range cm.m {
		nValue := cbfunc(iKey, iValue)
		cm.m[iKey] = nValue
	}
}

func NewCMap[K comparable, V any]() *CMap[K, V] {
	return &CMap[K, V]{
		m: make(map[K]V),
	}
}
