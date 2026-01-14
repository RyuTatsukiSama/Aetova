package ws

import (
	"context"
	"net/http"
)

var (
	ctx        context.Context
	cancelFunc context.CancelFunc
	chanClose  chan bool
)

func initGoroutineVar(r *http.Request) {

	ctx, cancelFunc = context.WithCancel(r.Context())
	chanClose = make(chan bool)
}
