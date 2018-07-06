// calculate.go

// result should abondon the unuseful '0' and '.'
// result should have symbol if negative

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func adjustNumber(x string) string { // make the input number more standard
	if isvalid(x) == false || len(x) == 0 {
		return "0"
	}
	negative := ""
	if x[0] == '-' {
		negative += "-"
		x = x[1:]
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
	var (
		length        = len(number)
		pointCnt int  = 0
		hasDigit bool = false
	)
	if length == 0 {
		return false
	}
	for i := 0; i < length; i++ {
		if number[i] < '0' || number[i] > '9' {
			if number[i] == '.' {
				pointCnt++
				if pointCnt > 1 { //two many decimal point
					return false
				}
				if i == length-1 { // there's no digit after points
					return false
				}
				if hasDigit == false { // there's no digit before points
					return false
				}

			} else if number[i] == '-' {
				if i != 0 { // symbol should at the beginning of number
					return false
				} else if i == 0 && length == 1 { // only symbol
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

	var (
		lenl     = len(lst)
		lenr     = len(rst)
		indexlst = strings.Index(lst, ".") + 1 // find the decimal point position
		indexrst = strings.Index(rst, ".") + 1
	)
	if lenl == 0 {
		lst = "0"
	}
	if lenr == 0 {
		rst = "0"
	}
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
func addd(lst, rst string) string { // two unsigned number addition
	lst, rst = adjustPoint(lst, rst)

	var (
		lenl = len(lst)
		lenr = len(rst)
		ans  = ""
		rest = 0
		lsti = lenl - 1
		rsti = lenr - 1
	)
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
	result := ""
	for i := len(ans) - 1; i >= 0; i-- {
		result += string(ans[i])
	}
	strings.Trim(result, "0")
	strings.TrimRight(result, ".")
	if len(result) == 0 {
		return "0"
	}
	if result[0] == '.' {
		result = "0" + result
	}
	return adjustNumber(result)
}
func add(lst, rst string) string { // deal with differ situation
	if isvalid(lst) == false || isvalid(rst) == false {
		fmt.Println("invalid number format")
		return "nil"
	}
	lst = adjustNumber(lst)
	rst = adjustNumber(rst)
	if lst[0] == '-' && rst[0] == '-' {
		return "-" + addd(lst[1:], rst[1:])
	} else if lst[0] == '-' && rst[0] != '-' {
		return subb(rst, lst[1:])
	} else if lst[0] != '-' && rst[0] == '-' {
		return subb(lst, rst[1:])
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
	lenl := len(lst)
	lenr := len(rst)
	if lenl == lenr {
		for i := 0; i < lenl; i++ {
			if lst[i] > rst[i] {
				return ">"
			} else if lst[i] < rst[i] {
				return "<"
			}
		}
		return "="
	} else if lenl > lenr {
		return ">"
	} else {
		return "<"
	}
}
func subb(lst, rst string) string { // two unsigned real number subtraction
	lst, rst = adjustPoint(lst, rst)
	com := compare(lst, rst)
	lenl := len(lst)
	lenr := len(rst)
	negative := ""
	if com == "<" {
		negative += "-"
		lst, rst = rst, lst
		lenl, lenr = lenr, lenl
	}
	array1 := []byte(lst) //directly get the slice of string 'lst'
	for i := 0; i < lenl || i < lenr; i++ {
		if lst[lenl-i-1] == '.' {
			continue
		}
		if array1[lenl-i-1] < byte('0') {
			array1[lenl-i-1] += 10
			if lenl-i-2 >= 0 && lst[lenl-i-2] != byte('.') {
				array1[lenl-i-2] -= 1
			} else {
				if lenl-i-3 >= 0 {
					array1[lenl-i-3] -= 1
				} else {
					fmt.Printf("unreachable position!!\n")
				}

			}
		}
		if i < lenr {
			array1[lenl-i-1] -= byte(rst[lenr-i-1])
			array1[lenl-i-1] += byte('0')
		}
		if array1[lenl-i-1] < byte('0') {
			array1[lenl-i-1] += 10
			if lenl-i-2 >= 0 && lst[lenl-i-2] != byte('.') {
				array1[lenl-i-2] -= 1
			} else {
				if lenl-i-3 >= 0 {
					array1[lenl-i-3] -= 1
				} else {
					fmt.Printf("something wrong\n")
				}
			}
		}
	}
	ans := ""
	for i := 0; i < lenl; i++ {
		ans += string(array1[i])
	}
	ans = negative + ans
	return adjustNumber(ans)
}
func sub(lst, rst string) string { // deal with differ situation
	if isvalid(lst) == false || isvalid(rst) == false {
		fmt.Println("invalid number format")
		return "nil"
	}
	lst = adjustNumber(lst)
	rst = adjustNumber(rst)
	if lst[0] == '-' && rst[0] == '-' {
		return subb(rst[1:], lst[1:])
	} else if lst[0] != '-' && rst[0] == '-' {
		return add(rst[1:], lst)
	} else if lst[0] != '-' && rst[0] != '-' {
		return subb(lst, rst)
	} else { //if lst[0] == '-' && rst[0] != '-'
		return "-" + add(rst, lst[1:])
	}
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func getInput() (string, string, string) {
	var err error
	reader := bufio.NewReader(os.Stdin)
	var s string
	for {
		s, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("input error")
			continue
		}
		s = strings.Replace(s, " ", "", -1)
		s = strings.Replace(s, "(", "", -1)
		s = strings.Replace(s, ")", "", -1)
		//	s = strings.Replace(s, "\n", "", -1)
		s = strings.TrimSpace(s)
		//		fmt.Printf("after :%s %d\n", s, len(s))
		firstSymbol := strings.Index(s, "+")
		if tmp := strings.Index(s, "-"); tmp != -1 {
			firstSymbol = min(firstSymbol, tmp)
		}

		if firstSymbol == -1 {
			return s, "", ""
		} else if firstSymbol == 0 {
			ss := s[1:]
			secondSymbol := strings.Index(ss, "+")
			if tmp := strings.Index(ss, "-"); tmp != -1 {
				secondSymbol = min(secondSymbol, tmp)
			}
			if secondSymbol == -1 {
				return s, "=", ""
			} else {
				return s[0 : secondSymbol+1], string(ss[secondSymbol]), ss[secondSymbol+1:]
			}
		} else {
			return s[0:firstSymbol], s[firstSymbol : firstSymbol+1], s[firstSymbol+1 : len(s)]
		}

	}

}
func main() {
	fmt.Println("This program Only support the addition and subtraction of real numbers")
	fmt.Println("Input Q/q to exit")
	for {
		fmt.Println("please input a expression")
		var num1, op, num2 string
		//fmt.Scanf("%s %s %s\n", &num1, &op, &num2)

		num1, op, num2 = getInput()
		//	fmt.Printf("getresult :%d %s %d %s %d %s\n", len(num1), num1, len(op), op, len(num2), num2)
		if num1 == "q" || num1 == "Q" {
			break
		}
		if op == "+" {
			fmt.Printf("result = %s\n", add(num1, num2))
		} else if op == "-" {
			fmt.Printf("result = %s\n", sub(num1, num2))
		} else if op == "=" {
			fmt.Println("result = %s\n", num1)
		} else {
			fmt.Println("Invalid operation")
		}
	}
}
