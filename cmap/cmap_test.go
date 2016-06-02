package cmap

import (
	"testing"
)

func TestLenSetAndGet(t *testing.T) {
	nm := NewCMap()
	act1 := nm.Len()
	exp1 := int(0)
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}

	nm.Set("hoge", 1000)
	act2_1, act2_2 := nm.Get("hoge")
	exp2_1 := 1000
	exp2_2 := true
	if act2_1 != exp2_1 {
		t.Errorf("error: act %v exp %v", act2_1, exp2_1)
	}
	if act2_2 != exp2_2 {
		t.Errorf("error: act %v exp %v", act2_2, exp2_2)
	}

	act3 := nm.Len()
	exp3 := int(1)
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp3)
	}

	_, act4_2 := nm.Get("abc")
	exp4_2 := false
	if act4_2 != exp4_2 {
		t.Errorf("error: act %v exp %v", act4_2, exp4_2)
	}
}
