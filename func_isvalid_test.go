// func_isvalid_test.go
package main

import (
	"testing"
)

func TestIsvalid(t *testing.T) {
	if ans := isvalid("-0.001230"); ans != true {
		t.Error("isvalid test1 failed ", ans)
	} else {
		t.Log("isvalid test1 passed ", ans)
	}
	if ans := isvalid("0.0.1"); ans != false {
		t.Error("isvalid test failed ", ans)
	} else {
		t.Log("isvalid test passed ", ans)
	}
	if ans := isvalid("-0.0-0"); ans != false {
		t.Error("isvalid test failed ", ans)
	} else {
		t.Log("isvalid test passed ", ans)
	}
	if ans := isvalid("-0."); ans != false {
		t.Error("isvalid test failed ", ans)
	} else {
		t.Log("isvalid test passed ", ans)
	}
	if ans := isvalid("-0abc"); ans != false {
		t.Error("isvalid test failed ", ans)
	} else {
		t.Log("isvalid test passed ", ans)
	}
	if ans := isvalid("0x3s-"); ans != false {
		t.Error("isvalid test failed ", ans)
	} else {
		t.Log("isvalid test passed ", ans)
	}
}
