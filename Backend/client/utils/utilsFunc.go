package utils

import (
	"errors"
	"fmt"
	"io/fs"
	"net"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func Exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, fs.ErrNotExist) {
		return false, nil
	}
	return false, errors.New("Error 19: " + err.Error())
}

func FindFreePort(start int) (int, error) {
	for ok := true; ok; {
		addr := fmt.Sprintf(":%d", start)
		listener, err := net.Listen("tcp", addr)
		if err == nil {
			listener.Close()
			writeFilePort(start)
			return start, nil
		} else {
			start++
		}
	}
	return 0, fmt.Errorf("there was an error")

}

func writeFilePort(port int) {
	dLog := docLogger.NewLoggerWithGOpts("Client/port")

	f, err := os.Create(".conf")
	if err != nil {
		dLog.Error("Error 11: " + err.Error())
		return
	}
	defer f.Close()

	fmt.Fprintf(f, "%d\n", port)
}
