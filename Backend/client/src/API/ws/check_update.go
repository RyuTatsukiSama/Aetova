package ws

import (
	mc "aetova/client/src/API/MutexConnection"
	"aetova/client/src/API/api_download"
	"aetova/client/utils"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/RyuTatsukiSama/docLogger/go/docLogger"
)

func CheckUpdate(mxConn mc.MutexConnection) error {
	dLog := docLogger.NewLoggerWithGOpts("client/CheckUpdate")

	client := &http.Client{}

	entries, err := os.ReadDir("./")
	if err != nil {
		return err
	}

	// Get all the app manifest present
	var manifests []string
	for _, entry := range entries {
		name := entry.Name()
		if strings.HasPrefix(name, "AppManifest_") && strings.HasSuffix(name, ".json") {
			manifests = append(manifests, name)
		}
	}

	// Load all the manifest and check if a update is need
	for _, mfs := range manifests {
		file, err := os.Open(mfs)
		if err != nil {
			return err
		}

		var gManifest utils.ManifestGame
		err = utils.FromJson(&gManifest, file)
		if err != nil {
			return err
		}

		params := url.Values{}
		params.Add("guid", fmt.Sprintf("%d", gManifest.Guid))

		req, err := http.NewRequest("GET", fmt.Sprintf("%s/version?%s", api_download.Server_Url, params.Encode()), nil) // TODO : Change that to avoid hard code port for server URL
		if err != nil {
			return err
		}

		req.Header.Add("api_key", "c7e642cc-9928-4248-bd3f-c9588490bb60")

		// Do the request and check the status code respond
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		if resp.StatusCode != 200 {
			return errors.New(resp.Status)
		}

		// Extract data from the body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var serverVersion uint
		err = json.Unmarshal(body, &serverVersion)
		if err != nil {
			return err
		}

		if gManifest.Version != serverVersion {
			dLog.Info(mfs + " need update")
			mxConn.WriteJSON(gManifest.Guid, mc.NeedUpdate)
		}
	}

	return nil
}
