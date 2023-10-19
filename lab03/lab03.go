package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func Calculator(w http.ResponseWriter, r *http.Request) {
	// TODO: implement a calculator
	str := r.URL.String()
	//fmt.Println(str)
	if str == "/favicon.ico" {
		fmt.Fprintf(w, "Error!")
		return
	}
	arr := strings.Split(str, "/")
	if len(arr) != 4 {
		fmt.Fprintf(w, "Error!")
		return
	}
	//fmt.Printf("%s\n", arr[1])
	//fmt.Printf("%s\n", arr[2])
	operation := arr[1]
	num1 := arr[2]
	num2 := arr[3]
	if operation == "add" {
		ans := ""
		ans += num1
		ans += " + "
		ans += num2
		ans += " = "
		a, err := strconv.Atoi(num1)
		if err != nil {
			fmt.Fprintf(w, "Error!")
			return
		}
		b, err := strconv.Atoi(num2)
		if err != nil {
			fmt.Fprintf(w, "Error!")
			return
		}
		fmt.Fprintf(w, ans)
		fmt.Fprintf(w, strconv.Itoa(a+b))
	} else if operation == "sub" {
		ans := ""
		ans += num1
		ans += " - "
		ans += num2
		ans += " = "
		a, err := strconv.Atoi(num1)
		if err != nil {
			fmt.Fprintf(w, "Error!")
			return
		}
		b, err := strconv.Atoi(num2)
		if err != nil {
			fmt.Fprintf(w, "Error!")
			return
		}
		fmt.Fprintf(w, ans)
		fmt.Fprintf(w, strconv.Itoa(a-b))
	} else if operation == "mul" {
		ans := ""
		ans += num1
		ans += " * "
		ans += num2
		ans += " = "
		a, err := strconv.Atoi(num1)
		if err != nil {
			fmt.Fprintf(w, "Error!")
			return
		}
		b, err := strconv.Atoi(num2)
		if err != nil {
			fmt.Fprintf(w, "Error!")
			return
		}
		fmt.Fprintf(w, ans)
		fmt.Fprintf(w, strconv.Itoa(a*b))
	} else if operation == "div" {
		ans := ""
		ans += num1
		ans += " / "
		ans += num2
		ans += " = "
		a, err := strconv.Atoi(num1)
		if err != nil {
			fmt.Fprintf(w, "Error!")
			return
		}
		b, err := strconv.Atoi(num2)
		if err != nil {
			fmt.Fprintf(w, "Error!")
			return
		}
		if b == 0 {
			fmt.Fprintf(w, "Error!")
			return
		}
		fmt.Fprintf(w, ans)
		fmt.Fprintf(w, strconv.Itoa(a/b))
		ans = ""
		ans += ", reminder = "
		fmt.Fprintf(w, ans)
		fmt.Fprintf(w, strconv.Itoa(a-b*(a/b)))
	} else {
		fmt.Fprintf(w, "Error!")
		return
	}
}

func main() {
	http.HandleFunc("/", Calculator)
	log.Fatal(http.ListenAndServe(":8083", nil))
}
