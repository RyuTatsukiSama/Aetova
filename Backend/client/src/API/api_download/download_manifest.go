package api_download

import (
	"aetova/client/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
)

var (
	client *http.Client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxDLWorkers,
		},
	}
)

func GetManifest(cm chan utils.ManifestDir, ce chan error) {
	guid := 0    // TODO : Change this to avoid hard coded guid
	version := 0 // TODO : Change this to avoid hard coded guid

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/manifest?guid=%d&version=%d", Server_Url, guid, version), nil) // TODO : Change that to avoid hard code port for server URL
	if err != nil {
		ce <- err
		return
	}

	req.Header.Add("api_key", "c7e642cc-9928-4248-bd3f-c9588490bb60")

	resp, err := client.Do(req)
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
