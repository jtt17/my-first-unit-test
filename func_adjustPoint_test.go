// func_adjustPoint_test.go
package main

import (
	"testing"
)

func TestAdjustPoint(t *testing.T) {
	if a, b := adjustPoint("-3.143", "1.0"); a != "-3.143" || b != "1.000" {
		t.Error("failed adjustPoint test1 ", a, " ", b)
	} else {
		t.Log("pass adjustPoint test1 ", a, " ", b)
	}
	if a, b := adjustPoint("", ""); a != "0" || b != "0" {
		t.Error("failed adjustPoint test2 ", a, " ", b)
	} else {
		t.Log("pass adjustPoint test2 ", a, " ", b)
	}
	if a, b := adjustPoint("1", "2.1"); a != "1.0" || b != "2.1" {
		t.Error("failed adjustPoint test3 ", a, " ", b)
	} else {
		t.Log("pass adjustPoint test3 ", a, " ", b)
	}
	if a, b := adjustPoint("1", "2"); a != "1" || b != "2" {
		t.Error("failed adjustPoint test4 ", a, " ", b)
	} else {
		t.Log("pass adjustPoint test4 ", a, " ", b)
	}
	if a, b := adjustPoint("1", "2.01"); a != "1.00" || b != "2.01" {
		t.Error("failed adjustPoint test5 ", a, " ", b)
	} else {
		t.Log("pass adjustPoint test5 ", a, " ", b)
	}
	if a, b := adjustPoint("0.001", "0.11"); a != "0.001" || b != "0.110" {
		t.Error("failed adjustPoint test6 ", a, " ", b)
	} else {
		t.Log("pass adjustPoint test6 ", a, " ", b)
	}
	if a, b := adjustPoint("110.001", "1"); a != "110.001" || b != "1.000" {
		t.Error("failed adjustPoint test7 ", a, " ", b)
	} else {
		t.Log("pass adjustPoint test7 ", a, " ", b)
	}
}
