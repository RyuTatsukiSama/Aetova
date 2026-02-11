package ws

import (
	mc "aetova/client/src/API/MutexConnection"
	"context"
	"net/http"
)

var (
	ctx       context.Context
	chanClose chan bool
)

func initGoroutineVar(r *http.Request) {

	ctx, mc.CancelFunc = context.WithCancel(r.Context())
	chanClose = make(chan bool)

}
