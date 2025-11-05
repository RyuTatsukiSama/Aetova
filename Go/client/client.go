package main

import (
	assemble "aetova/client/src"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	var err error

	if err = os.Chdir("wd"); err != nil {
		log.Fatal(err)
	}

	if err = dlNass(); err != nil {
		log.Fatal(err)
	}

	//http.HandleFunc("/health", handleHealth)

	//http.ListenAndServe(":8090", nil)
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintf(w, "Health check : ok")
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// download and assemble
func dlNass() (err error) {
	// download

	// assemble
	if err = assemble.AssembleGame("manifest.json"); err != nil {
		return err
	}

	return err

	// copy into the app folder
}
