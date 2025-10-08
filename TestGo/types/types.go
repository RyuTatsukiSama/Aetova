package docTypes

import "fmt"

func Type_main() {
	/* User specified types */
	const a rune = 12           // 32-bits integer (can also put int32)
	const b float32 = 20.5      // 32-bit float
	const c complex128 = 1 + 4i // 128-bits complex number
	const d uint16 = 14         // 16-bits unsigned int

	/* Default types */
	n := 42                      // int
	pi := 3.14159265358979323846 // float 64
	x, y := true, false          // bool
	z := "Go is awesome"         // string

	fmt.Printf("user-specified types:\n %T %T %T %T\n", a, b, c, d)
	fmt.Printf("default types:\n %T %T %T %T %T\n", n, pi, x, y, z)
}
