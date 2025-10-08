package variables

import (
	"fmt"
)

func Variables_refactored() {
	// Declare and assign with auto equivalent
	a1 := 42
	b1, c1 := true, 32.0
	d1 := "string"
	fmt.Println(a1, b1, c1, d1)
}
