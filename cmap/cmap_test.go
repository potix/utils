package cmap

import (
	"fmt"
	"testing"
)

func TestLen(t *testing.T) {
	nm := NewCMap[string, int]()
	act1 := nm.Len()
	exp1 := 0
	nm.Set("hoge", 1000)
	act2 := nm.Len()
	exp2 := 1
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
}

func TestSetGet(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	act1, act2 := nm.Get("hoge")
	exp1 := 1000
	exp2 := true
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	act3, act4 := nm.Get("abc")
	var exp3 int
	exp4 := false
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp3)
	}
	if act4 != exp4 {
		t.Errorf("error: act %v exp %v", act4, exp4)
	}
}

func updatecb(v int) int {
	return v + 1000
}

func TestSetUpdate(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	act1, act2 := nm.Get("hoge")
	exp1 := 1000
	exp2 := true
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	act3 := nm.Update("hoge", updatecb)
	exp3 := true
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp3)
	}
	act4, act5 := nm.Get("hoge")
	exp4 := 2000
	exp5 := true
	if act4 != exp4 {
		t.Errorf("error: act %v exp %v", act4, exp4)
	}
	if act5 != exp5 {
		t.Errorf("error: act %v exp %v", act5, exp5)
	}
}

func TestSetIfAbsend(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	act0 := nm.SetIfAbsent("hoge", 100)
	act1, act2 := nm.Get("hoge")
	exp0 := false
	exp1 := 1000
	exp2 := true
	if act0 != exp0 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	act3 := nm.SetIfAbsent("hoge2", 100)
	act4, act5 := nm.Get("hoge2")
	exp3 := true
	exp4 := 100
	exp5 := true
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp3)
	}
	if act4 != exp4 {
		t.Errorf("error: act %v exp %v", act4, exp4)
	}
	if act5 != exp5 {
		t.Errorf("error: act %v exp %v", act5, exp5)
	}
}

func TestSetIfExist(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	act0 := nm.SetIfExist("hoge", 100)
	act1, act2 := nm.Get("hoge")
	exp0 := true
	exp1 := 100
	exp2 := true
	if act0 != exp0 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	act3 := nm.SetIfExist("hoge2", 100)
	act4, act5 := nm.Get("hoge2")
	exp3 := false
	var exp4 int
	exp5 := false
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp3)
	}
	if act4 != exp4 {
		t.Errorf("error: act %v exp %v", act4, exp4)
	}
	if act5 != exp5 {
		t.Errorf("error: act %v exp %v", act5, exp5)
	}
}

func TestGetWithDefault(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	act1 := nm.GetWithDefault("hoge", 222)
	exp1 := 1000
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	act2 := nm.GetWithDefault("hoge2", 222)
	exp2 := 222
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	act3, act4 := nm.Get("hoge2")
	exp3 := 222
	exp4 := true
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp4)
	}
	if act4 != exp4 {
		t.Errorf("error: act %v exp %v", act3, exp4)
	}
}

func TestDelete(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	nm.Set("hoge2", 100)
	act1 := nm.Len()
	exp1 := 2
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	act2 := nm.Delete("hoge")
	exp2 := true
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	act3 := nm.Len()
	exp3 := 1
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp3)
	}
	act4 := nm.Delete("fuga")
	exp4 := false
	if act4 != exp4 {
		t.Errorf("error: act %v exp %v", act4, exp4)
	}
	act5 := nm.Len()
	exp5 := 1
	if act5 != exp5 {
		t.Errorf("error: act %v exp %v", act5, exp5)
	}
}

func TestClear(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	nm.Set("hoge2", 100)
	act1 := nm.Len()
	exp1 := 2
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	nm.Clear()
	act2 := nm.Len()
	exp2 := 0
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
}

func TestOverWriteMerge(t *testing.T) {
	nm1 := NewCMap[string, int]()
	nm1.Set("hoge", 1000)
	nm1.Set("aaaa", 222)
	act1 := nm1.Len()
	exp1 := 2
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	nm2 := NewCMap[string, int]()
	nm2.Set("fuga", 2000)
	nm2.Set("aaaa", 111)
	act2 := nm2.Len()
	exp2 := 2
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	nm1.OverwriteMerge(nm2)
	act3 := nm1.Len()
	exp3 := 3
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp3)
	}
	act4, act5 := nm1.Get("hoge")
	act6, act7 := nm1.Get("aaaa")
	act8, act9 := nm1.Get("fuga")
	exp4 := 1000
	exp5 := true
	exp6 := 111
	exp7 := true
	exp8 := 2000
	exp9 := true
	if act4 != exp4 {
		t.Errorf("error: act %v exp %v", act4, exp4)
	}
	if act5 != exp5 {
		t.Errorf("error: act %v exp %v", act5, exp5)
	}
	if act6 != exp6 {
		t.Errorf("error: act %v exp %v", act6, exp6)
	}
	if act7 != exp7 {
		t.Errorf("error: act %v exp %v", act7, exp7)
	}
	if act8 != exp8 {
		t.Errorf("error: act %v exp %v", act8, exp8)
	}
	if act9 != exp9 {
		t.Errorf("error: act %v exp %v", act9, exp9)
	}
}

func TestKeepMerge(t *testing.T) {
	nm1 := NewCMap[string, int]()
	nm1.Set("hoge", 1000)
	nm1.Set("aaaa", 222)
	act1 := nm1.Len()
	exp1 := 2
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	nm2 := NewCMap[string, int]()
	nm2.Set("fuga", 2000)
	nm2.Set("aaaa", 111)
	act2 := nm2.Len()
	exp2 := 2
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	nm1.KeepMerge(nm2)
	act3 := nm1.Len()
	exp3 := 3
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp3)
	}
	act4, act5 := nm1.Get("hoge")
	act6, act7 := nm1.Get("aaaa")
	act8, act9 := nm1.Get("fuga")
	exp4 := 1000
	exp5 := true
	exp6 := 222
	exp7 := true
	exp8 := 2000
	exp9 := true
	if act4 != exp4 {
		t.Errorf("error: act %v exp %v", act4, exp4)
	}
	if act5 != exp5 {
		t.Errorf("error: act %v exp %v", act5, exp5)
	}
	if act6 != exp6 {
		t.Errorf("error: act %v exp %v", act6, exp6)
	}
	if act7 != exp7 {
		t.Errorf("error: act %v exp %v", act7, exp7)
	}
	if act8 != exp8 {
		t.Errorf("error: act %v exp %v", act8, exp8)
	}
	if act9 != exp9 {
		t.Errorf("error: act %v exp %v", act9, exp9)
	}
}

func TestCopy(t *testing.T) {
	nm1 := NewCMap[string, int]()
	nm1.Set("hoge", 1000)
	nm1.Set("fuga", 2000)
	act1 := nm1.Len()
	exp1 := 2
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	nm2 := nm1.Copy()
	act2 := nm2.Len()
	exp2 := 2
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	act3, act4 := nm2.Get("hoge")
	act5, act6 := nm2.Get("fuga")
	exp3 := 1000
	exp4 := true
	exp5 := 2000
	exp6 := true
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp3)
	}
	if act4 != exp4 {
		t.Errorf("error: act %v exp %v", act4, exp4)
	}
	if act5 != exp5 {
		t.Errorf("error: act %v exp %v", act5, exp5)
	}
	if act6 != exp6 {
		t.Errorf("error: act %v exp %v", act6, exp6)
	}
}

func TestKeys(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	nm.Set("fuga", 2000)
	act1 := nm.Len()
	exp1 := 2
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	keys := nm.Keys()
	act2 := len(keys)
	exp2 := 2
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	for _, k := range keys {
		if k != "hoge" && k != "fuga" {
			t.Errorf("error: act %v", k)
		}
	}
}

func TestValues(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	nm.Set("fuga", 2000)
	act1 := nm.Len()
	exp1 := 2
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	values := nm.Values()
	act2 := len(values)
	exp2 := 2
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	for _, v := range values {
		if v != 1000 && v != 2000 {
			t.Errorf("error: act %v", v)
		}
	}
}

func TestItems(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	nm.Set("fuga", 2000)
	act1 := nm.Len()
	exp1 := 2
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	items := nm.Items()
	act2 := len(items)
	exp2 := 2
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	for _, i := range items {
		if i.Key != "hoge" && i.Key != "fuga" {
			t.Errorf("error: act %v", i.Key)
		}
		if i.Value != 1000 && i.Value != 2000 {
			t.Errorf("error: act %v", i.Value)
		}
	}
}

func TestPop(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	nm.Set("fuga", 2000)
	act1 := nm.Len()
	exp1 := 2
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	act2, act3 := nm.Pop("fuga")
	exp2 := 2000
	exp3 := true
	if act2 != exp2 {
		t.Errorf("error: act %v exp %v", act2, exp2)
	}
	if act3 != exp3 {
		t.Errorf("error: act %v exp %v", act3, exp3)
	}
	act4 := nm.Len()
	exp4 := 1
	if act4 != exp4 {
		t.Errorf("error: act %v exp %v", act4, exp4)
	}
	act5, act6 := nm.Pop("fuga")
	var exp5 int
	exp6 := false
	if act5 != exp5 {
		t.Errorf("error: act %v exp %v", act5, exp5)
	}
	if act6 != exp6 {
		t.Errorf("error: act %v exp %v", act6, exp6)
	}
}

func cbFunc1(key string, value int) {
	fmt.Printf("%v, %v\n", key, value)
}

func TestForeach(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	nm.Set("fuga", 2000)
	nm.Set("aaaa", 3000)
	nm.Set("bbbb", 4000)
	act1 := nm.Len()
	exp1 := 4
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	nm.Foreach(cbFunc1)
}

func cbFunc3(key string, value int) (int) {
	fmt.Printf("%v, %v\n", key, value)
	return value
}

func cbFunc4(key string, value int) (int) {
	return value + 500
}

func TestForeachUpdate(t *testing.T) {
	nm := NewCMap[string, int]()
	nm.Set("hoge", 1000)
	nm.Set("fuga", 2000)
	nm.Set("aaaa", 3000)
	nm.Set("bbbb", 4000)
	nm.Set("cccc", 4000)
	act1 := nm.Len()
	exp1 := 5
	if act1 != exp1 {
		t.Errorf("error: act %v exp %v", act1, exp1)
	}
	fmt.Printf("-----\n")
	nm.ForeachUpdate(cbFunc3)
	fmt.Printf("-----\n")
	nm.ForeachUpdate(cbFunc4)
	fmt.Printf("-----\n")
	nm.ForeachUpdate(cbFunc3)
	fmt.Printf("-----\n")
}
