package main

import (
	client_api "aetova/client/src/API"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func main() {
	opts := docLogger.NewOptionsBuilder().Build()
	docLogger.SetGlobalLoggerOptions(opts)
	dLog := docLogger.NewLogger("Client/main", *opts)

	if err := os.Chdir("wd"); err != nil {
		dLog.Log(docLogger.Error, err.Error())
		return
	}

	dLog.Log(docLogger.Error, client_api.LaunchAPI().Error())
}
