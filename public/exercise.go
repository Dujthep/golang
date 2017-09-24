package public

import (
	"strings"
)

func Multiple(a, b int) int {
	return a * b
}

func Stepup(a, b int) (int,int) {
	return b,b+a
}

func Factorial(a int) int {
	j := 1
	for i := 1; i <= a; i++ {
		j *= i
	}
	return j
}

func Factorial2(a int) int {
	return recursive(a, a-1)
}

func recursive(result, a int) int {
	if(a == 0) {
		return result
	}
	return Factorial2(result)
}

func Concat(a, b []string) []string{
	return append(a, b...)
}

func DelFirst(a []string) []string {
	return a[1:]
}

func DelLast(a []string) []string {
	return a[:len(a)-1]
}

func DelSecond(a []string) []string {
	result := []string {}
	result = append(result, a[0])
	result = append(result, a[2:]...)
	return result
}

func GetOdd(a []int) []int { 
	result := []int {}
	for i := 0; i < len(a); i++ {
		if(a[i] % 2 == 1) {
			result = append(result, a[i] )
		}
	}
	return result
}

func LettersCount(text string) map[string]int {
	arr := strings.Split(text, " ")
	mapping := map[string]int {}
	for i:=0; i < len(arr); i++ {
		mapping[arr[i]] = len(arr[i])
	}
	return mapping
}

func enc(xx, rot byte) byte {
	if xx == byte(32) {
		return xx - 97
	}
	return (xx + rot) % 26
}

func dec(text, shift byte) byte {
	if text == byte(32) {
		return text + 97
	}
	return (text + shift) % 26
}

func Caesar(s string, rot byte) string {
	result := []string{}
	for i:= 0; i < len(s); i++ {
		result = append(result, string(enc(s[i], rot) + 97))
	}
	return strings.Join(result,"")
}

func DecCaesar(s string, rot byte) string {
	result := []string{}
	for i:= 0; i < len(s); i++ {
		result = append(result, string(dec(s[i], rot) + 97))
	}
	return strings.Join(result,"")
}