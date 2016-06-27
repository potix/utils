package cmap

import (
	"reflect"
	"sync"
)

type KeyValue struct {
	Key   interface{}
	Value interface{}
}

type CMap struct {
	m map[interface{}]interface{}
	l *sync.RWMutex
}

func (cm *CMap) Len() (length int) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	return len(cm.m)
}

func (cm *CMap) IsEmpty() (empty bool) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	return len(cm.m) == 0
}

func (cm *CMap) Get(key interface{}) (value interface{}, ok bool) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	value, ok = cm.m[key]
	return value, ok
}

func (cm *CMap) Set(key interface{}, value interface{}) {
	cm.l.Lock()
	defer cm.l.Unlock()
	cm.m[key] = value
}

func (cm *CMap) SetIfAbsent(key interface{}, value interface{}) (ok bool) {
	cm.l.Lock()
	defer cm.l.Unlock()
	_, ok = cm.m[key]
	if !ok {
		cm.m[key] = value
	}
	return !ok
}

func (cm *CMap) Replace(key interface{}, value interface{}) (ok bool) {
	cm.l.Lock()
	defer cm.l.Unlock()
	_, ok = cm.m[key]
	if ok {
		cm.m[key] = value
	}
	return ok
}

func (cm *CMap) GetDefault(key interface{}, defValue interface{}) (value interface{}) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	value, ok := cm.m[key]
	if !ok {
		return defValue
	}
	return value
}

func (cm *CMap) GetAndSetDefault(key interface{}, defValue interface{}) (value interface{}) {
	cm.l.Lock()
	defer cm.l.Unlock()
	value, ok := cm.m[key]
	if ok {
		cm.m[key] = defValue
		return defValue
	}
	return value
}

func (cm *CMap) Delete(key interface{}) (ok bool) {
	cm.l.Lock()
	defer cm.l.Unlock()
	_, ok = cm.m[key]
	if ok {
		delete(cm.m, key)
	}
	return ok
}

func (cm *CMap) IsExistsKey(key interface{}) (ok bool) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	_, ok = cm.m[key]
	return ok
}

func (cm *CMap) IsExistsValue(value interface{}) (ok bool) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	for _, iValue := range cm.m {
		if reflect.DeepEqual(iValue, value) {
			return true
		}
	}
	return false
}

func (cm *CMap) Clear() {
	cm.l.Lock()
	defer cm.l.Unlock()
	cm.m = make(map[interface{}]interface{})
}

func (cm *CMap) Merge(src *CMap) {
	src.rLock()
	cm.l.Lock()
	defer func() {
		src.rUnlock()
		cm.l.Unlock()
	}()
	for iKey, iValue := range src.m {
		if _, ok := cm.m[iKey]; !ok {
			cm.m[iKey] = iValue
		}
	}
}

func (cm *CMap) Update(src *CMap) {
	src.rLock()
	cm.l.Lock()
	defer func() {
		src.rUnlock()
		cm.l.Unlock()
	}()
	for iKey, iValue := range src.m {
		cm.m[iKey] = iValue
	}
}

func (cm *CMap) Copy() (newCMap *CMap) {
	newCMap = NewCMap()
	cm.l.RLock()
	defer cm.l.RUnlock()
	for iKey, iValue := range cm.m {
		newCMap.unsafeSet(iKey, iValue)
	}
	return newCMap
}

func (cm *CMap) Keys() (keys []interface{}) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	length := len(cm.m)
	keys = make([]interface{}, 0, length)
	for iKey, _ := range cm.m {
		keys = append(keys, iKey)
	}
	return keys
}

func (cm *CMap) Values() (values []interface{}) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	length := len(cm.m)
	values = make([]interface{}, 0, length)
	for _, iValue := range cm.m {
		values = append(values, iValue)
	}
	return values
}

func (cm *CMap) Items() (items []*KeyValue) {
	cm.l.RLock()
	defer cm.l.RUnlock()
	length := len(cm.m)
	items = make([]*KeyValue, 0, length)
	for iKey, iValue := range cm.m {
		items = append(items, &KeyValue{
			Key:   iKey,
			Value: iValue,
		})
	}
	return items
}

func (cm *CMap) Pop(k interface{}) (value interface{}, ok bool) {
	cm.l.Lock()
	cm.l.Unlock()
	value, ok = cm.m[k]
	if !ok {
		return nil, ok
	}
	delete(cm.m, k)
	return value, ok
}

func (cm *CMap) PopItem() (key interface{}, value interface{}, ok bool) {
	cm.l.Lock()
	defer cm.l.Unlock()
	for iKey, iValue := range cm.m {
		key = iKey
		value = iValue
		ok = true
		break
	}
	if ok {
		delete(cm.m, key)
	}
	return key, value, ok
}

func (cm *CMap) ForeachKey(cbfunc func(key interface{}) (loopBreak bool)) {
	cm.l.Lock()
	defer cm.l.Unlock()
	for iKey, _ := range cm.m {
		loopBreak := cbfunc(iKey)
		if loopBreak {
			break
		}
	}
}

func (cm *CMap) ForeachValue(cbfunc func(value interface{}) (loopBreak bool)) {
	cm.l.Lock()
	defer cm.l.Unlock()
	for _, iValue := range cm.m {
		loopBreak := cbfunc(iValue)
		if loopBreak {
			break
		}
	}
}

func (cm *CMap) ForeachItem(cbfunc func(key interface{}, value interface{}) (loopBreak bool)) {
	cm.l.Lock()
	defer cm.l.Unlock()
	for iKey, iValue := range cm.m {
		loopBreak := cbfunc(iKey, iValue)
		if loopBreak {
			break
		}
	}
}

func (cm *CMap) rLock() {
	cm.l.RLock()
}

func (cm *CMap) rUnlock() {
	cm.l.RUnlock()
}

func (cm *CMap) unsafeSet(key interface{}, value interface{}) {
	cm.m[key] = value
}

func NewCMap() *CMap {
	return &CMap{
		m: make(map[interface{}]interface{}),
		l: new(sync.RWMutex),
	}
}
