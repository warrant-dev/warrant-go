package warrant

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const API_URL_BASE = "https://api.warrant.dev"
const API_VERSION = "/v1"

//const WARRANT_IGNORE_ID = "WARRANT_IGNORE_ID"

type ClientConfig struct {
	ApiKey string
}

type WarrantError struct {
	Message      string
	WrappedError error
}

func (err WarrantError) Error() string {
	if err.WrappedError != nil {
		return fmt.Sprintf("Warrant error: %s %s", err.Message, err.WrappedError.Error())
	}
	return fmt.Sprintf("Warrant error: %s", err.Message)
}

func wrapError(message string, err error) WarrantError {
	return WarrantError{
		Message:      message,
		WrappedError: err,
	}
}

type User struct {
	UserId string `json:"userId"`
}

type Warrant struct {
	ObjectType string      `json:"objectType"`
	ObjectId   string      `json:"objectId"`
	Relation   string      `json:"relation"`
	User       WarrantUser `json:"user"`
}

type WarrantUser struct {
	UserId string `json:"userId,omitempty"`
	*Userset
}

type Userset struct {
	ObjectType string `json:"objectType"`
	ObjectId   string `json:"objectId"`
	Relation   string `json:"relation"`
}

type WarrantClient struct {
	Client *http.Client
	Config ClientConfig
}

func NewClient(config ClientConfig) WarrantClient {
	return WarrantClient{
		Client: http.DefaultClient,
		Config: config,
	}
}

func (client WarrantClient) CreateUser(user User) (*User, error) {
	resp, err := client.makeRequest("POST", "/users", user)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		return nil, WarrantError{
			Message: fmt.Sprintf("Http %d", respStatus),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newUser User
	err = json.Unmarshal([]byte(body), &newUser)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newUser, nil
}

func (client WarrantClient) CreateUserWithGeneratedId() (*User, error) {
	return client.CreateUser(User{})
}

func (client WarrantClient) IsAuthorized(toCheck Warrant) (bool, error) {
	resp, err := client.makeRequest("POST", "/authorize", toCheck)
	if err != nil {
		return false, err
	}
	respStatus := resp.StatusCode
	if respStatus == 200 {
		return true, nil
	} else if respStatus == 401 {
		return false, nil
	}
	return false, WarrantError{
		Message: fmt.Sprintf("Http %d", respStatus),
	}
}

func (client WarrantClient) CreateWarrant(warrantToCreate Warrant) (*Warrant, error) {
	if warrantToCreate.User.UserId != "" && warrantToCreate.User.Userset != nil {
		return nil, WarrantError{
			Message: "Warrant cannot contain both a userId and userset",
		}
	}
	resp, err := client.makeRequest("POST", "/warrants", warrantToCreate)
	if err != nil {
		return nil, err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		msg, _ := ioutil.ReadAll(resp.Body)
		return nil, WarrantError{
			Message: fmt.Sprintf("Http %d %s", respStatus, string(msg)),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, wrapError("Error reading response", err)
	}
	var newWarrant Warrant
	err = json.Unmarshal([]byte(body), &newWarrant)
	if err != nil {
		return nil, wrapError("Invalid response from server", err)
	}
	return &newWarrant, nil
}

func (client WarrantClient) CreateSession(userId string) (string, error) {
	resp, err := client.makeRequest("POST", "/users/"+userId+"/sessions", make(map[string]string))
	if err != nil {
		return "", err
	}
	respStatus := resp.StatusCode
	if respStatus < 200 || respStatus >= 400 {
		return "", WarrantError{
			Message: fmt.Sprintf("Http %d", respStatus),
		}
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", wrapError("Error reading response", err)
	}
	var response map[string]string
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return "", wrapError("Invalid response from server", err)
	}
	return response["token"], nil
}

func (client WarrantClient) makeRequest(method string, requestUri string, payload interface{}) (*http.Response, error) {
	postBody, err := json.Marshal(payload)
	if err != nil {
		return nil, wrapError("Invalid request payload", err)
	}
	requestBody := bytes.NewBuffer(postBody)
	req, err := http.NewRequest(method, API_URL_BASE+API_VERSION+requestUri, requestBody)
	if err != nil {
		return nil, wrapError("Unable to create request", err)
	}
	req.Header.Add("Authorization", "ApiKey "+client.Config.ApiKey)
	resp, err := client.Client.Do(req)
	if err != nil {
		return nil, wrapError("Error making request", err)
	}
	return resp, nil
}
