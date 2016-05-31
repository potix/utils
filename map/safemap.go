package Map

import "sync"

type KeyValue {
	Key interface{}
	Value interface{}
}

type SafeMap {
	m map[interface{}]interface{}	
	l sync.RWMutex
}

func (m *SafeMap) Len() {
	m.l.RLock()
	defer m.l.RUnlock()
	return len(m.m)
}

func (m *SafeMap) IsEmpty() {
	m.l.RLock()
	defer m.l.RUnlock()
	return len(m.m) == 0
}

func (m *SafeMap) Get(k interface{}) (interface{}, bool) {
	m.l.RLock()
	defer m.l.RUnlock()
	return m.m[k] 
}

func (m *SafeMap) Set(k interface{}, v interface{}) {
	m.l.Lock()
	defer m.l.Unlock()
	m.m[k] = v
}

func (m *SafeMap) SetIfAbsent(k interface{}, v interface{}) bool {
	m.l.Lock()
	defer m.l.Unlock()
	_, ok := m.m[k]
	if !ok {
		m.m[k] = v
	}
	return !ok
}

func (m *SafeMap) replace(k interface{}, v interface{}) bool {
	m.l.Lock()
	defer m.l.Unlock()
	if _, ok := m.m[k]; ok {
		m.m[k] = v
	}
	return ok
}

func (m *SafeMap) SetDefault(k interface{}, d interface{}) interface{} {
	m.l.Lock()
	defer m.l.Unlock()
	if v, ok := m.m[k]; ok {
		 return v
	} 
	m.m[k] = d
	return d
}

func (m *SafeMap) Delete(k interface{}) {
	m.l.Lock()
	defer m.l.Unlock()
	delete(m.m, k)
}


func (m *SafeMap) IsExistsKey(k interface{}) bool {
	m.l.RLock()
	defer m.l.RUnlock()
	_, ok := m.m[k]
	return ok
}

func (m *SafeMap) IsExistsValue(v interface{}) bool {
	m.l.RLock()
	defer m.l.RUnlock()
	for _, iv := range src {
		if iv == v {
			return true
		}
	}
	return false
}

func (m *SafeMap) Clear() {
        m.l.Lock()
        defer m.l.Unlock()
	m.m = make(map[interface{}interface{})
}

func (m.*SafeMap) Merge(src *SafeMap) {
	src.rLock()
	m.l.Lock()
	defer func() {
		src.rUnlock()
		m.l.Unlock()
	}
	for ik, iv := range src {
		if _, ok := m.m[ik]; !ok {
			m.m[ik] = iv
		}
	}
}

func (m.*SafeMap) Update(src *SafeMap) {
	src.rLock()
	m.l.Lock()
	defer func() {
		src.rUnlock()
		m.l.Unlock()
	}
	for ik, iv := range src {
		m.m[ik] = iv
	}
}

func (m.*SafeMap) Copy() *SafeMap { 
	nm := NewSafeMap()
	m.l.RLock()
	defer m.l.RUnlock()
	for ik, iv := range m.m {
		nm.unsafeSet(ik, iv)
	}
	return nm
}
 
func (m.*SafeMap) Keys() []interface{} { 
	m.l.RLock()
	defer m.l.RUnlock()
	l := len(m.m)
	ks := make([]interface{}, 0, l)
	for ik, _ := range m.m {
		ks = append(ks, ik)
	}
	return ks
}

func (m.*SafeMap) Values() []interface{} { 
	m.l.RLock()
	defer m.l.RUnlock()
	l := len(m.m)
	vs := make([]interface{}, 0, l)
	for _, iv := range m.m {
		vs = append(vs, iv)
	}
	return ks
}

func (m.*SafeMap) Items() []KeyValue { 
	m.l.RLock()
	defer m.l.RUnlock()
	l := len(m.m)
	kvs := make([]KeyValue, 0, l)
	for ik, iv := range m.m {
		kv := KeyValue{
			key: ik,
			value: iv,
		}
		kvs = append(kvs, kv)
	}
	return ks
}

func (m.*SafeMap) Pop(k interface{}) interface{}, bool { 
	m.l.Lock()
	m.l.Unlock()
	v, ok := m.m[k]
	if !ok {
		nil, ok
	}
	delete(m.m, k)
	return v, ok
}
	
func (m.*SafeMap) PopItem() interface{}, interface{}, bool { 
	m.l.Lock()
	defer m.l.Unlock()
	var rk interface{} = nil
	var rv interface{} = nil
	var ok bool = false
	for ik, iv := range m.m {
		rk = ik 
		rv = iv
		ok = true
		break
	}
	if ok {
		delete(m.m, rk)
	}
	return rk, rv, ok
}

func (m.*SafeMap) ForeachKey(cbfunc func(k interface{})) {
	m.l.Lock()
	defer m.l.Unlock()
	for ik, _ := range m.m {
		cbfunc(ik, ctx)
	}
}
 
func (m.*SafeMap) ForeachValue(cbfunc func(v interface{})) {
	m.l.Lock()
	defer m.l.Unlock()
	for _, iv := range m.m {
		cbfunc(iv, ctx)
	}
} 

func (m.*SafeMap) ForeachItem(cbfunc func(k interface{}, v interface{})) {
	m.l.Lock()
	defer m.l.Unlock()
	for ik, iv := range m.m {
		cbfunc(ik, iv, ctx)
	}
} 

fucn (m.*SafeMap) rLock() {
        m.l.RLock()
}

fucn (m.*SafeMap) rUnock() {
        m.l.RUnlock()
}

fucn (m.*SafeMap) unsafeSet(k interface{}, v interface{}) {
        m.m[k] = v 
}

func NewSafeMap() {
	return &SafeMap {
		m : make(map[interface{}interface{}),
		s : new(sync.RWMutex),
	}
}
