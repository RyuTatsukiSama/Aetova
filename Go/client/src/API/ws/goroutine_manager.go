package ws

import (
	"context"
	"net/http"
)

var (
	ctx       context.Context
	cancel    context.CancelFunc
	chanClose chan bool
)

func initGoroutineVar(r *http.Request) {

	ctx, cancel = context.WithCancel(r.Context())
	chanClose = make(chan bool)
}
