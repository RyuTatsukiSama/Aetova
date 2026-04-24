package client_api

import (
	"aetova/client/src/API/api_download"
	"encoding/json"
	"io"
	"net/http"
)

type Health struct {
	Status       string
	Message      string
	ServerHealth map[string]string
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleGetHealth(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleGetHealth(w http.ResponseWriter, r *http.Request) {
	// launch the health check of the service

	s_Health := serverHealth()
	if s_Health["status"] != "ok" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(Health{
			Status:       "failed",
			Message:      "Server Health check failed",
			ServerHealth: s_Health,
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Health{
		Status:       "ok",
		Message:      "Client is healthy",
		ServerHealth: s_Health,
	})
}

func serverHealth() map[string]string {
	req, err := http.NewRequest("GET", api_download.Server_Url+"/health", nil)
	if err != nil {
		return map[string]string{
			"status":         "failed",
			"message":        "Error creating Request to server",
			"Detailed Error": err.Error(),
		}
	}

	req.Header.Add("api_key", "c7e642cc-9928-4248-bd3f-c9588490bb60")

	resp, err := client.Do(req)
	if err != nil {
		return map[string]string{
			"status":         "failed",
			"message":        "Error with the request",
			"Detailed Error": err.Error(),
		}
	}
	if resp.StatusCode != 200 {
		return map[string]string{
			"status":         "failed",
			"message":        "Error with the request",
			"Detailed Error": resp.Status,
		}
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return map[string]string{
			"status":         "failed",
			"message":        "Error reading respond body",
			"Detailed Error": err.Error(),
		}
	}

	var healthCheck map[string]string
	if err := json.Unmarshal(body, &healthCheck); err != nil {
		return map[string]string{
			"status":         "failed",
			"message":        "Error decoding body",
			"Detailed Error": err.Error(),
		}
	}

	err = resp.Body.Close()
	if err != nil {
		return map[string]string{
			"status":         "failed",
			"message":        "Error closing body",
			"Detailed Error": err.Error(),
		}
	}

	return healthCheck
}
