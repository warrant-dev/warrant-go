package user

import (
	"fmt"

	"github.com/warrant-dev/warrant-go/v6"
	"github.com/warrant-dev/warrant-go/v6/object"
)

type Client struct {
	apiClient *warrant.ApiClient
}

func NewClient(config warrant.ClientConfig) Client {
	return Client{
		apiClient: warrant.NewApiClient(config),
	}
}

func (c Client) Create(params *warrant.UserParams) (*warrant.User, error) {
	if params == nil {
		params = &warrant.UserParams{}
	}
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeUser,
		RequestOptions: params.RequestOptions,
	}
	if params.UserId != "" {
		objectParams.ObjectId = params.UserId
	}
	if params.Meta != nil {
		objectParams.Meta = params.Meta
	}
	object, err := object.Create(&objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.User{
		UserId: object.ObjectId,
		Meta:   object.Meta,
	}, nil
}

func Create(params *warrant.UserParams) (*warrant.User, error) {
	return getClient().Create(params)
}

func (c Client) BatchCreate(params []warrant.UserParams) ([]warrant.User, error) {
	objectsToCreate := make([]warrant.ObjectParams, 0)
	for _, userParam := range params {
		objectsToCreate = append(objectsToCreate, warrant.ObjectParams{
			RequestOptions: userParam.RequestOptions,
			ObjectType:     warrant.ObjectTypeUser,
			ObjectId:       userParam.UserId,
			Meta:           userParam.Meta,
		})
	}

	createdObjects, err := object.BatchCreate(objectsToCreate)
	if err != nil {
		return nil, err
	}

	users := make([]warrant.User, 0)
	for _, createdObject := range createdObjects {
		users = append(users, warrant.User{
			UserId: createdObject.ObjectId,
			Meta:   createdObject.Meta,
		})
	}

	return users, nil
}

func BatchCreate(params []warrant.UserParams) ([]warrant.User, error) {
	return getClient().BatchCreate(params)
}

func (c Client) Get(userId string, params *warrant.UserParams) (*warrant.User, error) {
	if params == nil {
		params = &warrant.UserParams{}
	}
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeUser,
		ObjectId:       userId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Get(warrant.ObjectTypeUser, userId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.User{
		UserId: object.ObjectId,
		Meta:   object.Meta,
	}, nil
}

func Get(userId string, params *warrant.UserParams) (*warrant.User, error) {
	return getClient().Get(userId, params)
}

func (c Client) Update(userId string, params *warrant.UserParams) (*warrant.User, error) {
	if params == nil {
		params = &warrant.UserParams{}
	}
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeUser,
		ObjectId:       userId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Update(warrant.ObjectTypeUser, userId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.User{
		UserId: object.ObjectId,
		Meta:   object.Meta,
	}, nil
}

func Update(userId string, params *warrant.UserParams) (*warrant.User, error) {
	return getClient().Update(userId, params)
}

func (c Client) Delete(userId string) (string, error) {
	return object.Delete(warrant.ObjectTypeUser, userId)
}

func Delete(userId string) (string, error) {
	return getClient().Delete(userId)
}

func (c Client) BatchDelete(params []warrant.UserParams) (string, error) {
	objectsToDelete := make([]warrant.ObjectParams, 0)
	for _, userParam := range params {
		objectsToDelete = append(objectsToDelete, warrant.ObjectParams{
			RequestOptions: userParam.RequestOptions,
			ObjectType:     warrant.ObjectTypeUser,
			ObjectId:       userParam.UserId,
			Meta:           userParam.Meta,
		})
	}

	warrantToken, err := object.BatchDelete(objectsToDelete)
	if err != nil {
		return "", err
	}

	return warrantToken, nil
}

func BatchDelete(params []warrant.UserParams) (string, error) {
	return getClient().BatchDelete(params)
}

func (c Client) ListUsers(listParams *warrant.ListUserParams) (warrant.ListResponse[warrant.User], error) {
	if listParams == nil {
		listParams = &warrant.ListUserParams{}
	}
	var usersListResponse warrant.ListResponse[warrant.User]

	objectsListResponse, err := object.ListObjects(&warrant.ListObjectParams{
		ListParams: listParams.ListParams,
		ObjectType: warrant.ObjectTypeUser,
	})
	if err != nil {
		return usersListResponse, err
	}

	users := make([]warrant.User, 0)
	for _, object := range objectsListResponse.Results {
		users = append(users, warrant.User{
			UserId: object.ObjectId,
			Meta:   object.Meta,
		})
	}

	usersListResponse = warrant.ListResponse[warrant.User]{
		Results:    users,
		PrevCursor: objectsListResponse.PrevCursor,
		NextCursor: objectsListResponse.NextCursor,
	}

	return usersListResponse, nil
}

func ListUsers(listParams *warrant.ListUserParams) (warrant.ListResponse[warrant.User], error) {
	return getClient().ListUsers(listParams)
}

func (c Client) ListUsersForTenant(tenantId string, listParams *warrant.ListUserParams) (warrant.ListResponse[warrant.User], error) {
	if listParams == nil {
		listParams = &warrant.ListUserParams{}
	}
	var usersListResponse warrant.ListResponse[warrant.User]

	queryResponse, err := warrant.Query(fmt.Sprintf("select * of type user for tenant:%s", tenantId), &warrant.QueryParams{
		ListParams: listParams.ListParams,
	})
	if err != nil {
		return usersListResponse, err
	}

	users := make([]warrant.User, 0)
	for _, queryResult := range queryResponse.Results {
		users = append(users, warrant.User{
			UserId: queryResult.ObjectId,
			Meta:   queryResult.Meta,
		})
	}

	usersListResponse = warrant.ListResponse[warrant.User]{
		Results:    users,
		PrevCursor: queryResponse.PrevCursor,
		NextCursor: queryResponse.NextCursor,
	}

	return usersListResponse, nil
}

func ListUsersForTenant(tenantId string, listParams *warrant.ListUserParams) (warrant.ListResponse[warrant.User], error) {
	return getClient().ListUsersForTenant(tenantId, listParams)
}

func (c Client) AssignUserToTenant(userId string, tenantId string, role string) (*warrant.Warrant, error) {
	return warrant.Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeTenant,
		ObjectId:   tenantId,
		Relation:   role,
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func AssignUserToTenant(userId string, tenantId string, role string) (*warrant.Warrant, error) {
	return getClient().AssignUserToTenant(userId, tenantId, role)
}

func (c Client) RemoveUserFromTenant(userId string, tenantId string, role string) (string, error) {
	return warrant.Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeTenant,
		ObjectId:   tenantId,
		Relation:   role,
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func RemoveUserFromTenant(userId string, tenantId string, role string) (string, error) {
	return getClient().RemoveUserFromTenant(userId, tenantId, role)
}

func getClient() Client {
	config := warrant.ClientConfig{
		ApiKey:                  warrant.ApiKey,
		ApiEndpoint:             warrant.ApiEndpoint,
		AuthorizeEndpoint:       warrant.AuthorizeEndpoint,
		SelfServiceDashEndpoint: warrant.SelfServiceDashEndpoint,
		HttpClient:              warrant.HttpClient,
	}

	return Client{
		&warrant.ApiClient{
			HttpClient: warrant.HttpClient,
			Config:     config,
		},
	}
}
