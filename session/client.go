package session

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	warrant "github.com/warrant-dev/warrant-go"
)

type Session struct {
	UserId string `json:"userId"`
	Token  string `json:"token"`
}

type SessionParams struct {
	Type     string `json:"type"`
	UserId   string `json:"userId"`
	TenantId string `json:"tenantId"`
	TTL      int64  `json:"ttl"`
}

func New(params *SessionParams) (*Session, error) {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: warrant.ApiKey,
	})

	resp, err := client.MakeRequest("POST", "/v1/sessions", params)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, warrant.Error{
			Message: fmt.Sprintf("HTTP %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, warrant.WrapError("Error reading response", err)
	}
	var newSession Session
	err = json.Unmarshal([]byte(body), &newSession)
	if err != nil {
		return nil, warrant.WrapError("Invalid response from server", err)
	}
	return &newSession, nil
}
