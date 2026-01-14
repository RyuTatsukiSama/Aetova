package client_api

import (
	"net/http"
)

func HandlePostCancel(w http.ResponseWriter, r *http.Request) {
	cancel()
	// delete all downloaded file
}
