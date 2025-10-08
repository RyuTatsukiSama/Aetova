package variables

import (
	"fmt"
)

var a int

var (
	b bool
	c float32
	d string
)

func Variables_main() {
	a = 42
	b, c = true, 32.0
	d = "string"
	fmt.Println(a, b, c, d)
}
