package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"net"
	"os"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, err
}

func FindFreePort(start int) (int, error) {

	for ok := true; ok; {
		addr := fmt.Sprintf(":%d", start)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			listener.Close()
			return start, nil
		}
	}
	return 0, fmt.Errorf("there was an error")
}
