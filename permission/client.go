package permission

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

func (c Client) Create(params *warrant.PermissionParams) (*warrant.Permission, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypePermission,
		RequestOptions: params.RequestOptions,
	}
	if params.PermissionId != "" {
		objectParams.ObjectId = params.PermissionId
	}
	if params.Meta != nil {
		objectParams.Meta = params.Meta
	}
	object, err := object.Create(&objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Permission{
		PermissionId: object.ObjectId,
		Meta:         object.Meta,
	}, nil
}

func Create(params *warrant.PermissionParams) (*warrant.Permission, error) {
	return getClient().Create(params)
}

func (c Client) Get(permissionId string, params *warrant.PermissionParams) (*warrant.Permission, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypePermission,
		ObjectId:       permissionId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Get(warrant.ObjectTypePermission, permissionId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Permission{
		PermissionId: object.ObjectId,
		Meta:         object.Meta,
	}, nil
}

func Get(permissionId string, params *warrant.PermissionParams) (*warrant.Permission, error) {
	return getClient().Get(permissionId, params)
}

func (c Client) Update(permissionId string, params *warrant.PermissionParams) (*warrant.Permission, error) {
	objectParams := warrant.ObjectParams{
		ObjectType:     warrant.ObjectTypePermission,
		ObjectId:       permissionId,
		RequestOptions: params.RequestOptions,
		Meta:           params.Meta,
	}
	object, err := object.Update(warrant.ObjectTypePermission, permissionId, &objectParams)
	if err != nil {
		return nil, err
	}
	return &warrant.Permission{
		PermissionId: object.ObjectId,
		Meta:         object.Meta,
	}, nil
}

func Update(permissionId string, params *warrant.PermissionParams) (*warrant.Permission, error) {
	return getClient().Update(permissionId, params)
}

func (c Client) Delete(permissionId string) error {
	return object.Delete(warrant.ObjectTypePermission, permissionId)
}

func Delete(permissionId string) error {
	return getClient().Delete(permissionId)
}

func (c Client) ListPermissions(listParams *warrant.ListPermissionParams) (warrant.ListResponse[warrant.Permission], error) {
	var permissionsListResponse warrant.ListResponse[warrant.Permission]

	objectsListResponse, err := object.ListObjects(&warrant.ListObjectParams{
		ListParams: listParams.ListParams,
		ObjectType: warrant.ObjectTypePermission,
	})
	if err != nil {
		return permissionsListResponse, err
	}

	permissions := make([]warrant.Permission, 0)
	for _, object := range objectsListResponse.Results {
		permissions = append(permissions, warrant.Permission{
			PermissionId: object.ObjectId,
			Meta:         object.Meta,
		})
	}

	permissionsListResponse = warrant.ListResponse[warrant.Permission]{
		Results:    permissions,
		PrevCursor: objectsListResponse.PrevCursor,
		NextCursor: objectsListResponse.NextCursor,
	}

	return permissionsListResponse, nil
}

func ListPermissions(listParams *warrant.ListPermissionParams) (warrant.ListResponse[warrant.Permission], error) {
	return getClient().ListPermissions(listParams)
}

func (c Client) ListPermissionsForRole(roleId string, listParams *warrant.ListPermissionParams) (warrant.ListResponse[warrant.Permission], error) {
	var permissionsListResponse warrant.ListResponse[warrant.Permission]

	queryResponse, err := warrant.Query(fmt.Sprintf("select permission where role:%s is *", roleId), &warrant.QueryParams{
		ListParams: listParams.ListParams,
	})
	if err != nil {
		return permissionsListResponse, err
	}

	permissions := make([]warrant.Permission, 0)
	for _, queryResult := range queryResponse.Results {
		permissions = append(permissions, warrant.Permission{
			PermissionId: queryResult.ObjectId,
			Meta:         queryResult.Meta,
		})
	}

	permissionsListResponse = warrant.ListResponse[warrant.Permission]{
		Results:    permissions,
		PrevCursor: queryResponse.PrevCursor,
		NextCursor: queryResponse.NextCursor,
	}

	return permissionsListResponse, nil
}

func ListPermissionsForRole(roleId string, listParams *warrant.ListPermissionParams) (warrant.ListResponse[warrant.Permission], error) {
	return getClient().ListPermissionsForRole(roleId, listParams)
}

func (c Client) AssignPermissionToRole(permissionId string, roleId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.apiClient.Config).Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePermission,
		ObjectId:   permissionId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeRole,
			ObjectId:   roleId,
		},
	})
}

func AssignPermissionToRole(permissionId string, roleId string) (*warrant.Warrant, error) {
	return getClient().AssignPermissionToRole(permissionId, roleId)
}

func (c Client) RemovePermissionFromRole(permissionId string, roleId string) error {
	return warrant.NewClient(c.apiClient.Config).Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePermission,
		ObjectId:   permissionId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeRole,
			ObjectId:   roleId,
		},
	})
}

func RemovePermissionFromRole(permissionId string, roleId string) error {
	return getClient().RemovePermissionFromRole(permissionId, roleId)
}

func (c Client) ListPermissionsForUser(userId string, listParams *warrant.ListPermissionParams) (warrant.ListResponse[warrant.Permission], error) {
	var permissionsListResponse warrant.ListResponse[warrant.Permission]

	queryResponse, err := warrant.Query(fmt.Sprintf("select permission where user:%s is *", userId), &warrant.QueryParams{
		ListParams: listParams.ListParams,
	})
	if err != nil {
		return permissionsListResponse, err
	}

	permissions := make([]warrant.Permission, 0)
	for _, queryResult := range queryResponse.Results {
		permissions = append(permissions, warrant.Permission{
			PermissionId: queryResult.ObjectId,
			Meta:         queryResult.Meta,
		})
	}

	permissionsListResponse = warrant.ListResponse[warrant.Permission]{
		Results:    permissions,
		PrevCursor: queryResponse.PrevCursor,
		NextCursor: queryResponse.NextCursor,
	}

	return permissionsListResponse, nil
}

func ListPermissionsForUser(userId string, listParams *warrant.ListPermissionParams) (warrant.ListResponse[warrant.Permission], error) {
	return getClient().ListPermissionsForUser(userId, listParams)
}

func (c Client) AssignPermissionToUser(permissionId string, userId string) (*warrant.Warrant, error) {
	return warrant.NewClient(c.apiClient.Config).Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePermission,
		ObjectId:   permissionId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func AssignPermissionToUser(permissionId string, userId string) (*warrant.Warrant, error) {
	return getClient().AssignPermissionToUser(permissionId, userId)
}

func (c Client) RemovePermissionFromUser(permissionId string, userId string) error {
	return warrant.NewClient(c.apiClient.Config).Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePermission,
		ObjectId:   permissionId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   userId,
		},
	})
}

func RemovePermissionFromUser(permissionId string, userId string) error {
	return getClient().RemovePermissionFromUser(permissionId, userId)
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
