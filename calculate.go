// calculate.go

// result should abondon the unuseful '0' and '.'
// result should have symbol

package main

import (
	"fmt"
	"strings"
)

func adjustNumber(x string) string { //
	negative := ""
	x = strings.Replace(x, " ", "", -1)
	if len(x) == 0 {
		return "0"
	}
	if x[0] == '-' {
		negative += "-"
		x = strings.TrimLeft(x, "-")
	}
	x = strings.Trim(x, "0")
	x = strings.TrimRight(x, ".")
	if len(x) == 0 {
		return "0"
	}
	if x[0] == '.' {
		x = "0" + x
	}
	return negative + x
}
func isvalid(number string) bool { //judge whether the number is valid
	// -xxx.xxx  xxx.xx
	var length = len(number)

	var pointCnt int = 0
	var hasDigit bool = false
	for i := 0; i < length; i++ {
		if number[i] < '0' || number[i] > '9' {
			if number[i] == '.' {
				pointCnt++
				if pointCnt > 1 {
					return false
				}
				if i == length-1 {
					return false
				}
				if hasDigit == false {
					return false
				}

			} else if number[i] == '-' {
				if i != 0 {
					return false
				} else if i == 0 && length == 1 {
					return false
				}
			} else {
				return false
			}
		} else {
			hasDigit = true
		}

	}
	return true
}
func adjustPoint(lst, rst string) (string, string) { //

	var lenl = len(lst)
	var lenr = len(rst)
	if lenl == 0 {
		lst = "0"
	}
	if lenr == 0 {
		rst = "0"
	}
	var indexlst = strings.Index(lst, ".") + 1
	var indexrst = strings.Index(rst, ".") + 1
	//point to point

	if indexlst == 0 {
		if indexrst == 0 { // two integer

		} else { // integer  + float
			lst += "." + strings.Repeat("0", lenr-indexrst)
			lenl = len(lst)
		}
	} else {
		if indexrst == 0 { //float + integer
			rst += "." + strings.Repeat("0", lenl-indexlst)
			lenr = len(rst)
		} else { //float + float
			var a = lenl - indexlst
			var b = lenr - indexrst
			if a > b {
				rst += strings.Repeat("0", a-b)
				lenr = len(rst)
			} else {
				lst += strings.Repeat("0", b-a)
				lenl = len(lst)
			}
		}
	}
	return lst, rst
}
func addd(lst, rst string) string { //// two positive number
	lst, rst = adjustPoint(lst, rst)

	var lenl = len(lst)
	var lenr = len(rst)
	var ans = ""
	var rest = 0
	var lsti = lenl - 1
	var rsti = lenr - 1
	for {
		if lsti < 0 && rsti < 0 {
			break
		}
		if lsti >= 0 {
			rest += int(lst[lsti]) - int(byte('0'))
			lsti--
		}
		if rsti >= 0 {
			rest += int(rst[rsti]) - int(byte('0'))
			rsti--
		}
		ans += string(byte('0' + (rest % 10)))
		rest /= 10
		if (lsti >= 0 && lst[lsti] == '.') || (rsti >= 0 && rst[rsti] == '.') {
			ans += "."
			lsti--
			rsti--
		}
	}
	if rest != 0 {
		ans += string(rest + '0')
		rest = 0
	}
	var result = ""
	for i := len(ans) - 1; i >= 0; i-- {
		result += string(ans[i])
	}
	strings.Trim(result, "0")
	strings.TrimRight(result, ".")
	if len(result) == 0 {
		return "0"
	}
	if result[0] == '.' {
		return "0" + result
	}
	return adjustNumber(result)
}
func add(lst, rst string) string {
	if isvalid(lst) == false || isvalid(rst) == false {
		fmt.Println("invalid number format")
		return "nil"
	}
	lst = adjustNumber(lst)
	rst = adjustNumber(rst)
	if lst[0] == '-' && rst[0] == '-' {
		return "-" + addd(strings.TrimLeft(lst, "-"), strings.TrimLeft(rst, "-"))
	} else if lst[0] == '-' && rst[0] != '-' {
		return subb(rst, strings.TrimLeft(lst, "-"))
	} else if lst[0] != '-' && rst[0] == '-' {
		return subb(lst, strings.TrimLeft(rst, "-"))
	} else { //if lst[0] != '-' && rst[0] != '-'
		return addd(lst, rst)
	}
}
func eraseSymbol(ele string) string {
	ele = strings.TrimLeft(ele, "-")
	ele = strings.TrimLeft(ele, "+")
	ele = strings.TrimLeft(ele, "-")
	return ele
}
func compare(lst, rst string) string { // input two unsigned number
	lst = eraseSymbol(lst)
	rst = eraseSymbol(rst)
	lst, rst = adjustPoint(lst, rst)
	len1 := len(lst)
	len2 := len(rst)
	if len1 == len2 {
		for i := 0; i < len1; i++ {
			if lst[i] > rst[i] {
				return ">"
			} else if lst[i] < rst[i] {
				return "<"
			}

		}
		return "="
	} else if len1 > len2 {
		return ">"
	} else {
		return "<"
	}
}
func subb(lst, rst string) string {
	lst, rst = adjustPoint(lst, rst)
	com := compare(lst, rst)
	len1 := len(lst)
	len2 := len(rst)
	negative := ""
	if com == "<" {
		negative += "-"
		lst, rst = rst, lst
		len1, len2 = len2, len1
	}
	fmt.Printf("%d %d %s %s\n", len1, len2, lst, rst)
	var tmp [500]rune
	var tmp1 [500]rune

	for i := 0; i < len1; i++ {
		tmp[i] = rune(lst[i])
	}
	for i := 0; i < len2; i++ {
		tmp1[i] = rune(rst[i])
	}
	for i := 0; i < len1 || i < len2; i++ {
		if lst[len1-i-1] == '.' {
			continue
		}
		if tmp[len1-i-1] < rune('0') {
			tmp[len1-i-1] += 10
			if tmp[len1-i-2] != rune('.') {
				tmp[len1-i-2] -= 1
			} else {
				tmp[len1-i-3] -= 1
			}
		}
		if i < len2 {
			tmp[len1-i-1] -= tmp1[len2-i-1] - '0'
		}
		if tmp[len1-i-1] < '0' {
			tmp[len1-i-1] += 10
			if tmp[len1-i-2] != '.' {
				tmp[len1-i-2] -= 1
			} else {
				tmp[len1-i-3] -= 1
			}
		}

	}
	ans := ""
	for i := 0; i < len(tmp); i++ {
		if tmp[i] == 0 {
			break
		}
		ans += string(tmp[i])
	}
	ans = negative + ans
	return adjustNumber(ans)
}
func sub(lst, rst string) string {
	if isvalid(lst) == false || isvalid(rst) == false {
		fmt.Println("invalid number format")
		return "nil"
	}
	lst = adjustNumber(lst)
	rst = adjustNumber(rst)
	if lst[0] == '-' && rst[0] == '-' {
		return subb(strings.TrimLeft(rst, "-"), strings.TrimLeft(lst, "-"))
	} else if lst[0] != '-' && rst[0] == '-' {
		return add(strings.TrimLeft(rst, "-"), lst)
	} else if lst[0] != '-' && rst[0] != '-' {
		return subb(lst, rst)
	} else { //if lst[0] == '-' && rst[0] != '-'
		return "-" + add(rst, strings.TrimLeft(lst, "-"))
	}
}

func main() {
	for {
		fmt.Println("please input a expression separate by space   format : a + b")
		var num1, op, num2 string
		fmt.Scanf("%s %s %s\n", &num1, &op, &num2)
		if op == "+" {
			fmt.Printf("result = %s\n", add(num1, num2))
		} else if op == "-" {
			fmt.Printf("result = %s\n", sub(num1, num2))
		}

		/*var num1, num2 string
		fmt.Scanf("%s %s\n", &num1, &num2)
		num1, num2 = justPoint(num1, num2)
		fmt.Printf("%s %s\n", num1, num2)
		*/
	}
}
