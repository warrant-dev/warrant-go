package warrant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const (
	ClientVersion string = "6.1.1"
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
	var request *http.Request
	var err error

	url := client.Config.ApiEndpoint + path
	if payload != nil {
		postBody, err := json.Marshal(payload)
		if err != nil {
			return nil, WrapError("Invalid request payload", err)
		}
		requestBody := bytes.NewBuffer(postBody)
		request, err = http.NewRequest(method, url, requestBody)
		if err != nil {
			return nil, WrapError("Unable to create request", err)
		}
	} else {
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			return nil, WrapError("Unable to create request", err)
		}
	}

	if client.Config.ApiKey != "" {
		request.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", client.Config.ApiKey))
	}
	if options != nil && options.WarrantToken != "" {
		request.Header.Add("Warrant-Token", options.WarrantToken)
	}
	request.Header.Add("User-Agent", fmt.Sprintf("warrant-go/%s", ClientVersion))

	resp, err := client.HttpClient.Do(request)
	if err != nil {
		return nil, WrapError("Error making request", err)
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, err := io.ReadAll(resp.Body)
		errMsg := ""
		if err == nil {
			errMsg = string(msg)
		}
		return nil, Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, errMsg),
		}
	}

	return resp, nil
}
