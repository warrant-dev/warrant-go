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
	newTenant, err := client.CreateTenant(warrant.Tenant{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Create new tenant with generated id %s\n", newTenant.TenantId)
	}

	// Create a new user with a Warrant generated id
	newUser, err := client.CreateUser(warrant.User{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new user with generated id %s\n", newUser.UserId)
	}

	// Assign the new user to the new tenant
	newWarrant, err := client.AssignUserToTenant(newTenant.TenantId, newUser.UserId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned user %s to tenant %s\n", newWarrant.User.UserId, newWarrant.ObjectId)
	}

	// Create test-role
	testRole, err := client.CreateRole("test-role")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new role %s\n", testRole.RoleId)
	}

	// Create granted-permission
	assignedPermission, err := client.CreatePermission("assigned-permission")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new permission %s\n", assignedPermission.PermissionId)
	}

	// Create unassigned-permission
	unassignedPermission, err := client.CreatePermission("unassigned-permission")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created new permission %s\n", unassignedPermission.PermissionId)
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
		fmt.Printf("Assigned role %s to user %s\n", testRole.RoleId, newWarrant.User.UserId)
	}

	// Create a session for user
	token, err := client.CreateAuthorizationSession(warrant.Session{
		UserId:   newUser.UserId,
		TenantId: newTenant.TenantId,
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

	// Check authorization (tenant membership) for a user
	isAuthorized, err := client.IsAuthorized(warrant.Warrant{
		ObjectType: "tenant",
		ObjectId:   newTenant.TenantId,
		Relation:   "member",
		User: warrant.WarrantUser{
			UserId: newUser.UserId,
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User %s is member of tenant %s? %t", newUser.UserId, newTenant.TenantId, isAuthorized)
	}

	// Check if user has assigned-permission
	hasPermission, err := client.HasPermission(assignedPermission.PermissionId, newUser.UserId)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User %s has permission %s? %t", newUser.UserId, assignedPermission.PermissionId, hasPermission)
	}

	// Check if user has unassigned-permission
	hasPermission, err = client.HasPermission(unassignedPermission.PermissionId, newUser.UserId)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("User %s has permission %s? %t", newUser.UserId, unassignedPermission.PermissionId, hasPermission)
	}

	// Remove test-role from user
	err = client.RemoveRoleFromUser(newUser.UserId, testRole.RoleId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Removed role %s from user %s\n", testRole.RoleId, newWarrant.User.UserId)
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

	// Delete assigned-permission
	err = client.DeletePermission(assignedPermission.PermissionId)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Deleted permission %s\n", assignedPermission.PermissionId)
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
