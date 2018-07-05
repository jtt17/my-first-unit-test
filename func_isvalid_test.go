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
		t.Error("isvalid test2 failed ", ans)
	} else {
		t.Log("isvalid test2 passed ", ans)
	}
	if ans := isvalid("-0.0-0"); ans != false {
		t.Error("isvalid test3 failed ", ans)
	} else {
		t.Log("isvalid test3 passed ", ans)
	}
	if ans := isvalid("-0."); ans != false {
		t.Error("isvalid test4 failed ", ans)
	} else {
		t.Log("isvalid test4 passed ", ans)
	}
	if ans := isvalid("-0abc"); ans != false {
		t.Error("isvalid test5 failed ", ans)
	} else {
		t.Log("isvalid test5 passed ", ans)
	}
	if ans := isvalid("0x3s-"); ans != false {
		t.Error("isvalid test6 failed ", ans)
	} else {
		t.Log("isvalid test6 passed ", ans)
	}
	if ans := isvalid(""); ans != false {
		t.Error("isvalid test7 failed ", ans)
	} else {
		t.Log("isvalid test7 passed ", ans)
	}
}
