package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int64

	fmt.Print("Enter a number: ")
	fmt.Scanln(&n)

	result := Sum(n)
	fmt.Println(result)
}

func Sum(n int64) string {
	// TODO: Finish this function
	var ans int64
	ans = 0
	var str string = ""
	var i int64 = 0
	if n%7 != 0 {
		for i = 0; i < n; i++ {
			if i%7 != 0 {
				ans += i
				str += strconv.FormatInt(i, 10)
				str += "+"
			}
		}
		ans += n
		str += strconv.FormatInt(n, 10)
		str += "="
		str += strconv.FormatInt(ans, 10)
	} else {
		for i = 0; i < n-1; i++ {
			if i%7 != 0 {
				ans += i
				str += strconv.FormatInt(i, 10)
				str += "+"
			}
		}
		ans += (n - 1)
		str += strconv.FormatInt(n-1, 10)
		str += "="
		str += strconv.FormatInt(ans, 10)
	}
	return str

}
