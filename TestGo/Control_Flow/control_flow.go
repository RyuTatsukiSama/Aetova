package ControlFlow

import "fmt"

func while_loop() {
	count := 1
	for count < 5 {
		count += count
	}
	fmt.Println(count)
}

func average_refactored(x []float64) (avg float64) {
	total := 0.0
	zero := 0
	/* No break needed, case can be variable, switch statement is not only an int  */
	switch len(x) {
	case zero:
		avg = 0
	default:
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))

	}
	return
}

func average(x []float64) (avg float64) {
	total := 0.0
	if len(x) == 0 {
		avg = 0
	} else {
		for _, v := range x {
			total += v
		}
		avg = total / float64(len(x))
	}
	return
}

func ControlFlow_main() {
	x := []float64{2.15, 3.14, 42.0, 29.5}
	fmt.Println(average(x))
	fmt.Println(average_refactored(x))
	while_loop()
}
