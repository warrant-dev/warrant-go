package main

import (
	"fmt"

	"github.com/warrant-dev/warrant-go"
)

const API_KEY = "YOUR_API_KEY"

func main() {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: API_KEY,
	})

	example(client)
}

func example(client warrant.WarrantClient) {
	// Create a new tenant with a Warrant generated id
	newTenant, err := client.CreateTenant(warrant.Tenant{
		TenantId: "test-tenant",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Create new tenant with generated id %s\n", newTenant.TenantId)
	}

	// Create a new tenant with a Warrant generated id
	fakeTenant, err := client.CreateTenant(warrant.Tenant{
		TenantId: "fake-tenant",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Create new tenant with generated id %s\n", fakeTenant.TenantId)
	}

	// Update tenant
	updatedTenant, err := client.UpdateTenant("test-tenant", warrant.Tenant{
		Name: "my tenant",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Updated tenant: %+v\n", updatedTenant)
	}

	// Get created tenant
	foundTenant, err := client.GetTenant("test-tenant")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found tenant: %+v\n", foundTenant)
	}

	// Create a new user with a Warrant generated id
	newUser, err := client.CreateUser(warrant.User{
		UserId: "test-user",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new user with generated id %s\n", newUser.UserId)
	}

	// Update user
	updatedUser, err := client.UpdateUser("test-user", warrant.User{
		Email: "my-user@example.com",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Updated user: %+v\n", updatedUser)
	}

	// Get created user
	foundUser, err := client.GetUser("test-user")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found user: %+v\n", foundUser)
	}

	// Assign the new user to the new tenant
	newWarrant, err := client.AssignUserToTenant(fakeTenant.TenantId, newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned user %s to tenant %s\n", newWarrant.Subject.ObjectId, newWarrant.ObjectId)
	}

	// Delete warrant
	err = client.DeleteWarrant(warrant.Warrant{
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
	createdWarrant, err := client.CreateWarrant(warrant.Warrant{
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
	testRole, err := client.CreateRole(warrant.Role{
		RoleId: "test-role",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new role %s\n", testRole.RoleId)
	}

	// Get created role
	foundRole, err := client.GetRole("test-role")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found role: %+v\n", foundRole)
	}

	// Create granted-permission
	assignedPermission, err := client.CreatePermission(warrant.Permission{
		PermissionId: "assigned-permission",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new permission %s\n", assignedPermission.PermissionId)
	}

	// Create user-specific-permission
	userAssignedPermission, err := client.CreatePermission(warrant.Permission{
		PermissionId: "user-specific-permission",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new permission %s\n", userAssignedPermission.PermissionId)
	}

	// Create unassigned-permission
	unassignedPermission, err := client.CreatePermission(warrant.Permission{
		PermissionId: "unassigned-permission",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new permission %s\n", unassignedPermission.PermissionId)
	}

	// Get permission
	foundPermission, err := client.GetPermission("assigned-permission")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found permission: %+v\n", foundPermission)
	}

	// Assign assigned-permission to test-role
	assignedPermission, err = client.AssignPermissionToRole(testRole.RoleId, assignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned permission %s to role %s\n", assignedPermission.PermissionId, testRole.RoleId)
	}

	// Assign view-self-service-dashboard to test-role
	_, err = client.AssignPermissionToRole(testRole.RoleId, "view-self-service-dashboard")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned permission %s to role %s\n", "view-self-service-dashboard", testRole.RoleId)
	}

	// Assign test-role to user
	_, err = client.AssignRoleToUser(newUser.UserId, testRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned role %s to user %s\n", testRole.RoleId, newWarrant.Subject.ObjectId)
	}

	// Assign user-specific-permission to user
	assignedUserPermission, err := client.AssignPermissionToUser(newUser.UserId, userAssignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned permission to user %s: %+v\n", newUser.UserId, assignedUserPermission)
	}

	// Create a session for user
	token, err := client.CreateAuthorizationSession(warrant.Session{
		UserId: newUser.UserId,
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created session token %s for user %s\n", token, newUser.UserId)
	}

	selfServiceDashUrl, err := client.CreateSelfServiceSession(warrant.Session{
		UserId:   newUser.UserId,
		TenantId: newTenant.TenantId,
	}, "https://warrant.dev")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created self service session %s for user %s\n", selfServiceDashUrl, newUser.UserId)
	}

	// List all warrants for organization
	warrants, err := client.ListWarrants(warrant.ListWarrantParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("\nWarrants: ")
		for _, w := range warrants {
			fmt.Printf("%+v\n", w)
		}
	}

	// List all tenants
	tenants, err := client.ListTenants(warrant.ListTenantParams{})
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
	tenantUsers, err := client.GetUsersForTenant(newTenant.TenantId)
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
	users, err := client.ListUsers(warrant.ListUserParams{})
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
	userTenants, err := client.GetTenantsForUser(newUser.UserId)
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
	userRoles, err := client.GetRolesForUser(newUser.UserId)
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
	userPermissions, err := client.GetPermissionsForUser(newUser.UserId)
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
	roles, err := client.ListRoles(warrant.ListRoleParams{})
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
	rolePermissions, err := client.GetPermissionsForRole(testRole.RoleId)
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
	permissions, err := client.ListPermissions(warrant.ListPermissionParams{})
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
	isAuthorized, err := client.IsAuthorized(warrant.WarrantCheckParams{
		Warrants: []warrant.Warrant{
			{
				ObjectType: "tenant",
				ObjectId:   newTenant.TenantId,
				Relation:   "member",
				Subject: warrant.Subject{
					ObjectType: "user",
					ObjectId:   newUser.UserId,
				},
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User %s is member of tenant %s? %t\n", newUser.UserId, newTenant.TenantId, isAuthorized)
	}

	// Check if user has assigned-permission
	hasPermission, err := client.HasPermission(warrant.PermissionCheckParams{
		PermissionId: assignedPermission.PermissionId,
		UserId:       newUser.UserId,
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User %s has permission %s? %t\n", newUser.UserId, assignedPermission.PermissionId, hasPermission)
	}

	// Check if user has unassigned-permission
	hasPermission, err = client.HasPermission(warrant.PermissionCheckParams{
		PermissionId: unassignedPermission.PermissionId,
		UserId:       newUser.UserId,
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User %s has permission %s? %t\n", newUser.UserId, unassignedPermission.PermissionId, hasPermission)
	}

	// Remove test-role from user
	err = client.RemoveRoleFromUser(newUser.UserId, testRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed role %s from user %s\n", testRole.RoleId, newWarrant.Subject.ObjectId)
	}

	// Remove view-self-service-dashboard from test-role
	err = client.RemovePermissionFromRole(testRole.RoleId, "view-self-service-dashboard")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed permission %s from role %s\n", "view-self-service-dashboard", testRole.RoleId)
	}

	// Remove assigned-permission from test-role
	err = client.RemovePermissionFromRole(testRole.RoleId, assignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed permission %s from role %s\n", assignedPermission.PermissionId, testRole.RoleId)
	}

	// Remove user-specific-permission from test-user
	err = client.RemovePermissionFromUser(newUser.UserId, userAssignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed permission %s from user %s\n", userAssignedPermission.PermissionId, newUser.UserId)
	}

	// Delete assigned-permission
	err = client.DeletePermission(assignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted permission %s\n", assignedPermission.PermissionId)
	}

	// Delete user-specific-permission
	err = client.DeletePermission(userAssignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted permission %s\n", userAssignedPermission.PermissionId)
	}

	// Delete unassigned-permission
	err = client.DeletePermission(unassignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted permission %s\n", unassignedPermission.PermissionId)
	}

	// Delete test-role
	err = client.DeleteRole(testRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted role %s\n", testRole.RoleId)
	}

	// Remove user from tenant
	err = client.RemoveUserFromTenant(newTenant.TenantId, newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed user %s from tenant %s\n", newUser.UserId, newTenant.TenantId)
	}

	// Delete user
	err = client.DeleteUser(newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted user %s\n", newUser.UserId)
	}

	// Delete tenant
	err = client.DeleteTenant(newTenant.TenantId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted tenant %s\n", newTenant.TenantId)
	}
}
