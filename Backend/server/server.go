package main

import (
	server_api "aetova/server/src/API"
	butcher "aetova/server/src/Butcher"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func main() {
	var err error = nil

	opts := docLogger.NewOptionsBuilder().Build()
	docLogger.SetGlobalLoggerOptions(opts)
	dLog := docLogger.NewLoggerWithGOpts("Server")

	if err = os.MkdirAll("Shipyard", 0700); err != nil {
		dLog.Log(docLogger.Critical, err.Error())
		return
	}

	if err = os.Chdir("Shipyard"); err != nil {
		dLog.Log(docLogger.Critical, err.Error())
		return
	}

	if err = butcher.ChopGame("BuildOranys.zip"); err != nil {
		dLog.Log(docLogger.Critical, err.Error())
		return
	}

	dLog.Log(docLogger.Critical, server_api.LaunchAPI().Error())
}
