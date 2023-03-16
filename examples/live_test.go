package main

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/warrant-dev/warrant-go/v3"
	"github.com/warrant-dev/warrant-go/v3/feature"
	"github.com/warrant-dev/warrant-go/v3/permission"
	"github.com/warrant-dev/warrant-go/v3/pricingtier"
	"github.com/warrant-dev/warrant-go/v3/role"
	"github.com/warrant-dev/warrant-go/v3/session"
	"github.com/warrant-dev/warrant-go/v3/tenant"
	"github.com/warrant-dev/warrant-go/v3/user"
)

func setup() {
	warrant.ApiKey = "YOUR_API_KEY"
}

func TestCrudUsers(t *testing.T) {
	setup()
	assert := assert.New(t)

	user1, err := user.Create(&warrant.UserParams{})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.NotNil(user1.UserId)
	assert.Empty(user1.Email)

	user2, err := user.Create(&warrant.UserParams{
		UserId: "some_id",
		Email:  "test@email.com",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	refetchedUser, err := user.Get(user2.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(user2.UserId, refetchedUser.UserId)
	assert.Equal(user2.Email, refetchedUser.Email)

	user2, err = user.Update("some_id", &warrant.UserParams{
		Email: "updated@email.com",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	refetchedUser, err = user.Get("some_id")
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal("some_id", refetchedUser.UserId)
	assert.Equal("updated@email.com", refetchedUser.Email)

	users, err := user.ListUsers(&warrant.ListUserParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(2, len(users))

	err = user.Delete(user1.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = user.Delete(user2.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	users, err = user.ListUsers(&warrant.ListUserParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(users))
}

func TestCrudTenants(t *testing.T) {
	setup()
	assert := assert.New(t)

	tenant1, err := tenant.Create(&warrant.TenantParams{})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.NotNil(tenant1.TenantId)
	assert.Empty(tenant1.Name)

	tenant2, err := tenant.Create(&warrant.TenantParams{
		TenantId: "some_tenant_id",
		Name:     "new_name",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	refetchedTenant, err := tenant.Get(tenant2.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(tenant2.TenantId, refetchedTenant.TenantId)
	assert.Equal(tenant2.Name, refetchedTenant.Name)

	tenant2, err = tenant.Update("some_tenant_id", &warrant.TenantParams{
		Name: "updated_name",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	refetchedTenant, err = tenant.Get("some_tenant_id")
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal("some_tenant_id", refetchedTenant.TenantId)
	assert.Equal("updated_name", refetchedTenant.Name)

	tenants, err := tenant.ListTenants(&warrant.ListTenantParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(2, len(tenants))

	err = tenant.Delete(tenant1.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tenant.Delete(tenant2.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}
	tenants, err = tenant.ListTenants(&warrant.ListTenantParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
	}
	refetchedRole, err := role.Get(viewerRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(viewerRole.RoleId, refetchedRole.RoleId)
	assert.Equal(viewerRole.Name, refetchedRole.Name)
	assert.Equal(viewerRole.Description, refetchedRole.Description)

	viewerRole, err = role.Update("viewer", &warrant.RoleParams{
		Name:        "Viewer Updated",
		Description: "Updated desc",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	refetchedRole, err = role.Get(viewerRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal("viewer", refetchedRole.RoleId)
	assert.Equal("Viewer Updated", refetchedRole.Name)
	assert.Equal("Updated desc", refetchedRole.Description)

	roles, err := role.ListRoles(&warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(2, len(roles))

	err = role.Delete(adminRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = role.Delete(viewerRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	}
	roles, err = role.ListRoles(&warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
	}
	assert.Equal("perm2", permission2.PermissionId)
	assert.Equal("Permission 2", permission2.Name)
	assert.Equal("Permission with id 2", permission2.Description)

	permission2, err = permission.Update("perm2", &warrant.PermissionParams{
		Name:        "Permission 2 Updated",
		Description: "Updated desc",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	refetchedPermission, err := permission.Get("perm2")
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal("perm2", refetchedPermission.PermissionId)
	assert.Equal("Permission 2 Updated", refetchedPermission.Name)
	assert.Equal("Updated desc", refetchedPermission.Description)

	permissions, err := permission.ListPermissions(&warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(2, len(permissions))

	err = permission.Delete(permission1.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = permission.Delete(permission2.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	}
	permissions, err = permission.ListPermissions(&warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
	}
	assert.Equal("new-feature", feature1.FeatureId)

	feature2, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "feature-2",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	refetchedFeature, err := feature.Get(feature2.FeatureId)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal("feature-2", refetchedFeature.FeatureId)

	features, err := feature.ListFeatures(&warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(2, len(features))

	err = feature.Delete(feature1.FeatureId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = feature.Delete(feature2.FeatureId)
	if err != nil {
		fmt.Println(err)
		return
	}
	features, err = feature.ListFeatures(&warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
	}
	assert.Equal("new-tier1", tier1.PricingTierId)

	tier2, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "tier-2",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	refetchedTier, err := pricingtier.Get(tier2.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(tier2.PricingTierId, refetchedTier.PricingTierId)

	tiers, err := pricingtier.ListPricingTiers(&warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(2, len(tiers))

	err = pricingtier.Delete(tier1.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pricingtier.Delete(tier2.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}
	tiers, err = pricingtier.ListPricingTiers(&warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
	}
	assert.Equal(2, len(createdUsers))
	assert.Equal("user-1", createdUsers[0].UserId)
	assert.Equal("user-2", createdUsers[1].UserId)

	createdTenants, err := tenant.BatchCreate([]warrant.TenantParams{
		{TenantId: "tenant-1"},
		{TenantId: "tenant-2"},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(2, len(createdTenants))
	assert.Equal("tenant-1", createdTenants[0].TenantId)
	assert.Equal("tenant-2", createdTenants[1].TenantId)

	err = user.Delete("user-1")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = user.Delete("user-2")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tenant.Delete("tenant-1")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tenant.Delete("tenant-2")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestMultiTenancy(t *testing.T) {
	setup()
	assert := assert.New(t)

	// Create users
	user1, err := user.Create(&warrant.UserParams{})
	if err != nil {
		fmt.Println(err)
		return
	}
	user2, err := user.Create(&warrant.UserParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create tenants
	tenant1, err := tenant.Create(&warrant.TenantParams{
		TenantId: "tenant-1",
		Name:     "Tenant 1",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	tenant2, err := tenant.Create(&warrant.TenantParams{
		TenantId: "tenant-2",
		Name:     "Tenant 2",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Assign user1 -> tenant1
	user1Tenants, err := tenant.ListTenantsForUser(user1.UserId, &warrant.ListTenantParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(user1Tenants))
	tenant1Users, err := user.ListUsersForTenant(tenant1.TenantId, &warrant.ListUserParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(tenant1Users))

	_, err = user.AssignUserToTenant(user1.UserId, tenant1.TenantId, "member")
	if err != nil {
		fmt.Println(err)
		return
	}

	user1Tenants, err = tenant.ListTenantsForUser(user1.UserId, &warrant.ListTenantParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(user1Tenants))
	assert.Equal("tenant-1", user1Tenants[0].TenantId)
	tenant1Users, err = user.ListUsersForTenant(tenant1.TenantId, &warrant.ListUserParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(tenant1Users))
	assert.Equal(user1.UserId, tenant1Users[0].UserId)

	err = user.RemoveUserFromTenant(user1.UserId, tenant1.TenantId, "member")
	if err != nil {
		fmt.Println(err)
		return
	}

	// Clean up
	err = user.Delete(user1.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = user.Delete(user2.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tenant.Delete(tenant1.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tenant.Delete(tenant2.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestRBAC(t *testing.T) {
	setup()
	assert := assert.New(t)

	// Create users
	adminUser, err := user.Create(&warrant.UserParams{})
	if err != nil {
		fmt.Println(err)
		return
	}
	viewerUser, err := user.Create(&warrant.UserParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create roles
	adminRole, err := role.Create(&warrant.RoleParams{
		RoleId:      "administrator",
		Name:        "Administrator",
		Description: "The admin role",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	viewerRole, err := role.Create(&warrant.RoleParams{
		RoleId:      "viewer",
		Name:        "Viewer",
		Description: "The viewer role",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create permissions
	createPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "create-report",
		Name:         "Create Report",
		Description:  "Permission to create reports",
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	viewPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "view-report",
		Name:         "View Report",
		Description:  "Permission to view reports",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Admin user tests
	adminUserRoles, err := role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(adminUserRoles))

	adminRolePermissions, err := permission.ListPermissionsForRole(adminRole.RoleId, &warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(adminRolePermissions))

	adminUserHasPermission, err := warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		PermissionId: "create-report",
		UserId:       adminUser.UserId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(adminUserHasPermission)

	// Assign create-report permission -> admin role -> admin user
	_, err = permission.AssignPermissionToRole(createPermission.PermissionId, adminRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = role.AssignRoleToUser(adminRole.RoleId, adminUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	adminRolePermissions, err = permission.ListPermissionsForRole(adminRole.RoleId, &warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(adminRolePermissions))
	assert.Equal("create-report", adminRolePermissions[0].PermissionId)

	adminUserHasPermission, err = warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		PermissionId: "create-report",
		UserId:       adminUser.UserId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.True(adminUserHasPermission)

	adminUserRoles, err = role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(adminUserRoles))
	assert.Equal("administrator", adminUserRoles[0].RoleId)

	// Remove create-report permission -> admin role -> admin user
	err = permission.RemovePermissionFromRole(createPermission.PermissionId, adminRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	}

	adminUserHasPermission, err = warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		PermissionId: "create-report",
		UserId:       adminUser.UserId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(adminUserHasPermission)

	adminUserRoles, err = role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(adminUserRoles))

	err = role.RemoveRoleFromUser(adminRole.RoleId, adminUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	adminUserRoles, err = role.ListRolesForUser(adminUser.UserId, &warrant.ListRoleParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(adminUserRoles))

	// Viewer user tests
	viewerUserHasPermission, err := warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		PermissionId: "view-report",
		UserId:       viewerUser.UserId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(viewerUserHasPermission)

	viewerUserPermissions, err := permission.ListPermissionsForUser(viewerUser.UserId, &warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(viewerUserPermissions))

	// Assign view-report permission -> viewer user
	_, err = permission.AssignPermissionToUser(viewPermission.PermissionId, viewerUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	viewerUserHasPermission, err = warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		PermissionId: "view-report",
		UserId:       viewerUser.UserId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.True(viewerUserHasPermission)

	viewerUserPermissions, err = permission.ListPermissionsForUser(viewerUser.UserId, &warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(viewerUserPermissions))
	assert.Equal("view-report", viewerUserPermissions[0].PermissionId)

	// Remove view-report permission -> viewer user
	err = permission.RemovePermissionFromUser(viewPermission.PermissionId, viewerUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	viewerUserHasPermission, err = warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		PermissionId: "view-report",
		UserId:       viewerUser.UserId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(viewerUserHasPermission)

	viewerUserPermissions, err = permission.ListPermissionsForUser(viewerUser.UserId, &warrant.ListPermissionParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(viewerUserPermissions))

	// Clean up
	err = user.Delete(adminUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = user.Delete(viewerUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = role.Delete(adminRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = role.Delete(viewerRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = permission.Delete(createPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = permission.Delete(viewPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestPricingTiersAndFeaturesUsers(t *testing.T) {
	setup()
	assert := assert.New(t)

	// Create users
	freeUser, err := user.Create(&warrant.UserParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	paidUser, err := user.Create(&warrant.UserParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create pricing tieres
	freeTier, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "free",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	paidTier, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "paid",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create features
	customFeature, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "custom-feature",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	feature1, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "feature-1",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	feature2, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "feature-2",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Paid user tests
	paidUserHasFeature, err := warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   paidUser.UserId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(paidUserHasFeature)

	paidUserFeatures, err := feature.ListFeaturesForUser(paidUser.UserId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(paidUserFeatures))

	// Assign custom feature -> paid user
	_, err = feature.AssignFeatureToUser(customFeature.FeatureId, paidUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	paidUserHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   paidUser.UserId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.True(paidUserHasFeature)

	paidUserFeatures, err = feature.ListFeaturesForUser(paidUser.UserId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(paidUserFeatures))
	assert.Equal("custom-feature", paidUserFeatures[0].FeatureId)

	err = feature.RemoveFeatureFromUser(customFeature.FeatureId, paidUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	paidUserHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   paidUser.UserId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(paidUserHasFeature)

	paidUserFeatures, err = feature.ListFeaturesForUser(paidUser.UserId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(paidUserFeatures))

	// Free user tests
	freeUserHasFeature, err := warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   freeUser.UserId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(freeUserHasFeature)

	freeTierFeatures, err := feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(freeTierFeatures))

	freeUserFeatures, err := feature.ListFeaturesForUser(freeUser.UserId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(freeUserFeatures))

	// Assign feature-1 -> free tier -> free user
	_, err = feature.AssignFeatureToPricingTier(feature1.FeatureId, freeTier.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = pricingtier.AssignPricingTierToUser(freeTier.PricingTierId, freeUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	freeUserHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   freeUser.UserId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.True(freeUserHasFeature)

	freeTierFeatures, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(freeTierFeatures))

	freeUserTiers, err := pricingtier.ListPricingTiersForUser(freeUser.UserId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(freeUserTiers))
	assert.Equal("free", freeUserTiers[0].PricingTierId)

	err = feature.RemoveFeatureFromPricingTier(feature1.FeatureId, freeTier.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}

	freeUserHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeUser,
			ObjectId:   freeUser.UserId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(freeUserHasFeature)

	freeTierFeatures, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(freeTierFeatures))

	freeUserTiers, err = pricingtier.ListPricingTiersForUser(freeUser.UserId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(freeUserTiers))

	err = pricingtier.RemovePricingTierFromUser(freeTier.PricingTierId, freeUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	freeUserTiers, err = pricingtier.ListPricingTiersForUser(freeUser.UserId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(freeUserTiers))

	// Clean up
	err = user.Delete(freeUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = user.Delete(paidUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pricingtier.Delete(freeTier.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pricingtier.Delete(paidTier.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = feature.Delete(customFeature.FeatureId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = feature.Delete(feature1.FeatureId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = feature.Delete(feature2.FeatureId)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestPricingTiersAndFeaturesTenants(t *testing.T) {
	setup()
	assert := assert.New(t)

	// Create tenants
	freeTenant, err := tenant.Create(&warrant.TenantParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	paidTenant, err := tenant.Create(&warrant.TenantParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create pricing tieres
	freeTier, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "free",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	paidTier, err := pricingtier.Create(&warrant.PricingTierParams{
		PricingTierId: "paid",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create features
	customFeature, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "custom-feature",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	feature1, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "feature-1",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	feature2, err := feature.Create(&warrant.FeatureParams{
		FeatureId: "feature-2",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// Paid tenant tests
	paidTenantHasFeature, err := warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   paidTenant.TenantId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(paidTenantHasFeature)

	paidTenantFeatures, err := feature.ListFeaturesForTenant(paidTenant.TenantId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(paidTenantFeatures))

	// Assign custom feature -> paid tenant
	_, err = feature.AssignFeatureToTenant(customFeature.FeatureId, paidTenant.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}

	paidTenantHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   paidTenant.TenantId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.True(paidTenantHasFeature)

	paidTenantFeatures, err = feature.ListFeaturesForTenant(paidTenant.TenantId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(paidTenantFeatures))
	assert.Equal("custom-feature", paidTenantFeatures[0].FeatureId)

	err = feature.RemoveFeatureFromTenant(customFeature.FeatureId, paidTenant.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}

	paidTenantHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "custom-feature",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   paidTenant.TenantId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(paidTenantHasFeature)

	paidTenantFeatures, err = feature.ListFeaturesForTenant(paidTenant.TenantId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(paidTenantFeatures))

	// Free tenant tests
	freeTenantHasFeature, err := warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   freeTenant.TenantId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(freeTenantHasFeature)

	freeTierFeatures, err := feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(freeTierFeatures))

	freeTenantFeatures, err := feature.ListFeaturesForTenant(freeTenant.TenantId, &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(freeTenantFeatures))

	// Assign feature-1 -> free tier -> free tenant
	_, err = feature.AssignFeatureToPricingTier(feature1.FeatureId, freeTier.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = pricingtier.AssignPricingTierToTenant(freeTier.PricingTierId, freeTenant.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}

	freeTenantHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   freeTenant.TenantId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.True(freeTenantHasFeature)

	freeTierFeatures, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(freeTierFeatures))

	freeTenantTiers, err := pricingtier.ListPricingTiersForTenant(freeTenant.TenantId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(freeTenantTiers))
	assert.Equal("free", freeTenantTiers[0].PricingTierId)

	err = feature.RemoveFeatureFromPricingTier(feature1.FeatureId, freeTier.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}

	freeTenantHasFeature, err = warrant.CheckHasFeature(&warrant.FeatureCheckParams{
		FeatureId: "feature-1",
		Subject: warrant.Subject{
			ObjectType: warrant.ObjectTypeTenant,
			ObjectId:   freeTenant.TenantId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.False(freeTenantHasFeature)

	freeTierFeatures, err = feature.ListFeaturesForPricingTier("free", &warrant.ListFeatureParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(freeTierFeatures))

	freeTenantTiers, err = pricingtier.ListPricingTiersForTenant(freeTenant.TenantId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(1, len(freeTenantTiers))

	err = pricingtier.RemovePricingTierFromTenant(freeTier.PricingTierId, freeTenant.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}

	freeTenantTiers, err = pricingtier.ListPricingTiersForTenant(freeTenant.TenantId, &warrant.ListPricingTierParams{
		ListParams: warrant.ListParams{
			Page:  1,
			Limit: 100,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.Equal(0, len(freeTenantTiers))

	// Clean up
	err = tenant.Delete(freeTenant.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = tenant.Delete(paidTenant.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pricingtier.Delete(freeTier.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = pricingtier.Delete(paidTier.PricingTierId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = feature.Delete(customFeature.FeatureId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = feature.Delete(feature1.FeatureId)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = feature.Delete(feature2.FeatureId)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestSessions(t *testing.T) {
	setup()
	assert := assert.New(t)

	user1, err := user.Create(&warrant.UserParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	tenant1, err := tenant.Create(&warrant.TenantParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = user.AssignUserToTenant(user1.UserId, tenant1.TenantId, "admin")
	if err != nil {
		fmt.Println(err)
		return
	}

	authzSessionToken, err := session.CreateAuthorizationSession(&warrant.AuthorizationSessionParams{
		UserId: user1.UserId,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.NotEmpty(authzSessionToken)

	ssDashUrl, err := session.CreateSelfServiceSession((&warrant.SelfServiceSessionParams{
		UserId:              user1.UserId,
		TenantId:            tenant1.TenantId,
		RedirectUrl:         "http://localhost:8080",
		SelfServiceStrategy: warrant.SelfServiceStrategyFGAC,
	}))
	if err != nil {
		fmt.Println(err)
		return
	}
	assert.NotEmpty(ssDashUrl)

	// Clean up
	err = user.Delete(user1.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = tenant.Delete(tenant1.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestWarrants(t *testing.T) {
	setup()
	assert := assert.New(t)

	newUser, err := user.Create(&warrant.UserParams{})
	if err != nil {
		fmt.Println(err)
		return
	}

	newPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "perm1",
		Name:         "Permission 1",
		Description:  "Permission with id 1",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	checkResult, err := warrant.Check(&warrant.WarrantCheckParams{
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
		fmt.Println(err)
		return
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
		fmt.Println(err)
		return
	}

	checkResult, err = warrant.Check(&warrant.WarrantCheckParams{
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
		fmt.Println(err)
		return
	}
	assert.True(checkResult)

	queryResult, err := warrant.Query(fmt.Sprintf("SELECT warrant FOR subject=%s:%s WHERE subject=%s:%s", "user", newUser.UserId, "user", newUser.UserId), &warrant.ListWarrantParams{})
	if err != nil {
		fmt.Println(err)
		return
	}
	queryBytes, _ := json.Marshal(queryResult.Result)
	var result []warrant.Warrant
	json.Unmarshal(queryBytes, &result)

	assert.Equal(1, len(result))
	assert.Equal("permission", result[0].ObjectType)
	assert.Equal("perm1", result[0].ObjectId)
	assert.Equal("member", result[0].Relation)

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
		fmt.Println(err)
		return
	}

	checkResult, err = warrant.Check(&warrant.WarrantCheckParams{
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
		fmt.Println(err)
		return
	}
	assert.False(checkResult)

	// Clean up
	err = user.Delete(newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = permission.Delete(newPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	}
}
