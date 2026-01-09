package main

import (
	client_api "aetova/client/src/API"
	"context"
	"os"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func main() {

	opts := docLogger.NewOptionsBuilder().Build()
	docLogger.SetGlobalLoggerOptions(opts)
	dLog, ctx, _ := docLogger.NewLogger("Client/main", *opts, context.Background())

	if err := os.Chdir("wd"); err != nil {
		dLog.Log(docLogger.Error, err.Error())
		return
	}

	dLog.Log(docLogger.Error, client_api.LaunchAPI(ctx).Error())
}
