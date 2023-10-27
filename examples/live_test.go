package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warrant-dev/warrant-go/v5"
	"github.com/warrant-dev/warrant-go/v5/feature"
	"github.com/warrant-dev/warrant-go/v5/object"
	"github.com/warrant-dev/warrant-go/v5/objecttype"
	"github.com/warrant-dev/warrant-go/v5/permission"
	"github.com/warrant-dev/warrant-go/v5/pricingtier"
	"github.com/warrant-dev/warrant-go/v5/role"
	"github.com/warrant-dev/warrant-go/v5/session"
	"github.com/warrant-dev/warrant-go/v5/tenant"
	"github.com/warrant-dev/warrant-go/v5/user"
)

func setup() {
	warrant.ApiKey = ""
	warrant.ApiEndpoint = "https://api.warrant.dev"
	warrant.AuthorizeEndpoint = "https://api.warrant.dev"
}

func TestCrudUsers(t *testing.T) {
	setup()
	assert := assert.New(t)

	user1, err := user.Create(&warrant.UserParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(user1.UserId)
	assert.Nil(user1.Meta)

	user2, err := user.Create(&warrant.UserParams{
		UserId: "some_id",
		Meta: map[string]interface{}{
			"email": "test@email.com",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fetchUserParams := &warrant.UserParams{}
	fetchUserParams.SetWarrantToken("latest")
	refetchedUser, err := user.Get(user2.UserId, fetchUserParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(user2.UserId, refetchedUser.UserId)
	assert.Equal(user2.Meta, refetchedUser.Meta)

	user2, err = user.Update("some_id", &warrant.UserParams{
		Meta: map[string]interface{}{
			"email": "updated@email.com",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	refetchedUser, err = user.Get("some_id", fetchUserParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("some_id", refetchedUser.UserId)
	assert.Equal(map[string]interface{}{"email": "updated@email.com"}, refetchedUser.Meta)

	usersList, err := user.ListUsers(&warrant.ListUserParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(2, len(usersList.Results))

	err = user.Delete(user1.UserId)
	if err != nil {
		t.Fatal(err)
	}
	err = user.Delete(user2.UserId)
	if err != nil {
		t.Fatal(err)
	}
	usersList, err = user.ListUsers(&warrant.ListUserParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(usersList.Results))
}

func TestCrudTenants(t *testing.T) {
	setup()
	assert := assert.New(t)

	tenant1, err := tenant.Create(&warrant.TenantParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(tenant1.TenantId)
	assert.Nil(tenant1.Meta)

	tenant2, err := tenant.Create(&warrant.TenantParams{
		TenantId: "some_tenant_id",
		Meta: map[string]interface{}{
			"name": "new_name",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fetchTenantParams := &warrant.TenantParams{}
	fetchTenantParams.SetWarrantToken("latest")
	refetchedTenant, err := tenant.Get(tenant2.TenantId, fetchTenantParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(tenant2.TenantId, refetchedTenant.TenantId)
	assert.Equal(tenant2.Meta, refetchedTenant.Meta)

	tenant2, err = tenant.Update("some_tenant_id", &warrant.TenantParams{
		Meta: map[string]interface{}{
			"name": "updated_name",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	refetchedTenant, err = tenant.Get("some_tenant_id", fetchTenantParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("some_tenant_id", refetchedTenant.TenantId)
	assert.Equal(map[string]interface{}{"name": "updated_name"}, refetchedTenant.Meta)

	tenantsList, err := tenant.ListTenants(&warrant.ListTenantParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(2, len(tenantsList.Results))

	err = tenant.Delete(tenant1.TenantId)
	if err != nil {
		t.Fatal(err)
	}
	err = tenant.Delete(tenant2.TenantId)
	if err != nil {
		t.Fatal(err)
	}
	tenantsList, err = tenant.ListTenants(&warrant.ListTenantParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(tenantsList.Results))
}

func TestCrudRoles(t *testing.T) {
	setup()
	assert := assert.New(t)

	adminRole, err := role.Create(&warrant.RoleParams{
		RoleId: "admin",
		Meta: map[string]interface{}{
			"name":        "Admin",
			"description": "The admin role",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("admin", adminRole.RoleId)
	assert.Equal(map[string]interface{}{
		"name":        "Admin",
		"description": "The admin role",
	}, adminRole.Meta)

	viewerRole, err := role.Create(&warrant.RoleParams{
		RoleId: "viewer",
		Meta: map[string]interface{}{
			"name":        "Viewer",
			"description": "The viewer role",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fetchRoleParams := &warrant.RoleParams{}
	fetchRoleParams.SetWarrantToken("latest")
	refetchedRole, err := role.Get(viewerRole.RoleId, fetchRoleParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(viewerRole.RoleId, refetchedRole.RoleId)
	assert.Equal(viewerRole.Meta, refetchedRole.Meta)

	viewerRole, err = role.Update("viewer", &warrant.RoleParams{
		Meta: map[string]interface{}{
			"name":        "Viewer Updated",
			"description": "Updated desc",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	refetchedRole, err = role.Get(viewerRole.RoleId, fetchRoleParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("viewer", refetchedRole.RoleId)
	assert.Equal(map[string]interface{}{
		"name":        "Viewer Updated",
		"description": "Updated desc",
	}, refetchedRole.Meta)

	rolesList, err := role.ListRoles(&warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(2, len(rolesList.Results))

	err = role.Delete(adminRole.RoleId)
	if err != nil {
		t.Fatal(err)
	}
	err = role.Delete(viewerRole.RoleId)
	if err != nil {
		t.Fatal(err)
	}
	rolesList, err = role.ListRoles(&warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(rolesList.Results))
}

func TestCrudPermissions(t *testing.T) {
	setup()
	assert := assert.New(t)

	permission1, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "perm1",
		Meta: map[string]interface{}{
			"name":        "Permission 1",
			"description": "Permission with id 1",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("perm1", permission1.PermissionId)
	assert.Equal(map[string]interface{}{
		"name":        "Permission 1",
		"description": "Permission with id 1",
	}, permission1.Meta)

	permission2, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "perm2",
		Meta: map[string]interface{}{
			"name":        "Permission 2",
			"description": "Permission with id 2",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("perm2", permission2.PermissionId)
	assert.Equal(map[string]interface{}{
		"name":        "Permission 2",
		"description": "Permission with id 2",
	}, permission2.Meta)

	permission2, err = permission.Update("perm2", &warrant.PermissionParams{
		Meta: map[string]interface{}{
			"name":        "Permission 2 Updated",
			"description": "Updated desc",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fetchPermissionParams := &warrant.PermissionParams{}
	fetchPermissionParams.SetWarrantToken("latest")
	refetchedPermission, err := permission.Get("perm2", fetchPermissionParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("perm2", refetchedPermission.PermissionId)
	assert.Equal(map[string]interface{}{
		"name":        "Permission 2 Updated",
		"description": "Updated desc",
	}, refetchedPermission.Meta)

	permissionsList, err := permission.ListPermissions(&warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(2, len(permissionsList.Results))

	err = permission.Delete(permission1.PermissionId)
	if err != nil {
		t.Fatal(err)
	}
	err = permission.Delete(permission2.PermissionId)
	if err != nil {
		t.Fatal(err)
	}
	permissionsList, err = permission.ListPermissions(&warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(permissionsList.Results))
}

func TestCrudFeatures(t *testing.T) {
	setup()
	assert := assert.New(t)

	feature1, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "new-feature",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("new-feature", feature1.FeatureId)

	feature2, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "feature-2",
		Meta: map[string]interface{}{
			"name": "Feature 2",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fetchFeatureParams := &warrant.FeatureParams{}
	fetchFeatureParams.SetWarrantToken("latest")
	refetchedFeature, err := feature.Get(feature2.FeatureId, fetchFeatureParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(feature2.FeatureId, refetchedFeature.FeatureId)
	assert.Equal(feature2.Meta, refetchedFeature.Meta)

	feature2, err = feature.Update(feature2.FeatureId, &warrant.FeatureParams{
		Meta: map[string]interface{}{
			"name": "Updated Feature 2",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	refetchedFeature, err = feature.Get(feature2.FeatureId, fetchFeatureParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("feature-2", refetchedFeature.FeatureId)
	assert.Equal(map[string]interface{}{"name": "Updated Feature 2"}, refetchedFeature.Meta)

	featuresList, err := feature.ListFeatures(&warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(2, len(featuresList.Results))

	err = feature.Delete(feature1.FeatureId)
	if err != nil {
		t.Fatal(err)
	}
	err = feature.Delete(feature2.FeatureId)
	if err != nil {
		t.Fatal(err)
	}
	featuresList, err = feature.ListFeatures(&warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(featuresList.Results))
}

func TestCrudPricingTiers(t *testing.T) {
	setup()
	assert := assert.New(t)

	tier1, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "new-tier1",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("new-tier1", tier1.PricingTierId)

	tier2, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "tier-2",
		Meta: map[string]interface{}{
			"name": "Tier 2",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fetchTierParams := &warrant.PricingTierParams{}
	fetchTierParams.SetWarrantToken("latest")
	refetchedTier, err := pricingtier.Get(tier2.PricingTierId, fetchTierParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(tier2.PricingTierId, refetchedTier.PricingTierId)
	assert.Equal(tier2.Meta, refetchedTier.Meta)

	tier2, err = pricingtier.Update(tier2.PricingTierId, &warrant.PricingTierParams{
		Meta: map[string]interface{}{
			"name": "Updated Tier 2",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("tier-2", tier2.PricingTierId)
	assert.Equal(map[string]interface{}{"name": "Updated Tier 2"}, tier2.Meta)

	tiersList, err := pricingtier.ListPricingTiers(&warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(2, len(tiersList.Results))

	err = pricingtier.Delete(tier1.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}
	err = pricingtier.Delete(tier2.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}
	tiersList, err = pricingtier.ListPricingTiers(&warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(tiersList.Results))
}

func TestBatchCreateUsersAndTenants(t *testing.T) {
	setup()
	assert := assert.New(t)

	createdUsers, err := user.BatchCreate([]warrant.UserParams{
		{UserId: "user-1"},
		{UserId: "user-2"},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(2, len(createdUsers))
	assert.Equal("user-1", createdUsers[0].UserId)
	assert.Equal("user-2", createdUsers[1].UserId)

	createdTenants, err := tenant.BatchCreate([]warrant.TenantParams{
		{TenantId: "tenant-1"},
		{TenantId: "tenant-2"},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(2, len(createdTenants))
	assert.Equal("tenant-1", createdTenants[0].TenantId)
	assert.Equal("tenant-2", createdTenants[1].TenantId)

	err = user.Delete("user-1")
	if err != nil {
		t.Fatal(err)
	}
	err = user.Delete("user-2")
	if err != nil {
		t.Fatal(err)
	}
	err = tenant.Delete("tenant-1")
	if err != nil {
		t.Fatal(err)
	}
	err = tenant.Delete("tenant-2")
	if err != nil {
		t.Fatal(err)
	}
}

func TestMultiTenancy(t *testing.T) {
	setup()
	assert := assert.New(t)

	// Create users
	user1, err := user.Create(&warrant.UserParams{})
	if err != nil {
		t.Fatal(err)
	}
	user2, err := user.Create(&warrant.UserParams{})
	if err != nil {
		t.Fatal(err)
	}

	// Create tenants
	tenant1, err := tenant.Create(&warrant.TenantParams{
		TenantId: "tenant-1",
		Meta: map[string]interface{}{
			"name": "Tenant 1",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	tenant2, err := tenant.Create(&warrant.TenantParams{
		TenantId: "tenant-2",
		Meta: map[string]interface{}{
			"name": "Tenant 2",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Assign user1 -> tenant1
	user1TenantsList, err := tenant.ListTenantsForUser(user1.UserId, &warrant.ListTenantParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(user1TenantsList.Results))
	tenant1UsersList, err := user.ListUsersForTenant(tenant1.TenantId, &warrant.ListUserParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(tenant1UsersList.Results))

	_, err = user.AssignUserToTenant(user1.UserId, tenant1.TenantId, "member")
	if err != nil {
		t.Fatal(err)
	}

	user1TenantsList, err = tenant.ListTenantsForUser(user1.UserId, &warrant.ListTenantParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(user1TenantsList.Results))
	assert.Equal("tenant-1", user1TenantsList.Results[0].TenantId)
	tenant1UsersList, err = user.ListUsersForTenant(tenant1.TenantId, &warrant.ListUserParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(tenant1UsersList.Results))
	assert.Equal(user1.UserId, tenant1UsersList.Results[0].UserId)

	err = user.RemoveUserFromTenant(user1.UserId, tenant1.TenantId, "member")
	if err != nil {
		t.Fatal(err)
	}

	// Clean up
	err = user.Delete(user1.UserId)
	if err != nil {
		t.Fatal(err)
	}
	err = user.Delete(user2.UserId)
	if err != nil {
		t.Fatal(err)
	}
	err = tenant.Delete(tenant1.TenantId)
	if err != nil {
		t.Fatal(err)
	}
	err = tenant.Delete(tenant2.TenantId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestRBAC(t *testing.T) {
	setup()
	assert := assert.New(t)

	// Create users
	adminUser, err := user.Create(&warrant.UserParams{})
	if err != nil {
		t.Fatal(err)
	}
	viewerUser, err := user.Create(&warrant.UserParams{})
	if err != nil {
		t.Fatal(err)
	}

	// Create roles
	adminRole, err := role.Create(&warrant.RoleParams{
		RoleId: "administrator",
		Meta: map[string]interface{}{
			"name":        "Administrator",
			"description": "The admin role",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	viewerRole, err := role.Create(&warrant.RoleParams{
		RoleId: "viewer",
		Meta: map[string]interface{}{
			"name":        "Viewer",
			"description": "The viewer role",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create permissions
	createPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "create-report",
		Meta: map[string]interface{}{
			"name":        "Create Report",
			"description": "Permission to create reports",
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	viewPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "view-report",
		Meta: map[string]interface{}{
			"name":        "View Report",
			"description": "Permission to view reports",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	// Admin user tests
	adminUserRolesList, err := role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(adminUserRolesList.Results))

	adminRolePermissionsList, err := permission.ListPermissionsForRole(adminRole.RoleId, &warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(adminRolePermissionsList.Results))

	adminUserHasPermission, err := warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		PermissionId: "create-report",
		UserId:       adminUser.UserId,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(adminUserHasPermission)

	// Assign create-report permission -> admin role -> admin user
	_, err = permission.AssignPermissionToRole(createPermission.PermissionId, adminRole.RoleId)
	if err != nil {
		t.Fatal(err)
	}

	_, err = role.AssignRoleToUser(adminRole.RoleId, adminUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	adminRolePermissionsList, err = permission.ListPermissionsForRole(adminRole.RoleId, &warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(adminRolePermissionsList.Results))
	assert.Equal("create-report", adminRolePermissionsList.Results[0].PermissionId)

	adminUserHasPermission, err = warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		PermissionId: "create-report",
		UserId:       adminUser.UserId,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.True(adminUserHasPermission)

	adminUserRolesList, err = role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(adminUserRolesList.Results))
	assert.Equal("administrator", adminUserRolesList.Results[0].RoleId)

	// Remove create-report permission -> admin role -> admin user
	err = permission.RemovePermissionFromRole(createPermission.PermissionId, adminRole.RoleId)
	if err != nil {
		t.Fatal(err)
	}

	adminUserHasPermission, err = warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		PermissionId: "create-report",
		UserId:       adminUser.UserId,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(adminUserHasPermission)

	adminUserRolesList, err = role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(adminUserRolesList.Results))

	err = role.RemoveRoleFromUser(adminRole.RoleId, adminUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	adminUserRolesList, err = role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(adminUserRolesList.Results))

	// Viewer user tests
	viewerUserHasPermission, err := warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		PermissionId: "view-report",
		UserId:       viewerUser.UserId,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(viewerUserHasPermission)

	viewerUserPermissionsList, err := permission.ListPermissionsForUser(viewerUser.UserId, &warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(viewerUserPermissionsList.Results))

	// Assign view-report permission -> viewer user
	_, err = permission.AssignPermissionToUser(viewPermission.PermissionId, viewerUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	viewerUserHasPermission, err = warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		PermissionId: "view-report",
		UserId:       viewerUser.UserId,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.True(viewerUserHasPermission)

	viewerUserPermissionsList, err = permission.ListPermissionsForUser(viewerUser.UserId, &warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(viewerUserPermissionsList.Results))
	assert.Equal("view-report", viewerUserPermissionsList.Results[0].PermissionId)

	// Remove view-report permission -> viewer user
	err = permission.RemovePermissionFromUser(viewPermission.PermissionId, viewerUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	viewerUserHasPermission, err = warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		PermissionId: "view-report",
		UserId:       viewerUser.UserId,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(viewerUserHasPermission)

	viewerUserPermissionsList, err = permission.ListPermissionsForUser(viewerUser.UserId, &warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(viewerUserPermissionsList.Results))

	// Clean up
	err = user.Delete(adminUser.UserId)
	if err != nil {
		t.Fatal(err)
	}
	err = user.Delete(viewerUser.UserId)
	if err != nil {
		t.Fatal(err)
	}
	err = role.Delete(adminRole.RoleId)
	if err != nil {
		t.Fatal(err)
	}
	err = role.Delete(viewerRole.RoleId)
	if err != nil {
		t.Fatal(err)
	}
	err = permission.Delete(createPermission.PermissionId)
	if err != nil {
		t.Fatal(err)
	}
	err = permission.Delete(viewPermission.PermissionId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPricingTiersAndFeaturesUsers(t *testing.T) {
	setup()
	assert := assert.New(t)

	// Create users
	freeUser, err := user.Create(&warrant.UserParams{})
	if err != nil {
		t.Fatal(err)
	}

	paidUser, err := user.Create(&warrant.UserParams{})
	if err != nil {
		t.Fatal(err)
	}

	// Create pricing tieres
	freeTier, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "free",
	})
	if err != nil {
		t.Fatal(err)
	}

	paidTier, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "paid",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create features
	customFeature, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "custom-feature",
	})
	if err != nil {
		t.Fatal(err)
	}

	feature1, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "feature-1",
	})
	if err != nil {
		t.Fatal(err)
	}

	feature2, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "feature-2",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Paid user tests
	paidUserHasFeature, err := warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   paidUser.UserId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(paidUserHasFeature)

	paidUserFeaturesList, err := feature.ListFeaturesForUser(paidUser.UserId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(paidUserFeaturesList.Results))

	// Assign custom feature -> paid user
	_, err = feature.AssignFeatureToUser(customFeature.FeatureId, paidUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	paidUserHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   paidUser.UserId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.True(paidUserHasFeature)

	paidUserFeaturesList, err = feature.ListFeaturesForUser(paidUser.UserId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(paidUserFeaturesList.Results))
	assert.Equal("custom-feature", paidUserFeaturesList.Results[0].FeatureId)

	err = feature.RemoveFeatureFromUser(customFeature.FeatureId, paidUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	paidUserHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   paidUser.UserId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(paidUserHasFeature)

	paidUserFeaturesList, err = feature.ListFeaturesForUser(paidUser.UserId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(paidUserFeaturesList.Results))

	// Free user tests
	freeUserHasFeature, err := warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   freeUser.UserId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(freeUserHasFeature)

	freeTierFeaturesList, err := feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(freeTierFeaturesList.Results))

	freeUserFeaturesList, err := feature.ListFeaturesForUser(freeUser.UserId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(freeUserFeaturesList.Results))

	// Assign feature-1 -> free tier -> free user
	_, err = feature.AssignFeatureToPricingTier(feature1.FeatureId, freeTier.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}

	_, err = pricingtier.AssignPricingTierToUser(freeTier.PricingTierId, freeUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	freeUserHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   freeUser.UserId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.True(freeUserHasFeature)

	freeTierFeaturesList, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(freeTierFeaturesList.Results))

	freeUserTiersList, err := pricingtier.ListPricingTiersForUser(freeUser.UserId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(freeUserTiersList.Results))
	assert.Equal("free", freeUserTiersList.Results[0].PricingTierId)

	err = feature.RemoveFeatureFromPricingTier(feature1.FeatureId, freeTier.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}

	freeUserHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   freeUser.UserId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(freeUserHasFeature)

	freeTierFeaturesList, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(freeTierFeaturesList.Results))

	freeUserTiersList, err = pricingtier.ListPricingTiersForUser(freeUser.UserId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(freeUserTiersList.Results))

	err = pricingtier.RemovePricingTierFromUser(freeTier.PricingTierId, freeUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	freeUserTiersList, err = pricingtier.ListPricingTiersForUser(freeUser.UserId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(freeUserTiersList.Results))

	// Clean up
	err = user.Delete(freeUser.UserId)
	if err != nil {
		t.Fatal(err)
	}
	err = user.Delete(paidUser.UserId)
	if err != nil {
		t.Fatal(err)
	}
	err = pricingtier.Delete(freeTier.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}
	err = pricingtier.Delete(paidTier.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}
	err = feature.Delete(customFeature.FeatureId)
	if err != nil {
		t.Fatal(err)
	}
	err = feature.Delete(feature1.FeatureId)
	if err != nil {
		t.Fatal(err)
	}
	err = feature.Delete(feature2.FeatureId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestPricingTiersAndFeaturesTenants(t *testing.T) {
	setup()
	assert := assert.New(t)

	// Create tenants
	freeTenant, err := tenant.Create(&warrant.TenantParams{})
	if err != nil {
		t.Fatal(err)
	}

	paidTenant, err := tenant.Create(&warrant.TenantParams{})
	if err != nil {
		t.Fatal(err)
	}

	// Create pricing tieres
	freeTier, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "free",
	})
	if err != nil {
		t.Fatal(err)
	}

	paidTier, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "paid",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create features
	customFeature, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "custom-feature",
	})
	if err != nil {
		t.Fatal(err)
	}

	feature1, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "feature-1",
	})
	if err != nil {
		t.Fatal(err)
	}

	feature2, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "feature-2",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Paid tenant tests
	paidTenantHasFeature, err := warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   paidTenant.TenantId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(paidTenantHasFeature)

	paidTenantFeaturesList, err := feature.ListFeaturesForTenant(paidTenant.TenantId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(paidTenantFeaturesList.Results))

	// Assign custom feature -> paid tenant
	_, err = feature.AssignFeatureToTenant(customFeature.FeatureId, paidTenant.TenantId)
	if err != nil {
		t.Fatal(err)
	}

	paidTenantHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   paidTenant.TenantId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.True(paidTenantHasFeature)

	paidTenantFeaturesList, err = feature.ListFeaturesForTenant(paidTenant.TenantId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(paidTenantFeaturesList.Results))
	assert.Equal("custom-feature", paidTenantFeaturesList.Results[0].FeatureId)

	err = feature.RemoveFeatureFromTenant(customFeature.FeatureId, paidTenant.TenantId)
	if err != nil {
		t.Fatal(err)
	}

	paidTenantHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   paidTenant.TenantId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(paidTenantHasFeature)

	paidTenantFeaturesList, err = feature.ListFeaturesForTenant(paidTenant.TenantId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(paidTenantFeaturesList.Results))

	// Free tenant tests
	freeTenantHasFeature, err := warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   freeTenant.TenantId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(freeTenantHasFeature)

	freeTierFeaturesList, err := feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(freeTierFeaturesList.Results))

	freeTenantFeaturesList, err := feature.ListFeaturesForTenant(freeTenant.TenantId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(freeTenantFeaturesList.Results))

	// Assign feature-1 -> free tier -> free tenant
	_, err = feature.AssignFeatureToPricingTier(feature1.FeatureId, freeTier.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}

	_, err = pricingtier.AssignPricingTierToTenant(freeTier.PricingTierId, freeTenant.TenantId)
	if err != nil {
		t.Fatal(err)
	}

	freeTenantHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   freeTenant.TenantId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.True(freeTenantHasFeature)

	freeTierFeaturesList, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(freeTierFeaturesList.Results))

	freeTenantTiersList, err := pricingtier.ListPricingTiersForTenant(freeTenant.TenantId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(freeTenantTiersList.Results))
	assert.Equal("free", freeTenantTiersList.Results[0].PricingTierId)

	err = feature.RemoveFeatureFromPricingTier(feature1.FeatureId, freeTier.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}

	freeTenantHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   freeTenant.TenantId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(freeTenantHasFeature)

	freeTierFeaturesList, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(freeTierFeaturesList.Results))

	freeTenantTiersList, err = pricingtier.ListPricingTiersForTenant(freeTenant.TenantId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(1, len(freeTenantTiersList.Results))

	err = pricingtier.RemovePricingTierFromTenant(freeTier.PricingTierId, freeTenant.TenantId)
	if err != nil {
		t.Fatal(err)
	}

	freeTenantTiersList, err = pricingtier.ListPricingTiersForTenant(freeTenant.TenantId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 100,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(0, len(freeTenantTiersList.Results))

	// Clean up
	err = tenant.Delete(freeTenant.TenantId)
	if err != nil {
		t.Fatal(err)
	}
	err = tenant.Delete(paidTenant.TenantId)
	if err != nil {
		t.Fatal(err)
	}
	err = pricingtier.Delete(freeTier.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}
	err = pricingtier.Delete(paidTier.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}
	err = feature.Delete(customFeature.FeatureId)
	if err != nil {
		t.Fatal(err)
	}
	err = feature.Delete(feature1.FeatureId)
	if err != nil {
		t.Fatal(err)
	}
	err = feature.Delete(feature2.FeatureId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestSessions(t *testing.T) {
	setup()
	assert := assert.New(t)

	user1, err := user.Create(&warrant.UserParams{})
	if err != nil {
		t.Fatal(err)
	}

	tenant1, err := tenant.Create(&warrant.TenantParams{})
	if err != nil {
		t.Fatal(err)
	}

	_, err = user.AssignUserToTenant(user1.UserId, tenant1.TenantId, "admin")
	if err != nil {
		t.Fatal(err)
	}

	authzSessionToken, err := session.CreateAuthorizationSession(&warrant.AuthorizationSessionParams{
		UserId: user1.UserId,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(authzSessionToken)

	ssDashUrl, err := session.CreateSelfServiceSession((&warrant.SelfServiceSessionParams{
		UserId:              user1.UserId,
		TenantId:            tenant1.TenantId,
		RedirectUrl:         "http://localhost:8080",
		SelfServiceStrategy: warrant.SelfServiceStrategyFGAC,
	}))
	if err != nil {
		t.Fatal(err)
	}
	assert.NotEmpty(ssDashUrl)

	// Clean up
	err = user.Delete(user1.UserId)
	if err != nil {
		t.Fatal(err)
	}

	err = tenant.Delete(tenant1.TenantId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWarrants(t *testing.T) {
	setup()
	assert := assert.New(t)

	newUser, err := user.Create(&warrant.UserParams{})
	if err != nil {
		t.Fatal(err)
	}

	newPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "perm1",
		Meta: map[string]interface{}{
			"name":        "Permission 1",
			"description": "Permission with id 1",
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	checkResult, err := warrant.Check(&warrant.WarrantCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		WarrantCheck: warrant.WarrantCheck{
			Object: warrant.Object{
				ObjectType: warrant.ObjectTypePermission,
				ObjectId:   newPermission.PermissionId,
			},
			Relation: "member",
			Subject: warrant.Subject{
				ObjectType: warrant.ObjectTypeUser,
				ObjectId:   newUser.UserId,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(checkResult)

	_, err = warrant.Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePermission,
		ObjectId:   newPermission.PermissionId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   newUser.UserId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	checkResult, err = warrant.Check(&warrant.WarrantCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		WarrantCheck: warrant.WarrantCheck{
			Object: warrant.Object{
				ObjectType: warrant.ObjectTypePermission,
				ObjectId:   newPermission.PermissionId,
			},
			Relation: "member",
			Subject: warrant.Subject{
				ObjectType: warrant.ObjectTypeUser,
				ObjectId:   newUser.UserId,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.True(checkResult)

	queryResult, err := warrant.Query(fmt.Sprintf("select * where %s:%s is *", "user", newUser.UserId), &warrant.QueryParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	assert.Equal(1, len(queryResult.Results))
	assert.Equal("permission", queryResult.Results[0].ObjectType)
	assert.Equal("perm1", queryResult.Results[0].ObjectId)
	assert.Equal("member", queryResult.Results[0].Warrant.Relation)
	assert.NotNil(queryResult.Results[0].Meta)
	assert.Equal("Permission 1", queryResult.Results[0].Meta["name"])
	assert.Equal("Permission with id 1", queryResult.Results[0].Meta["description"])

	err = warrant.Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePermission,
		ObjectId:   newPermission.PermissionId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   newUser.UserId,
		},
	})
	if err != nil {
		t.Fatal(err)
	}

	checkResult, err = warrant.Check(&warrant.WarrantCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		WarrantCheck: warrant.WarrantCheck{
			Object: warrant.Object{
				ObjectType: warrant.ObjectTypePermission,
				ObjectId:   newPermission.PermissionId,
			},
			Relation: "member",
			Subject: warrant.Subject{
				ObjectType: warrant.ObjectTypeUser,
				ObjectId:   newUser.UserId,
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(checkResult)

	// Clean up
	err = user.Delete(newUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	err = permission.Delete(newPermission.PermissionId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestWarrantPolicies(t *testing.T) {
	setup()
	assert := assert.New(t)

	newUser, err := user.Create(&warrant.UserParams{
		UserId: "user-1",
	})
	if err != nil {
		t.Fatal(err)
	}

	newPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "test-permission",
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = warrant.Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePermission,
		ObjectId:   newPermission.PermissionId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   newUser.UserId,
		},
		Policy: `geo == "us"`,
	})
	if err != nil {
		t.Fatal(err)
	}

	checkResult, err := warrant.Check(&warrant.WarrantCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		WarrantCheck: warrant.WarrantCheck{
			Object: warrant.Object{
				ObjectType: warrant.ObjectTypePermission,
				ObjectId:   newPermission.PermissionId,
			},
			Relation: "member",
			Subject: warrant.Subject{
				ObjectType: warrant.ObjectTypeUser,
				ObjectId:   "user-1",
			},
			Context: warrant.PolicyContext{
				"geo": "us",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.True(checkResult)

	checkResult, err = warrant.Check(&warrant.WarrantCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		WarrantCheck: warrant.WarrantCheck{
			Object: warrant.Object{
				ObjectType: warrant.ObjectTypePermission,
				ObjectId:   newPermission.PermissionId,
			},
			Relation: "member",
			Subject: warrant.Subject{
				ObjectType: warrant.ObjectTypeUser,
				ObjectId:   "user-1",
			},
			Context: warrant.PolicyContext{
				"geo": "gb",
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.False(checkResult)

	// Clean up
	err = warrant.Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePermission,
		ObjectId:   newPermission.PermissionId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   "user-1",
		},
		Policy: `geo == "us"`,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = user.Delete(newUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	err = permission.Delete(newPermission.PermissionId)
	if err != nil {
		t.Fatal(err)
	}
}

func TestObjectTypes(t *testing.T) {
	setup()
	assert := assert.New(t)

	relations := make(map[string]interface{})
	relations["relation-1"] = struct{}{}

	newType, err := objecttype.Create(&warrant.ObjectTypeParams{
		Type:      "new-type",
		Relations: relations,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("new-type", newType.Type)
	assert.NotNil(newType.Relations["relation-1"])
	assert.Nil(newType.Relations["relation-2"])

	objType, err := objecttype.Get("new-type", &warrant.ObjectTypeParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("new-type", objType.Type)
	assert.Len(objType.Relations, 1)
	assert.NotNil(objType.Relations["relation-1"])

	types, err := objecttype.ListObjectTypes(&warrant.ListObjectTypeParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(types, 7)

	newRelations := make(map[string]interface{})
	newRelations["relation-1"] = struct{}{}
	newRelations["relation-2"] = struct{}{}
	objType, err = objecttype.Update("new-type", &warrant.ObjectTypeParams{
		Type:      "new-type",
		Relations: newRelations,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("new-type", objType.Type)
	assert.Len(objType.Relations, 2)
	assert.NotNil(objType.Relations["relation-1"])
	assert.NotNil(objType.Relations["relation-2"])

	err = objecttype.Delete("new-type")
	if err != nil {
		t.Fatal(err)
	}

	types, err = objecttype.ListObjectTypes(&warrant.ListObjectTypeParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(types, 6)
}

func TestObjects(t *testing.T) {
	setup()
	assert := assert.New(t)

	newObj, err := object.Create(&warrant.ObjectParams{
		ObjectType: "role",
		ObjectId:   "admin2",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("role", newObj.ObjectType)
	assert.Equal("admin2", newObj.ObjectId)
	assert.Len(newObj.Meta, 0)

	obj, err := object.Get("role", "admin2", &warrant.ObjectParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("role", obj.ObjectType)
	assert.Equal("admin2", obj.ObjectId)
	assert.Len(obj.Meta, 0)

	objectsList, err := object.ListObjects(&warrant.ListObjectParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(objectsList.Results, 1)
	assert.Equal("role", objectsList.Results[0].ObjectType)
	assert.Equal("admin2", objectsList.Results[0].ObjectId)

	meta := make(map[string]interface{})
	meta["name"] = "new name"
	obj, err = object.Update("role", "admin2", &warrant.ObjectParams{
		ObjectType: "role",
		ObjectId:   "admin2",
		Meta:       meta,
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("role", obj.ObjectType)
	assert.Equal("admin2", obj.ObjectId)
	assert.Len(obj.Meta, 1)
	assert.Equal("new name", obj.Meta["name"])

	err = object.Delete("role", "admin2")
	if err != nil {
		t.Fatal(err)
	}

	objectsList, err = object.ListObjects(&warrant.ListObjectParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(objectsList.Results, 0)
}

func TestBatchObjects(t *testing.T) {
	setup()
	assert := assert.New(t)

	createdObjects, err := object.BatchCreate([]warrant.ObjectParams{
		{ObjectType: "document", ObjectId: "document-a"},
		{ObjectType: "document", ObjectId: "document-b"},
		{ObjectType: "folder", ObjectId: "resources", Meta: map[string]interface{}{"description": "Helpful documents"}},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(createdObjects, 3)

	fetchedObjects, err := object.ListObjects(&warrant.ListObjectParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(fetchedObjects.Results, 3)
	assert.Equal("document", fetchedObjects.Results[0].ObjectType)
	assert.Equal("document-a", fetchedObjects.Results[0].ObjectId)
	assert.Equal("document", fetchedObjects.Results[1].ObjectType)
	assert.Equal("document-b", fetchedObjects.Results[1].ObjectId)
	assert.Equal("folder", fetchedObjects.Results[2].ObjectType)
	assert.Equal("resources", fetchedObjects.Results[2].ObjectId)
	assert.Equal(map[string]interface{}{"description": "Helpful documents"}, fetchedObjects.Results[2].Meta)

	fetchedObjects, err = object.ListObjects(&warrant.ListObjectParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
		Query: "resource",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(fetchedObjects.Results, 1)
	assert.Equal("folder", fetchedObjects.Results[0].ObjectType)
	assert.Equal("resources", fetchedObjects.Results[0].ObjectId)
	assert.Equal(map[string]interface{}{"description": "Helpful documents"}, fetchedObjects.Results[0].Meta)

	err = object.BatchDelete([]warrant.ObjectParams{
		{ObjectType: "document", ObjectId: "document-a"},
		{ObjectType: "document", ObjectId: "document-b"},
		{ObjectType: "folder", ObjectId: "resources"},
	})
	if err != nil {
		t.Fatal(err)
	}

	fetchedObjects, err = object.ListObjects(&warrant.ListObjectParams{
		ListParams: warrant.ListParams{
			RequestOptions: warrant.RequestOptions{
				WarrantToken: "latest",
			},
			Limit: 10,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(fetchedObjects.Results, 0)
}
