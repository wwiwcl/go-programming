package main

import (
	"fmt"
	"math/big"
	"strconv"
	"syscall/js"
)

func CheckPrime(this js.Value, args []js.Value) interface{} {
	if len(args) != 1 {
		return "Invalid number of arguments. Expected 1 argument"
	}

	number := args[0].String()
	n, err := strconv.Atoi(number)
	if err != nil {
		return "Invalid input. Please provide a valid number"
	}

	bigInt := big.NewInt(int64(n))

	isPrime := bigInt.ProbablyPrime(20) // Adjust the Miller-Rabin iterations for higher certainty

	return isPrime
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