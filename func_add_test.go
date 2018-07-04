// caculate_test.go
package main

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if ans := add("-1.0", "1.0"); ans != "0" {
		t.Error("add test1 failed ", ans)
	} else {
		t.Log("add test1 passed ", ans)
	}

	if ans := add("-1.0-", "1.0"); ans != "nil" {
		t.Error("add test2 failed ", ans)
	} else {
		t.Log("add test2 passed ", ans)
	}

	if ans := add("786456456.134531", "-4123.112001"); ans != "786452333.02253" {
		t.Error("add test3 failed ", ans)
	} else {
		t.Log("add test1 passed ", ans)
	}

	if ans := add("-45.1", "-0.1"); ans != "-45.2" {
		t.Error("add test4 failed ", ans)
	} else {
		t.Log("add test1 passed ", ans)
	}
	if ans := add("-45.1", "-0.1"); ans != "-45.2" {
		t.Error("add test4 failed ", ans)
	} else {
		t.Log("add test1 passed ", ans)
	}
	if ans := add("1.0", "-0.1"); ans != "0.9" {
		t.Error("add test5 failed ", ans)
	} else {
		t.Log("add test1 passed ", ans)
	}
	if ans := add("-00.1", "99.00"); ans != "98.9" {
		t.Error("add test6 failed ", ans)
	} else {
		t.Log("add test1 passed ", ans)
	}
	if ans := add("-0.100", "94.001"); ans != "93.901" {
		t.Error("add test4 failed ", ans)
	} else {
		t.Log("add test1 passed ", ans)
	}
}
