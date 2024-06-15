package fizzbuzz

import (
	"strconv"
)

// FizzBuzz関数
// 3の倍数の場合はFizzを返す
// 5の倍数の場合はBuzzを返す
// 3の倍数かつ5の倍数の場合はFizzBuzzを返す
func FizzBuzz(n int) string {
	if n%3 == 0 || n%5 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	} else {
		return strconv.Itoa(n)
	}
}
