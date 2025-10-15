package beyond100lines

import "fmt"

func Pointers_main() {
	var address *int  // declare an int pointer
	number := 42      // int
	address = &number // address stores the memory address of number
	value := *address // dereferencing the value

	fmt.Printf("address: %v\n", address) // address: 0xc0000ae008
	fmt.Printf("value: %v\n", value)     // value: 42
}
