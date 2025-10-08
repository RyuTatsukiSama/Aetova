package utils

import (
	"bufio"
	"fmt"
	"os"
)

func Pause() {
	fmt.Print("Press 'Enter' to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
