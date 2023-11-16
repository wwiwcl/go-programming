package main

import (
	"fmt"
	"math/big"
	"strconv"
	"syscall/js"
)

func CheckPrime(this js.Value, args []js.Value) interface{} {
	val := js.Global().Get("value").Get("value").String()
	fmt.Println("CheckPrime called with value: " + val)
	n, _ := strconv.Atoi(val)
	result := big.NewInt(int64(n)).ProbablyPrime(0)
	fmt.Println("CheckPrime result: " + strconv.FormatBool(result))
	return result
}

func registerCallbacks() {
	// TODO: Register the function CheckPrime
	js.Global().Set("CheckPrime", js.FuncOf(CheckPrime))
}

func main() {
	fmt.Println("Golang main function executed")
	registerCallbacks()

	//need block the main thread forever
	select {}
}