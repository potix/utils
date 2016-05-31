package ssmap

import (
	"testing"
)

func TestLenSetAndGet(t *testing.T) {
	nm := NewSafeMap()
	nm.ROnly(func(m map[interface{}]interface{}) {
		act1 := len(m)
		exp1 := int(0)
		if act1 != exp1 {
			t.Errorf("error: act %v exp %v", act1, exp1)
		}
	})

	nm.RW(func(m map[interface{}]interface{}) {
		m["hoge"] =  1000
		act2_1, act2_2 := m["hoge"]
		exp2_1 := 1000
		exp2_2 := true
		if act2_1 != exp2_1 {
			t.Errorf("error: act %v exp %v", act2_1, exp2_1)
		}
		if act2_2 != exp2_2 {
			t.Errorf("error: act %v exp %v", act2_2, exp2_2)
		}
	})

	nm.ROnly(func(m map[interface{}]interface{}) {
		act3 := len(m)
		exp3 := int(1)
		if act3 != exp3 {
			t.Errorf("error: act %v exp %v", act3, exp3)
		}
	})

	nm.ROnly(func(m map[interface{}]interface{}) {
		_, act4_2 := m["abc"]
		exp4_2 := false
		if act4_2 != exp4_2 {
			t.Errorf("error: act %v exp %v", act4_2, exp4_2)
		}
	})
}
