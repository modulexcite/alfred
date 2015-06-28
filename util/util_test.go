package util

import (
	"reflect"
	"testing"
)

func TestToIntf(t *testing.T) {
	if reflect.TypeOf(ToIntf([]int{1, 2, 3})) != reflect.TypeOf([]interface{}{}) {
		t.Fatal("Did not convert slice to intf slice")
	}
}

func TestIn(t *testing.T) {
	s := []string{"foo", "bar", "kuku", "kiki"}
	for _, v := range s {
		if !In(s, v) {
			t.Error("Should be in")
		}
	}
	if In(s, "foobar") {
		t.Error("Should not be in")
	}
}

func TestToLower(t *testing.T) {
	s := []string{"MyString12Str"}
	res := ToLower(s)
	if res[0] != "mystring12str" {
		t.Error(res)
	}

}

func TestRandStr(t *testing.T) {
	s1 := RandStr(32)
	s2 := RandStr(32)
	if len(s1) != 32 {
		t.Errorf("Rand str len not enforced s1 %d", len(s1))
	}
	if len(s2) != 32 {
		t.Errorf("Rand str len not enforced s2 %d", len(s2))
	}
	if s1 == s2 {
		t.Error("Random string is not random")
	}

}

func TestCannoncialize(t *testing.T) {
	url1 := "http://first/test.com"
	url2 := "https://second/test.com"
	res := Canonicalize(url1, url2)
	if len(res) != 2 {
		t.Errorf("length problem: %d", len(res))
	}
	if res[0] != "http://first" || res[1] != "https://second" {
		t.Error(res)
	}
}
