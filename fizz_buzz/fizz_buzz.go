package fizz_buzz

import (
	"strconv"
)

func CallFizzBuzz(a int) string {
	var result string
	if a % 3 == 0 {
		result += "fizz"
	}  
	if a % 5 == 0 {
		result += "buzz"
	} 
	if result == ""{
		result = strconv.Itoa(a)
	}
	return result
}