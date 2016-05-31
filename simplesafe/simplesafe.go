package simplesafe

import "sync"

type SimpleSafe struct {
	v interface{}
	l *sync.RWMutex
}

func (ss *SimpleSafe) RD(cbfunc func(v interface{})) {
	ss.l.RLock()
	defer ss.l.RUnlock()
	cbfunc(ss.v)
}

func (ss *SimpleSafe) RW(cbfunc func(v interface{})) {
	ss.l.Lock()
	defer ss.l.Unlock()
	cbfunc(ss.v)
}

func NewSimpleSafe(v interface{}) *SimpleSafe {
	return &SimpleSafe {
		v : v,
		l : new(sync.RWMutex),
	}
}
