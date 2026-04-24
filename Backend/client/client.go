package main

import (
	client_api "aetova/client/src/API"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func folderForger() error {
	err := os.MkdirAll("Shipyard", 0700)
	if err != nil {
		return err
	}

	if err := os.Chdir("Shipyard"); err != nil {
		return err
	}

	err = os.MkdirAll("app", 0700)
	if err != nil {
		return err
	}

	err = os.MkdirAll("downloads", 0700)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	opts := docLogger.NewOptionsBuilder().Build()
	docLogger.SetGlobalLoggerOptions(opts)
	dLog := docLogger.NewLogger("Client/main", *opts)

	err := folderForger()
	if err != nil {
		dLog.Error("Error 13: " + err.Error())
	}

	dLog.Log(docLogger.Error, client_api.LaunchAPI().Error())
}
