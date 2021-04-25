package random

import (
	"fmt"
	"testing"
)

func TestGetRandomString(t *testing.T) {
	r, err := GetRandomString(32)
	if err != nil {
		t.Fatalf("GetRandomString error %v", err)
	}
	if len(r) != 32 {
		t.Fatalf("length mismatch %v", err)
	}
	fmt.Println(r)

	r, err = GetRandomString(0)
	if err == nil {
		t.Fatalf("GetRandomString error %v", err)
	}
	if len(r) != 0 {
		t.Fatalf("length mismatch %v", err)
	}
	fmt.Println(r)
}


func TestGetAlphanumericRandomString(t *testing.T) {
	r, err := GetAlphanumericRandomString(32)
	if err != nil {
		t.Fatalf("GetRandomString error %v", err)
	}
	if len(r) != 32 {
		t.Fatalf("length mismatch %v", err)
	}
	fmt.Println(r)

	r, err = GetAlphanumericRandomString(0)
	if err == nil {
		t.Fatalf("GetRandomString error %v", err)
	}
	if len(r) != 0 {
		t.Fatalf("length mismatch %v", err)
	}
	fmt.Println(r)
}
