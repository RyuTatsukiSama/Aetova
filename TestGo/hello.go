package main

import (
	"fmt"

	"rsc.io/quote"
)

func hello() {
	var eight int = 8
	fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
	fmt.Println(eight)
}
