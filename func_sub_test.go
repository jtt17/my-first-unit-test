// func_sub_test.go
package main

import (
	"testing"
)

func TestSub(t *testing.T) {
	if ans := sub("-1.0", "1.0"); ans != "-2" {
		t.Error("sub test1 failed ", ans)
	} else {
		t.Log("sub test1 passed ", ans)
	}
	if ans := sub("-1-", "1.0"); ans != "nil" {
		t.Error("sub test2 failed ", ans)
	} else {
		t.Log("sub test2 passed ", ans)
	}
	if ans := sub("1231537894565614564864164.12", "-1.0"); ans != "1231537894565614564864165.12" {
		t.Error("sub test3 failed ", ans)
	} else {
		t.Log("sub test3 passed ", ans)
	}
	if ans := sub("3.1", "123.9"); ans != "-120.8" {
		t.Error("sub test4 failed ", ans)
	} else {
		t.Log("sub test4 passed ", ans)
	}
	if ans := sub("-0.08", "123.9"); ans != "-123.98" {
		t.Error("sub test5 failed ", ans)
	} else {
		t.Log("sub test5 passed ", ans)
	}
	if ans := sub("-123.1", "-123.9"); ans != "0.8" {
		t.Error("sub test6 failed ", ans)
	} else {
		t.Log("sub test6 passed ", ans)
	}
}
