package role

import (
	"fmt"

	"github.com/warrant-dev/warrant-go/v5"
	"github.com/warrant-dev/warrant-go/v5/object"
)

type Client struct {
	apiClient *warrant.ApiClient
}

func NewClient(config warrant.ClientConfig) Client {
	return Client{
		apiClient: warrant.NewApiClient(config),
	}
}

func (c Client) Create(params *warrant.RoleParams) (*warrant.Role, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeRole,
		RequestOptions: params.RequestOptions,
	}
	if params.RoleId != "" {
		objectParams.ObjectId = params.RoleId
	}
	if params.Meta != nil {
		objectParams.Meta = params.Meta
	}
	object, err := object.Create(&objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Role{
		RoleId: object.ObjectId,
		Meta:   object.Meta,
	}, nil
}

func Create(params *warrant.RoleParams) (*warrant.Role, error) {
	return getClient().Create(params)
}

func (c Client) Get(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeRole,
		ObjectId:       roleId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Get(warrant.ObjectTypeRole, roleId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Role{
		RoleId: object.ObjectId,
		Meta:   object.Meta,
	}, nil
}

func Get(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	return getClient().Get(roleId, params)
}

func (c Client) Update(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypeRole,
		ObjectId:       roleId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Update(warrant.ObjectTypeRole, roleId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Role{
		RoleId: object.ObjectId,
		Meta:   object.Meta,
	}, nil
}

func Update(roleId string, params *warrant.RoleParams) (*warrant.Role, error) {
	return getClient().Update(roleId, params)
}

func (c Client) Delete(roleId string) (string, error) {
	return object.Delete(warrant.ObjectTypeRole, roleId)
}

func Delete(roleId string) (string, error) {
	return getClient().Delete(roleId)
}

func (c Client) ListRoles(listParams *warrant.ListRoleParams) (warrant.ListResponse[warrant.Role], error) {
	var rolesListResponse warrant.ListResponse[warrant.Role]

	objectsListResponse, err := object.ListObjects(&warrant.ListObjectParams{
		ListParams: listParams.ListParams,
		ObjectType: warrant.ObjectTypeRole,
	})
	if err != nil {
		return rolesListResponse, err
	}

	roles := make([]warrant.Role, 0)
	for _, object := range objectsListResponse.Results {
		roles = append(roles, warrant.Role{
			RoleId: object.ObjectId,
			Meta:   object.Meta,
		})
	}

	rolesListResponse = warrant.ListResponse[warrant.Role]{
		Results:    roles,
		PrevCursor: objectsListResponse.PrevCursor,
		NextCursor: objectsListResponse.NextCursor,
	}

	return rolesListResponse, nil
}

func ListRoles(listParams *warrant.ListRoleParams) (warrant.ListResponse[warrant.Role], error) {
	return getClient().ListRoles(listParams)
}

func (c Client) ListRolesForUser(userId string, listParams *warrant.ListRoleParams) (warrant.ListResponse[warrant.Role], error) {
	var rolesListResponse warrant.ListResponse[warrant.Role]

	queryResponse, err := warrant.Query(fmt.Sprintf("select role where user:%s is *", userId), &warrant.QueryParams{
		ListParams: listParams.ListParams,
	})
	if err != nil {
		return rolesListResponse, err
	}

	users := make([]warrant.Role, 0)
	for _, queryResult := range queryResponse.Results {
		users = append(users, warrant.Role{
			RoleId: queryResult.ObjectId,
			Meta:   queryResult.Meta,
		})
	}

	rolesListResponse = warrant.ListResponse[warrant.Role]{
		Results:    users,
		PrevCursor: queryResponse.PrevCursor,
		NextCursor: queryResponse.NextCursor,
	}

	return rolesListResponse, nil
}

func ListRolesForUser(userId string, listParams *warrant.ListRoleParams) (warrant.ListResponse[warrant.Role], error) {
	return getClient().ListRolesForUser(userId, listParams)
}

func (c Client) AssignRoleToUser(roleId string, userId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.apiClient.Config).Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeRole,
		ObjectId:   roleId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func AssignRoleToUser(roleId string, userId string) (*warrant.Warrant, error) {
	return getClient().AssignRoleToUser(roleId, userId)
}

func (c Client) RemoveRoleFromUser(roleId string, userId string) (string, error) {
	return warrant.NewClient(c.apiClient.Config).Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypeRole,
		ObjectId:   roleId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func RemoveRoleFromUser(roleId string, userId string) (string, error) {
	return getClient().RemoveRoleFromUser(roleId, userId)
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
