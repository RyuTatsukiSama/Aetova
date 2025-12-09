package api_download

import (
	"aetova/client/utils"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func GetManifest(cm chan utils.ManifestDir, ce chan error) {
	req, err := http.NewRequest("GET", "http://aetova.duckdns.org:15369"+"/manifest", nil) // TODO : Change that to avoid hard code port
	if err != nil {
		ce <- err
		return
	}

	req.Header.Add("api_key", "c7e642cc-9928-4248-bd3f-c9588490bb60")

	resp, err := Client.Do(req)
	if err != nil {
		ce <- err
		return
	}
	if resp.StatusCode != 200 {
		ce <- errors.New(resp.Status)
		return
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ce <- err
		return
	}

	var manifest utils.ManifestDir
	if err := json.Unmarshal(body, &manifest); err != nil {
		ce <- err
		return
	}

	err = resp.Body.Close()
	if err != nil {
		ce <- err
	}

	cm <- manifest
}
