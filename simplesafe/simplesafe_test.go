package simplesafe

import (
	"testing"
)

func TestLenSetAndGet(t *testing.T) {
	ss := NewSimpleSafe(make(map[string]int))
	ss.RD(func(v interface{}) {
		act1 := len(v.(map[string]int))
		exp1 := int(0)
		if act1 != exp1 {
			t.Errorf("error: act %v exp %v", act1, exp1)
		}
	})

	ss.RW(func(v interface{}) {
		v.(map[string]int)["hoge"] =  1000
		act2_1, act2_2 := v.(map[string]int)["hoge"]
		exp2_1 := 1000
		exp2_2 := true
		if act2_1 != exp2_1 {
			t.Errorf("error: act %v exp %v", act2_1, exp2_1)
		}
		if act2_2 != exp2_2 {
			t.Errorf("error: act %v exp %v", act2_2, exp2_2)
		}
	})

	ss.RD(func(v interface{}) {
		act3 := len(v.(map[string]int))
		exp3 := int(1)
		if act3 != exp3 {
			t.Errorf("error: act %v exp %v", act3, exp3)
		}
	})

	ss.RD(func(v interface{}) {
		_, act4_2 := v.(map[string]int)["abc"]
		exp4_2 := false
		if act4_2 != exp4_2 {
			t.Errorf("error: act %v exp %v", act4_2, exp4_2)
		}
	})
}
