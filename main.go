package main

import (
	"fmt"
)

import "github.com/mikcheal101/data-structures-in-go/fibonacciseries"

func main () {
	// create a fibonicci
	items := fibonacciseries.FibonacciSeries{}
	items.FibonacciConstructor(3)
	fmt.Println(items.FibonacciNumber)
}