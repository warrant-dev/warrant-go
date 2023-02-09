package warrant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type WarrantClient struct {
	HttpClient *http.Client
	Config     ClientConfig
}

func (client WarrantClient) MakeRequest(method string, path string, payload interface{}) (*http.Response, error) {
	url := ApiEndpoint + path
	if payload == nil {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, WrapError("Unable to create request", err)
		}
		req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", client.Config.ApiKey))
		resp, err := client.HttpClient.Do(req)
		if err != nil {
			return nil, WrapError("Error making request", err)
		}
		return resp, nil
	}

	postBody, err := json.Marshal(payload)
	if err != nil {
		return nil, WrapError("Invalid request payload", err)
	}
	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, WrapError("Unable to create request", err)
	}
	req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", client.Config.ApiKey))
	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, WrapError("Error making request", err)
	}
	return resp, nil
}
