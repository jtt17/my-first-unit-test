// func_adjustNumber_test.go
package main

import (
	"testing"
)

func TestAdjustNumber(t *testing.T) {
	if ans := adjustNumber("-0.001230"); ans != "-0.00123" {
		t.Error("adjustNumber test1 failed ", ans)
	} else {
		t.Log("adjustNumber test1 passed ", ans)
	}
	if ans := adjustNumber("0123.00"); ans != "123" {
		t.Error("adjustNumber test2 failed ", ans)
	} else {
		t.Log("adjustNumber test2 passed ", ans)
	}
	if ans := adjustNumber("-12245.000"); ans != "-12245" {
		t.Error("adjustNumber test3 failed ", ans)
	} else {
		t.Log("adjustNumber test3 passed ", ans)
	}
	if ans := adjustNumber(""); ans != "0" {
		t.Error("adjustNumber test4 failed ", ans)
	} else {
		t.Log("adjustNumber test4 passed ", ans)
	}
	if ans := adjustNumber("-000.000"); ans != "0" {
		t.Error("adjustNumber test5 failed ", ans)
	} else {
		t.Log("adjustNumber test5 passed ", ans)
	}
}
