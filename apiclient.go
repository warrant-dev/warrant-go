package warrant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ClientVersion string = "4.0.0"
)

type ApiClient struct {
	HttpClient *http.Client
	Config     ClientConfig
}

func NewApiClient(config ClientConfig) *ApiClient {
	httpClient := config.HttpClient
	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	return &ApiClient{
		HttpClient: httpClient,
		Config:     config,
	}
}

func (client ApiClient) MakeRequest(method string, path string, payload interface{}, options *RequestOptions) (*http.Response, error) {
	url := client.Config.ApiEndpoint + path
	if payload == nil {
		req, err := http.NewRequest(method, url, nil)
		if err != nil {
			return nil, WrapError("Unable to create request", err)
		}
		if client.Config.ApiKey != "" {
			req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", client.Config.ApiKey))
		}
		if options.WarrantToken != "" {
			req.Header.Add("Warrant-Token", options.WarrantToken)
		}
		req.Header.Add("User-Agent", fmt.Sprintf("warrant-go/%s", ClientVersion))
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
	if client.Config.ApiKey != "" {
		req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", client.Config.ApiKey))
	}
	if options.WarrantToken != "" {
		req.Header.Add("Warrant-Token", options.WarrantToken)
	}
	req.Header.Add("User-Agent", fmt.Sprintf("warrant-go/%s", ClientVersion))
	resp, err := client.HttpClient.Do(req)
	if err != nil {
		return nil, WrapError("Error making request", err)
	}

	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := io.ReadAll(resp.Body)
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}

	return resp, nil
}
