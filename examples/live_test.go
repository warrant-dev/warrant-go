package main

import (
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
	assert.Empty(user1.Email)

	user2, err := user.Create(&warrant.UserParams{
		UserId: "some_id",
		Email:  "test@email.com",
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
	assert.Equal(user2.Email, refetchedUser.Email)

	user2, err = user.Update("some_id", &warrant.UserParams{
		Email: "updated@email.com",
	})
	if err != nil {
		t.Fatal(err)
	}
	refetchedUser, err = user.Get("some_id", fetchUserParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("some_id", refetchedUser.UserId)
	assert.Equal("updated@email.com", refetchedUser.Email)

	users, err := user.ListUsers(&warrant.ListUserParams{
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
	assert.Equal(2, len(users))

	err = user.Delete(user1.UserId)
	if err != nil {
		t.Fatal(err)
	}
	err = user.Delete(user2.UserId)
	if err != nil {
		t.Fatal(err)
	}
	users, err = user.ListUsers(&warrant.ListUserParams{
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
	assert.Equal(0, len(users))
}

func TestCrudTenants(t *testing.T) {
	setup()
	assert := assert.New(t)

	tenant1, err := tenant.Create(&warrant.TenantParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.NotNil(tenant1.TenantId)
	assert.Empty(tenant1.Name)

	tenant2, err := tenant.Create(&warrant.TenantParams{
		TenantId: "some_tenant_id",
		Name:     "new_name",
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
	assert.Equal(tenant2.Name, refetchedTenant.Name)

	tenant2, err = tenant.Update("some_tenant_id", &warrant.TenantParams{
		Name: "updated_name",
	})
	if err != nil {
		t.Fatal(err)
	}
	refetchedTenant, err = tenant.Get("some_tenant_id", fetchTenantParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("some_tenant_id", refetchedTenant.TenantId)
	assert.Equal("updated_name", refetchedTenant.Name)

	tenants, err := tenant.ListTenants(&warrant.ListTenantParams{
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
	assert.Equal(2, len(tenants))

	err = tenant.Delete(tenant1.TenantId)
	if err != nil {
		t.Fatal(err)
	}
	err = tenant.Delete(tenant2.TenantId)
	if err != nil {
		t.Fatal(err)
	}
	tenants, err = tenant.ListTenants(&warrant.ListTenantParams{
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
	assert.Equal(0, len(tenants))
}

func TestCrudRoles(t *testing.T) {
	setup()
	assert := assert.New(t)

	adminRole, err := role.Create(&warrant.RoleParams{
		RoleId:      "admin",
		Name:        "Admin",
		Description: "The admin role",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("admin", adminRole.RoleId)
	assert.Equal("Admin", adminRole.Name)
	assert.Equal("The admin role", adminRole.Description)

	viewerRole, err := role.Create(&warrant.RoleParams{
		RoleId:      "viewer",
		Name:        "Viewer",
		Description: "The viewer role",
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
	assert.Equal(viewerRole.Name, refetchedRole.Name)
	assert.Equal(viewerRole.Description, refetchedRole.Description)

	viewerRole, err = role.Update("viewer", &warrant.RoleParams{
		Name:        "Viewer Updated",
		Description: "Updated desc",
	})
	if err != nil {
		t.Fatal(err)
	}
	refetchedRole, err = role.Get(viewerRole.RoleId, fetchRoleParams)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("viewer", refetchedRole.RoleId)
	assert.Equal("Viewer Updated", refetchedRole.Name)
	assert.Equal("Updated desc", refetchedRole.Description)

	roles, err := role.ListRoles(&warrant.ListRoleParams{
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
	assert.Equal(2, len(roles))

	err = role.Delete(adminRole.RoleId)
	if err != nil {
		t.Fatal(err)
	}
	err = role.Delete(viewerRole.RoleId)
	if err != nil {
		t.Fatal(err)
	}
	roles, err = role.ListRoles(&warrant.ListRoleParams{
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
	assert.Equal(0, len(roles))
}

func TestCrudPermissions(t *testing.T) {
	setup()
	assert := assert.New(t)

	permission1, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "perm1",
		Name:         "Permission 1",
		Description:  "Permission with id 1",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("perm1", permission1.PermissionId)
	assert.Equal("Permission 1", permission1.Name)
	assert.Equal("Permission with id 1", permission1.Description)

	permission2, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "perm2",
		Name:         "Permission 2",
		Description:  "Permission with id 2",
	})
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal("perm2", permission2.PermissionId)
	assert.Equal("Permission 2", permission2.Name)
	assert.Equal("Permission with id 2", permission2.Description)

	permission2, err = permission.Update("perm2", &warrant.PermissionParams{
		Name:        "Permission 2 Updated",
		Description: "Updated desc",
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
	assert.Equal("Permission 2 Updated", refetchedPermission.Name)
	assert.Equal("Updated desc", refetchedPermission.Description)

	permissions, err := permission.ListPermissions(&warrant.ListPermissionParams{
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
	assert.Equal(2, len(permissions))

	err = permission.Delete(permission1.PermissionId)
	if err != nil {
		t.Fatal(err)
	}
	err = permission.Delete(permission2.PermissionId)
	if err != nil {
		t.Fatal(err)
	}
	permissions, err = permission.ListPermissions(&warrant.ListPermissionParams{
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
	assert.Equal(0, len(permissions))
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
	assert.Equal("feature-2", refetchedFeature.FeatureId)

	features, err := feature.ListFeatures(&warrant.ListFeatureParams{
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
	assert.Equal(2, len(features))

	err = feature.Delete(feature1.FeatureId)
	if err != nil {
		t.Fatal(err)
	}
	err = feature.Delete(feature2.FeatureId)
	if err != nil {
		t.Fatal(err)
	}
	features, err = feature.ListFeatures(&warrant.ListFeatureParams{
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
	assert.Equal(0, len(features))
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

	tiers, err := pricingtier.ListPricingTiers(&warrant.ListPricingTierParams{
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
	assert.Equal(2, len(tiers))

	err = pricingtier.Delete(tier1.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}
	err = pricingtier.Delete(tier2.PricingTierId)
	if err != nil {
		t.Fatal(err)
	}
	tiers, err = pricingtier.ListPricingTiers(&warrant.ListPricingTierParams{
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
	assert.Equal(0, len(tiers))
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
		Name:     "Tenant 1",
	})
	if err != nil {
		t.Fatal(err)
	}
	tenant2, err := tenant.Create(&warrant.TenantParams{
		TenantId: "tenant-2",
		Name:     "Tenant 2",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Assign user1 -> tenant1
	user1Tenants, err := tenant.ListTenantsForUser(user1.UserId, &warrant.ListTenantParams{
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
	assert.Equal(0, len(user1Tenants))
	tenant1Users, err := user.ListUsersForTenant(tenant1.TenantId, &warrant.ListUserParams{
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
	assert.Equal(0, len(tenant1Users))

	_, err = user.AssignUserToTenant(user1.UserId, tenant1.TenantId, "member")
	if err != nil {
		t.Fatal(err)
	}

	user1Tenants, err = tenant.ListTenantsForUser(user1.UserId, &warrant.ListTenantParams{
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
	assert.Equal(1, len(user1Tenants))
	assert.Equal("tenant-1", user1Tenants[0].TenantId)
	tenant1Users, err = user.ListUsersForTenant(tenant1.TenantId, &warrant.ListUserParams{
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
	assert.Equal(1, len(tenant1Users))
	assert.Equal(user1.UserId, tenant1Users[0].UserId)

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
		RoleId:      "administrator",
		Name:        "Administrator",
		Description: "The admin role",
	})
	if err != nil {
		t.Fatal(err)
	}
	viewerRole, err := role.Create(&warrant.RoleParams{
		RoleId:      "viewer",
		Name:        "Viewer",
		Description: "The viewer role",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Create permissions
	createPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "create-report",
		Name:         "Create Report",
		Description:  "Permission to create reports",
	})
	if err != nil {
		t.Fatal(err)
	}
	viewPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "view-report",
		Name:         "View Report",
		Description:  "Permission to view reports",
	})
	if err != nil {
		t.Fatal(err)
	}

	// Admin user tests
	adminUserRoles, err := role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
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
	assert.Equal(0, len(adminUserRoles))

	adminRolePermissions, err := permission.ListPermissionsForRole(adminRole.RoleId, &warrant.ListPermissionParams{
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
	assert.Equal(0, len(adminRolePermissions))

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

	adminRolePermissions, err = permission.ListPermissionsForRole(adminRole.RoleId, &warrant.ListPermissionParams{
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
	assert.Equal(1, len(adminRolePermissions))
	assert.Equal("create-report", adminRolePermissions[0].PermissionId)

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

	adminUserRoles, err = role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
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
	assert.Equal(1, len(adminUserRoles))
	assert.Equal("administrator", adminUserRoles[0].RoleId)

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

	adminUserRoles, err = role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
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
	assert.Equal(1, len(adminUserRoles))

	err = role.RemoveRoleFromUser(adminRole.RoleId, adminUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	adminUserRoles, err = role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
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
	assert.Equal(0, len(adminUserRoles))

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

	viewerUserPermissions, err := permission.ListPermissionsForUser(viewerUser.UserId, &warrant.ListPermissionParams{
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
	assert.Equal(0, len(viewerUserPermissions))

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

	viewerUserPermissions, err = permission.ListPermissionsForUser(viewerUser.UserId, &warrant.ListPermissionParams{
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
	assert.Equal(1, len(viewerUserPermissions))
	assert.Equal("view-report", viewerUserPermissions[0].PermissionId)

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

	viewerUserPermissions, err = permission.ListPermissionsForUser(viewerUser.UserId, &warrant.ListPermissionParams{
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
	assert.Equal(0, len(viewerUserPermissions))

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

	paidUserFeatures, err := feature.ListFeaturesForUser(paidUser.UserId, &warrant.ListFeatureParams{
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
	assert.Equal(0, len(paidUserFeatures))

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

	paidUserFeatures, err = feature.ListFeaturesForUser(paidUser.UserId, &warrant.ListFeatureParams{
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
	assert.Equal(1, len(paidUserFeatures))
	assert.Equal("custom-feature", paidUserFeatures[0].FeatureId)

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

	paidUserFeatures, err = feature.ListFeaturesForUser(paidUser.UserId, &warrant.ListFeatureParams{
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
	assert.Equal(0, len(paidUserFeatures))

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

	freeTierFeatures, err := feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
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
	assert.Equal(0, len(freeTierFeatures))

	freeUserFeatures, err := feature.ListFeaturesForUser(freeUser.UserId, &warrant.ListFeatureParams{
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
	assert.Equal(0, len(freeUserFeatures))

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

	freeTierFeatures, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
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
	assert.Equal(1, len(freeTierFeatures))

	freeUserTiers, err := pricingtier.ListPricingTiersForUser(freeUser.UserId, &warrant.ListPricingTierParams{
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
	assert.Equal(1, len(freeUserTiers))
	assert.Equal("free", freeUserTiers[0].PricingTierId)

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

	freeTierFeatures, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
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
	assert.Equal(0, len(freeTierFeatures))

	freeUserTiers, err = pricingtier.ListPricingTiersForUser(freeUser.UserId, &warrant.ListPricingTierParams{
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
	assert.Equal(1, len(freeUserTiers))

	err = pricingtier.RemovePricingTierFromUser(freeTier.PricingTierId, freeUser.UserId)
	if err != nil {
		t.Fatal(err)
	}

	freeUserTiers, err = pricingtier.ListPricingTiersForUser(freeUser.UserId, &warrant.ListPricingTierParams{
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
	assert.Equal(0, len(freeUserTiers))

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

	paidTenantFeatures, err := feature.ListFeaturesForTenant(paidTenant.TenantId, &warrant.ListFeatureParams{
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
	assert.Equal(0, len(paidTenantFeatures))

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

	paidTenantFeatures, err = feature.ListFeaturesForTenant(paidTenant.TenantId, &warrant.ListFeatureParams{
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
	assert.Equal(1, len(paidTenantFeatures))
	assert.Equal("custom-feature", paidTenantFeatures[0].FeatureId)

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

	paidTenantFeatures, err = feature.ListFeaturesForTenant(paidTenant.TenantId, &warrant.ListFeatureParams{
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
	assert.Equal(0, len(paidTenantFeatures))

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

	freeTierFeatures, err := feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
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
	assert.Equal(0, len(freeTierFeatures))

	freeTenantFeatures, err := feature.ListFeaturesForTenant(freeTenant.TenantId, &warrant.ListFeatureParams{
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
	assert.Equal(0, len(freeTenantFeatures))

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

	freeTierFeatures, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
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
	assert.Equal(1, len(freeTierFeatures))

	freeTenantTiers, err := pricingtier.ListPricingTiersForTenant(freeTenant.TenantId, &warrant.ListPricingTierParams{
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
	assert.Equal(1, len(freeTenantTiers))
	assert.Equal("free", freeTenantTiers[0].PricingTierId)

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

	freeTierFeatures, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
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
	assert.Equal(0, len(freeTierFeatures))

	freeTenantTiers, err = pricingtier.ListPricingTiersForTenant(freeTenant.TenantId, &warrant.ListPricingTierParams{
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
	assert.Equal(1, len(freeTenantTiers))

	err = pricingtier.RemovePricingTierFromTenant(freeTier.PricingTierId, freeTenant.TenantId)
	if err != nil {
		t.Fatal(err)
	}

	freeTenantTiers, err = pricingtier.ListPricingTiersForTenant(freeTenant.TenantId, &warrant.ListPricingTierParams{
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
	assert.Equal(0, len(freeTenantTiers))

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
		Name:         "Permission 1",
		Description:  "Permission with id 1",
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

	// queryResult, err := warrant.Query(fmt.Sprintf("SELECT warrant FOR subject=%s:%s WHERE subject=%s:%s", "user", newUser.UserId, "user", newUser.UserId), &warrant.ListWarrantParams{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// queryBytes, _ := json.Marshal(queryResult.Result)
	// var result []warrant.Warrant
	// json.Unmarshal(queryBytes, &result)

	// assert.Equal(1, len(result))
	// assert.Equal("permission", result[0].ObjectType)
	// assert.Equal("perm1", result[0].ObjectId)
	// assert.Equal("member", result[0].Relation)

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

	_, err := warrant.Create(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePermission,
		ObjectId:   "test-permission",
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

	checkResult, err := warrant.Check(&warrant.WarrantCheckParams{
		RequestOptions: warrant.RequestOptions{
			WarrantToken: "latest",
		},
		WarrantCheck: warrant.WarrantCheck{
			Object: warrant.Object{
				ObjectType: warrant.ObjectTypePermission,
				ObjectId:   "test-permission",
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
				ObjectId:   "test-permission",
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

	err = warrant.Delete(&warrant.WarrantParams{
		ObjectType: warrant.ObjectTypePermission,
		ObjectId:   "test-permission",
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

	objects, err := object.ListObjects(&warrant.ListObjectParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(objects, 1)

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

	objects, err = object.ListObjects(&warrant.ListObjectParams{})
	if err != nil {
		t.Fatal(err)
	}
	assert.Len(objects, 0)
}
