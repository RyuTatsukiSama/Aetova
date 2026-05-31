package api_download

import (
	"aetova/client/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

var (
	client *http.Client = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxDLWorkers,
		},
	}
)

func GetManifest(cm chan utils.ManifestDir, ce chan error, guid uint, version uint) {
	// Create the request
	params := url.Values{}
	params.Add("guid", fmt.Sprintf("%d", guid))
	params.Add("mType", "app")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/manifest?%s", Server_Url, params.Encode()), nil) // TODO : Change that to avoid hard code port for server URL
	if err != nil {
		ce <- err
		return
	}
	req.Header.Add("api_key", "c7e642cc-9928-4248-bd3f-c9588490bb60")

	// Launch the request and check resp status code
	resp, err := client.Do(req)
	if err != nil {
		ce <- err
		return
	}
	if resp.StatusCode != 200 {
		ce <- errors.New(resp.Status)
		return
	}

	// Extract body data
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ce <- err
		return
	}

	// Unmarshall it
	var gameManifest utils.ManifestGame
	if err := json.Unmarshal(body, &gameManifest); err != nil {
		ce <- err
		return
	}

	// Close the body
	err = resp.Body.Close()
	if err != nil {
		ce <- err
	}

	// Save the app manifest into a file
	err = os.WriteFile(fmt.Sprintf("AppManifest_%d.json", gameManifest.Guid), body, 0777) // TODO: When PostgreSQL is here, change 0 by the guid in the manifest
	if err != nil {
		ce <- err
		return
	}

	// Get the dir manifest
	dirManifest, err := getDirManifest(gameManifest)
	if err != nil {
		ce <- err
		return
	}

	cm <- dirManifest
}

func getDirManifest(gameManifest utils.ManifestGame) (utils.ManifestDir, error) {
	// Create the request
	params := url.Values{}
	params.Add("guid", fmt.Sprintf("%d", gameManifest.Guid))
	params.Add("version", fmt.Sprintf("%d", gameManifest.Version))
	params.Add("mType", "dl")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/manifest?%s", Server_Url, params.Encode()), nil) // TODO : Change that to avoid hard code port for server URL
	if err != nil {
		return utils.ManifestDir{}, err
	}

	req.Header.Add("api_key", "c7e642cc-9928-4248-bd3f-c9588490bb60")

	// Do the request and check the status code respond
	resp, err := client.Do(req)
	if err != nil {
		return utils.ManifestDir{}, err
	}
	if resp.StatusCode != 200 {
		return utils.ManifestDir{}, errors.New(resp.Status)
	}

	// Extract data from the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return utils.ManifestDir{}, err
	}

	// Convert it to manifestDir
	var dirManifest utils.ManifestDir
	if err := json.Unmarshal(body, &dirManifest); err != nil {
		return utils.ManifestDir{}, err
	}

	// Close the body
	err = resp.Body.Close()
	if err != nil {
		return utils.ManifestDir{}, err
	}

	// save the manifest into a file
	err = os.WriteFile(fmt.Sprintf(TargetDl+"Mfs_%d.json", gameManifest.Guid), body, 0777) // TODO: When PostgreSQL is here, change 0 by the guid in the manifest
	if err != nil {
		return utils.ManifestDir{}, err
	}

	return dirManifest, nil
}

func GetUManifest(cm chan utils.ManifestUDir, ce chan error, guid uint, version uint) {
	// Create the request
	params := url.Values{}
	params.Add("guid", fmt.Sprintf("%d", guid))
	params.Add("version", fmt.Sprintf("%d", version))
	params.Add("mType", "upt")

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/manifest?%s", Server_Url, params.Encode()), nil) // TODO : Change that to avoid hard code port for server URL
	if err != nil {
		ce <- err
		return
	}

	req.Header.Add("api_key", "c7e642cc-9928-4248-bd3f-c9588490bb60")

	// Do the request and check the status code respond
	resp, err := client.Do(req)
	if err != nil {
		ce <- err
		return
	}
	if resp.StatusCode != 200 {
		ce <- errors.New(resp.Status)
		return
	}

	// Extract data from the body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		ce <- err
		return
	}

	// Convert it to manifestDir
	var dirManifest utils.ManifestUDir
	if err := json.Unmarshal(body, &dirManifest); err != nil {
		ce <- err
		return
	}

	// Close the body
	err = resp.Body.Close()
	if err != nil {
		ce <- err
		return
	}

	// save the manifest into a file
	err = os.WriteFile(fmt.Sprintf(TargetDl+"UMfs_%d.json", guid), body, 0777)
	if err != nil {
		ce <- err
		return
	}

	cm <- dirManifest
}
