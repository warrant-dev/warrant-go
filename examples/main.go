package main

import (
	"fmt"

	"github.com/warrant-dev/warrant-go"
	"github.com/warrant-dev/warrant-go/permission"
	"github.com/warrant-dev/warrant-go/role"
	"github.com/warrant-dev/warrant-go/session"
	"github.com/warrant-dev/warrant-go/tenant"
	"github.com/warrant-dev/warrant-go/user"
)

func main() {
	warrant.ApiKey = "YOUR_KEY"

	example()
}

func example() {
	// Create a new tenant with a Warrant generated id
	newTenant, err := tenant.Create(&warrant.TenantParams{
		TenantId: "test-tenant",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Create new tenant with generated id %s\n", newTenant.TenantId)
	}

	// Create a new tenant with a Warrant generated id
	fakeTenant, err := tenant.Create(&warrant.TenantParams{
		TenantId: "fake-tenant",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Create new tenant with generated id %s\n", fakeTenant.TenantId)
	}

	// Update tenant
	updatedTenant, err := tenant.Update("test-tenant", &warrant.TenantParams{
		Name: "my tenant",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Updated tenant: %+v\n", updatedTenant)
	}

	// Get created tenant
	foundTenant, err := tenant.Get("test-tenant")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found tenant: %+v\n", foundTenant)
	}

	// Create a new user with a Warrant generated id
	newUser, err := user.Create(&warrant.UserParams{
		UserId: "test-user",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new user with generated id %s\n", newUser.UserId)
	}

	// Update user
	updatedUser, err := user.Update("test-user", &warrant.UserParams{
		Email: "my-user@example.com",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Updated user: %+v\n", updatedUser)
	}

	// Get created user
	foundUser, err := user.Get("test-user")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found user: %+v\n", foundUser)
	}

	// Assign the new user to the new tenant
	newWarrant, err := user.AssignUserToTenant(fakeTenant.TenantId, newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned user %s to tenant %s\n", newWarrant.Subject.ObjectId, newWarrant.ObjectId)
	}

	// Delete warrant
	err = warrant.Delete(&warrant.WarrantParams{
		ObjectType: "tenant",
		ObjectId:   fakeTenant.TenantId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: "user",
			ObjectId:   newUser.UserId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted warrant assigned user to tenant\n")
	}

	// Create warrant
	createdWarrant, err := warrant.Create(&warrant.WarrantParams{
		ObjectType: "tenant",
		ObjectId:   newTenant.TenantId,
		Relation:   "member",
		Subject: warrant.Subject{
			ObjectType: "user",
			ObjectId:   newUser.UserId,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created warrant: %+v\n", createdWarrant)
	}

	// Create test-role
	testRole, err := role.Create(&warrant.RoleParams{
		RoleId: "test-role",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new role %s\n", testRole.RoleId)
	}

	// Get created role
	foundRole, err := role.Get("test-role")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found role: %+v\n", foundRole)
	}

	// Create granted-permission
	assignedPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "assigned-permission",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new permission %s\n", assignedPermission.PermissionId)
	}

	// Create user-specific-permission
	userAssignedPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "user-specific-permission",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new permission %s\n", userAssignedPermission.PermissionId)
	}

	// Create unassigned-permission
	unassignedPermission, err := permission.Create(&warrant.PermissionParams{
		PermissionId: "unassigned-permission",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new permission %s\n", unassignedPermission.PermissionId)
	}

	// Get permission
	foundPermission, err := permission.Get("assigned-permission")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found permission: %+v\n", foundPermission)
	}

	// Assign assigned-permission to test-role
	assignedPermission, err = permission.AssignPermissionToRole(assignedPermission.PermissionId, testRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned permission %s to role %s\n", assignedPermission.PermissionId, testRole.RoleId)
	}

	// Assign view-self-service-dashboard to test-role
	_, err = permission.AssignPermissionToRole("view-self-service-dashboard", testRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned permission %s to role %s\n", "view-self-service-dashboard", testRole.RoleId)
	}

	// Assign test-role to user
	_, err = role.AssignRoleToUser(newUser.UserId, testRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned role %s to user %s\n", testRole.RoleId, newWarrant.Subject.ObjectId)
	}

	// Assign user-specific-permission to user
	assignedUserPermission, err := permission.AssignPermissionToUser(userAssignedPermission.PermissionId, newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned permission to user %s: %+v\n", newUser.UserId, assignedUserPermission)
	}

	// Create a session for user
	token, err := session.CreateAuthorizationSession(&warrant.AuthorizationSessionParams{
		UserId: newUser.UserId,
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created session token %s for user %s\n", token, newUser.UserId)
	}

	selfServiceDashUrl, err := session.CreateSelfServiceSession(&warrant.SelfServiceSessionParams{
		UserId:      newUser.UserId,
		TenantId:    newTenant.TenantId,
		RedirectUrl: "https://warrant.dev",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created self service session %s for user %s\n", selfServiceDashUrl, newUser.UserId)
	}

	// Query warrants for role test-role
	testRoleQuery, err := warrant.Query("SELECT warrant FOR object=role:test-role")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("\nRole test-role's warrants: %+v", testRoleQuery.Result)
	}

	// List all tenants
	tenants, err := tenant.ListTenants(&warrant.ListTenantParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("\nTenants: ")
		for _, w := range tenants {
			fmt.Printf("%+v\n", w)
		}
	}

	// List users for tenant
	tenantUsers, err := user.ListUsersForTenant(newTenant.TenantId, &warrant.ListUserParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("\nUsers for tenant %s: \n", newTenant.TenantId)
		for _, w := range tenantUsers {
			fmt.Printf("%+v\n", w)
		}
	}

	// List all users
	users, err := user.ListUsers(&warrant.ListUserParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("\nUsers: ")
		for _, w := range users {
			fmt.Printf("%+v\n", w)
		}
	}

	// Get tenants for user
	userTenants, err := tenant.ListTenantsForUser(newUser.UserId, &warrant.ListTenantParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("\nTenants for user %s: \n", newUser.UserId)
		for _, w := range userTenants {
			fmt.Printf("%+v\n", w)
		}
	}

	// Get roles for user
	userRoles, err := role.ListRolesForUser(newUser.UserId, &warrant.ListRoleParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("\nRoles for user %s: \n", newUser.UserId)
		for _, w := range userRoles {
			fmt.Printf("%+v\n", w)
		}
	}

	// Get permissions for user
	userPermissions, err := permission.ListPermissionsForUser(newUser.UserId, &warrant.ListPermissionParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("\nPermissions for user %s: \n", newUser.UserId)
		for _, w := range userPermissions {
			fmt.Printf("%+v\n", w)
		}
	}

	// List all roles
	roles, err := role.ListRoles(&warrant.ListRoleParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("\nRoles: ")
		for _, w := range roles {
			fmt.Printf("%+v\n", w)
		}
	}

	// Get permissions for role
	rolePermissions, err := permission.ListPermissionsForRole(testRole.RoleId, &warrant.ListPermissionParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("\nPermissions for role %s: \n", testRole.RoleId)
		for _, w := range rolePermissions {
			fmt.Printf("%+v\n", w)
		}
	}

	// List all permissions
	permissions, err := permission.ListPermissions(&warrant.ListPermissionParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("\nPermissions: ")
		for _, w := range permissions {
			fmt.Printf("%+v\n", w)
		}
	}

	// Check authorization (tenant membership) for a user
	isAuthorized, err := warrant.Check(&warrant.WarrantCheckParams{
		Object: &warrant.WarrantObject{
			ObjectType: "tenant",
			ObjectId:   newTenant.TenantId,
		},
		Relation: "member",
		Subject: &warrant.Subject{
			ObjectType: "user",
			ObjectId:   newUser.UserId,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User %s is member of tenant %s? %t\n", newUser.UserId, newTenant.TenantId, isAuthorized)
	}

	// Check if user has assigned-permission
	hasPermission, err := warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		PermissionId: assignedPermission.PermissionId,
		UserId:       newUser.UserId,
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User %s has permission %s? %t\n", newUser.UserId, assignedPermission.PermissionId, hasPermission)
	}

	// Check if user has unassigned-permission
	hasPermission, err = warrant.CheckUserHasPermission(&warrant.PermissionCheckParams{
		PermissionId: unassignedPermission.PermissionId,
		UserId:       newUser.UserId,
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User %s has permission %s? %t\n", newUser.UserId, unassignedPermission.PermissionId, hasPermission)
	}

	// Remove test-role from user
	err = role.RemoveRoleFromUser(testRole.RoleId, newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed role %s from user %s\n", testRole.RoleId, newWarrant.Subject.ObjectId)
	}

	// Remove view-self-service-dashboard from test-role
	err = permission.RemovePermissionFromRole("view-self-service-dashboard", testRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed permission %s from role %s\n", "view-self-service-dashboard", testRole.RoleId)
	}

	// Remove assigned-permission from test-role
	err = permission.RemovePermissionFromRole(assignedPermission.PermissionId, testRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed permission %s from role %s\n", assignedPermission.PermissionId, testRole.RoleId)
	}

	// Remove user-specific-permission from test-user
	err = permission.RemovePermissionFromUser(userAssignedPermission.PermissionId, newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed permission %s from user %s\n", userAssignedPermission.PermissionId, newUser.UserId)
	}

	// Delete assigned-permission
	err = permission.Delete(assignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted permission %s\n", assignedPermission.PermissionId)
	}

	// Delete user-specific-permission
	err = permission.Delete(userAssignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted permission %s\n", userAssignedPermission.PermissionId)
	}

	// Delete unassigned-permission
	err = permission.Delete(unassignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted permission %s\n", unassignedPermission.PermissionId)
	}

	// Delete test-role
	err = role.Delete(testRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted role %s\n", testRole.RoleId)
	}

	// Remove user from tenant
	err = user.RemoveUserFromTenant(newTenant.TenantId, newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed user %s from tenant %s\n", newUser.UserId, newTenant.TenantId)
	}

	// Delete user
	err = user.Delete(newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted user %s\n", newUser.UserId)
	}

	// Delete tenant
	err = tenant.Delete(newTenant.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted tenant %s\n", newTenant.TenantId)
	}

	err = tenant.Delete(fakeTenant.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted tenant %s\n", newTenant.TenantId)
	}
}
