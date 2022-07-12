package main

import (
	"fmt"

	warrantdev "github.com/warrant-dev/warrant-go"
	"github.com/warrant-dev/warrant-go/permission"
	"github.com/warrant-dev/warrant-go/role"
	"github.com/warrant-dev/warrant-go/session"
	"github.com/warrant-dev/warrant-go/tenant"
	"github.com/warrant-dev/warrant-go/user"
	"github.com/warrant-dev/warrant-go/warrant"
)

const API_KEY = "YOUR_API_KEY"

func main() {
	warrantdev.ApiKey = "api_prod_fgozJoE0i8cNb6gCCgWFlZBHSjTxUInf81Q1x1zXGAc="

	// Roles
	fmt.Println("ROLES:")
	newRole, err := role.New(&role.RoleParams{RoleId: "go-role"})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("New role: %+v\n", newRole)
	}

	newRole, err = role.New(&role.RoleParams{RoleId: "go-role-2"})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("New role: %+v\n", newRole)
	}

	foundRole, err := role.Get("go-role")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found role go-role: %+v\n", foundRole)
	}

	roles, err := role.List(&role.RoleListParams{
		ListParams: warrantdev.ListParams{
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("All roles:")
		for _, r := range roles {
			fmt.Printf("%+v\n", r)
		}
	}

	err = role.Delete("go-role-2")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Role go-role-2 deleted")
	}

	roles, err = role.List(nil)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("All roles:")
		for _, r := range roles {
			fmt.Printf("%+v\n", r)
		}
	}

	// Permissions
	fmt.Println("\nPERMISSIONS:")
	newPermission, err := permission.New(&permission.PermissionParams{PermissionId: "go-permission"})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("New permission: %+v\n", newPermission)
	}

	newPermission, err = permission.New(&permission.PermissionParams{PermissionId: "go-permission-2"})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("New permission: %+v\n", newPermission)
	}

	foundPermission, err := permission.Get("go-permission")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found permission go-permission: %+v\n", foundPermission)
	}

	permissions, err := permission.List(nil)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("All permissions:")
		for _, r := range permissions {
			fmt.Printf("%+v\n", r)
		}
	}

	err = permission.Delete("go-permission-2")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Permission go-permission-2 deleted")
	}

	permissions, err = permission.List(&permission.PermissionListParams{
		ListParams: warrantdev.ListParams{
			Limit: 10,
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("All permissions:")
		for _, r := range permissions {
			fmt.Printf("%+v\n", r)
		}
	}

	// Users
	fmt.Println("\nUSERS:")
	newUser, err := user.New(&user.UserParams{UserId: "go-user"})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("New user: %+v\n", newUser)
	}

	newUser, err = user.New(&user.UserParams{UserId: "go-user-2"})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("New user: %+v\n", newUser)
	}

	newUser, err = user.New(&user.UserParams{UserId: "go-user-3"})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("New user: %+v\n", newUser)
	}

	updatedUser, err := user.Update("go-user-2", &user.UserParams{
		Email: "updated-email@example.com",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Updated user: %+v\n", updatedUser)
	}

	foundUser, err := user.Get("go-user")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found user go-user: %+v\n", foundUser)
	}

	users, err := user.List(&user.UserListParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("All users:")
		for _, r := range users {
			fmt.Printf("%+v\n", r)
		}
	}

	err = user.Delete("go-user-2")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("User go-user-2 deleted")
	}

	users, err = user.List(nil)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("All users:")
		for _, r := range users {
			fmt.Printf("%+v\n", r)
		}
	}

	assignedRole, err := foundUser.AssignRole("admin")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned role to user go-user: %+v\n", assignedRole)
	}

	assignedRole, err = foundUser.AssignRole("member")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned role to user go-user: %+v\n", assignedRole)
	}

	err = foundUser.RemoveRole("member")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Removed role member from user go-user")
	}

	userRoles, err := foundUser.ListRoles()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("All roles for user go-user:")
		for _, r := range userRoles {
			fmt.Printf("%+v\n", r)
		}
	}

	assignedPermission, err := foundUser.AssignPermission("go-permission")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned permission to user go-user: %+v\n", assignedPermission)
	}

	assignedPermission, err = foundUser.AssignPermission("permission1")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned permission to user go-user: %+v\n", assignedPermission)
	}

	// hasPermission, err := foundUser.HasPermission("go-permission")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("User go-user has permission go-permission? %+v\n", hasPermission)
	// }

	// hasPermission, err = foundUser.HasPermission("fake-permission")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("User go-user has permission fake-permission? %+v\n", hasPermission)
	// }

	err = foundUser.RemovePermission("permission1")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Removed permission permission1")
	}

	userPermissions, err := foundUser.ListPermissions()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("User go-user's permissions:")
		for _, r := range userPermissions {
			fmt.Printf("%+v\n", r)
		}
	}

	// Tenants
	fmt.Println("\nTENANTS:")
	newTenant, err := tenant.New(&tenant.TenantParams{TenantId: "go-tenant"})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("New tenant: %+v\n", newTenant)
	}

	newTenant, err = tenant.New(&tenant.TenantParams{TenantId: "go-tenant-2"})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("New tenant: %+v\n", newTenant)
	}

	foundTenant, err := tenant.Get("go-tenant")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Found tenant: %+v\n", foundTenant)
	}

	updatedTenant, err := tenant.Update("go-tenant-2", &tenant.TenantParams{
		Name: "updated tenant name",
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Updated tenant: %+v\n", updatedTenant)
	}

	tenants, err := tenant.List(&tenant.TenantListParams{})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("All tenants:")
		for _, r := range tenants {
			fmt.Printf("%+v\n", r)
		}
	}

	err = tenant.Delete("go-tenant-2")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Tenant go-tenant-2 deleted")
	}

	tenants, err = tenant.List(nil)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("All tenants:")
		for _, r := range tenants {
			fmt.Printf("%+v\n", r)
		}
	}

	// Users/Tenants
	fmt.Println("\nUSERS/TENANTS:")
	newWarrant, err := foundTenant.AddUser("go-user")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned go-user to tenant: %+v", newWarrant)
	}

	newWarrant, err = foundTenant.AddUser("go-user-3")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Assigned go-user-3 to tenant: %+v", newWarrant)
	}

	userTenants, err := foundUser.ListTenants()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("User's tenants:")
		for _, r := range userTenants {
			fmt.Printf("%+v\n", r)
		}
	}

	tenantUsers, err := foundTenant.ListUsers()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Tenant's users:")
		for _, r := range tenantUsers {
			fmt.Printf("%+v\n", r)
		}
	}

	err = foundTenant.RemoveUser("go-user-3")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Removed go-user-3 from tenant")
	}

	// Sessions
	fmt.Println("\nSESSIONS:")
	token, err := session.New(&session.SessionParams{
		Type:     "sess",
		UserId:   foundUser.UserId,
		TenantId: foundTenant.TenantId,
		TTL:      3600,
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created session: %+v", token)
	}

	token, err = session.New(&session.SessionParams{
		Type:     "ssdash",
		UserId:   foundUser.UserId,
		TenantId: foundTenant.TenantId,
		TTL:      3600,
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created self service session: %+v", token)
	}

	// Warrants
	fmt.Println("\nWARRANTS:")

	warrants, err := warrant.List(nil)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("All warrants:")
		for _, r := range warrants {
			fmt.Printf("%+v\n", r)
		}
	}

	newWarrant, err = warrant.New(&warrant.WarrantParams{
		ObjectType: "permission",
		ObjectId:   "go-permission",
		Relation:   "member",
		Subject: &warrant.SubjectParams{
			ObjectType: "user",
			ObjectId:   "go-user",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("New warrant: %+v\n", newWarrant)
	}

	updatedWarrant, err := warrant.Update(&warrant.WarrantUpdateParams{
		Before: &warrant.WarrantParams{
			ObjectType: "permission",
			ObjectId:   "go-permission",
			Relation:   "member",
			Subject: &warrant.SubjectParams{
				ObjectType: "user",
				ObjectId:   "go-user",
			},
		},
		After: &warrant.WarrantParams{
			ObjectType: "permission",
			ObjectId:   "go-permission",
			Relation:   "member",
			Subject: &warrant.SubjectParams{
				ObjectType: "user",
				ObjectId:   "go-user-3",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Updated warrant: %+v\n", updatedWarrant)
	}

	// isAuthorized, err := warrant.IsAuthorized(&warrant.WarrantCheckParams{
	// 	Warrants: []*warrant.WarrantParams{
	// 		{
	// 			ObjectType: "permission",
	// 			ObjectId:   "go-permission",
	// 			Relation:   "member",
	// 			Subject: &warrant.SubjectParams{
	// 				ObjectType: "user",
	// 				ObjectId:   "go-user-nonexistent",
	// 			},
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Is authorized?: %+v\n", isAuthorized)
	// }

	// isAuthorized, err = warrant.IsAuthorized(&warrant.WarrantCheckParams{
	// 	Warrants: []*warrant.WarrantParams{
	// 		{
	// 			ObjectType: "permission",
	// 			ObjectId:   "go-permission",
	// 			Relation:   "member",
	// 			Subject: &warrant.SubjectParams{
	// 				ObjectType: "user",
	// 				ObjectId:   "go-user-3",
	// 			},
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Is authorized?: %+v\n", isAuthorized)
	// }

	// hasPermission, err = warrant.HasPermission(&warrant.PermissionCheckParams{
	// 	PermissionId: "go-permission",
	// 	UserId: "go-user",
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("User go-user has permission go-permission? %+v\n", hasPermission)
	// }

	err = warrant.Delete(&warrant.WarrantParams{
		ObjectType: "tenant",
		ObjectId:   "go-tenant",
		Relation:   "member",
		Subject: &warrant.SubjectParams{
			ObjectType: "user",
			ObjectId:   "go-user-3",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Deleted warrant")
	}
}

func example(client warrantdev.WarrantClient) {
	// // Create a new tenant with a Warrant generated id
	// newTenant, err := client.CreateTenant(warrant.Tenant{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Create new tenant with generated id %s\n", newTenant.TenantId)
	// }

	// // Create a new user with a Warrant generated id
	// newUser, err := client.CreateUser(warrant.User{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Created new user with generated id %s\n", newUser.UserId)
	// }

	// // Assign the new user to the new tenant
	// newWarrant, err := client.AssignUserToTenant(newTenant.TenantId, newUser.UserId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Assigned user %s to tenant %s\n", newWarrant.Subject.ObjectId, newWarrant.ObjectId)
	// }

	// // Create test-role
	// testRole, err := client.CreateRole("test-role")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Created new role %s\n", testRole.RoleId)
	// }

	// // Create granted-permission
	// assignedPermission, err := client.CreatePermission("assigned-permission")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Created new permission %s\n", assignedPermission.PermissionId)
	// }

	// // Create unassigned-permission
	// unassignedPermission, err := client.CreatePermission("unassigned-permission")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Created new permission %s\n", unassignedPermission.PermissionId)
	// }

	// // Assign assigned-permission to test-role
	// assignedPermission, err = client.AssignPermissionToRole(testRole.RoleId, assignedPermission.PermissionId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Assigned permission %s to role %s\n", assignedPermission.PermissionId, testRole.RoleId)
	// }

	// // Assign view-self-service-dashboard to test-role
	// _, err = client.AssignPermissionToRole(testRole.RoleId, "view-self-service-dashboard")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Assigned permission %s to role %s\n", "view-self-service-dashboard", testRole.RoleId)
	// }

	// // Assign test-role to user
	// _, err = client.AssignRoleToUser(newUser.UserId, testRole.RoleId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Assigned role %s to user %s\n", testRole.RoleId, newWarrant.Subject.ObjectId)
	// }

	// // Create a session for user
	// token, err := client.CreateAuthorizationSession(warrant.Session{
	// 	UserId:   newUser.UserId,
	// 	TenantId: newTenant.TenantId,
	// })
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Created session token %s for user %s\n", token, newUser.UserId)
	// }

	// selfServiceDashUrl, err := client.CreateSelfServiceSession(warrant.Session{
	// 	UserId:   newUser.UserId,
	// 	TenantId: newTenant.TenantId,
	// }, "https://warrant.dev")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Created self service session %s for user %s\n", selfServiceDashUrl, newUser.UserId)
	// }

	// // List all warrants for organization
	// warrants, err := client.ListWarrants(warrant.ListWarrantFilters{})
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Println("Warrants: ")
	// 	for _, w := range warrants {
	// 		fmt.Printf("%+v\n", w)
	// 	}
	// }

	// // Check authorization (tenant membership) for a user
	// isAuthorized, err := client.IsAuthorized(warrant.Warrant{
	// 	ObjectType: "tenant",
	// 	ObjectId:   newTenant.TenantId,
	// 	Relation:   "member",
	// 	Subject: warrant.Subject{
	// 		ObjectType: "user",
	// 		ObjectId:   newUser.UserId,
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("User %s is member of tenant %s? %t", newUser.UserId, newTenant.TenantId, isAuthorized)
	// }

	// // Check if user has assigned-permission
	// hasPermission, err := client.HasPermission(assignedPermission.PermissionId, newUser.UserId)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("User %s has permission %s? %t", newUser.UserId, assignedPermission.PermissionId, hasPermission)
	// }

	// // Check if user has unassigned-permission
	// hasPermission, err = client.HasPermission(unassignedPermission.PermissionId, newUser.UserId)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else {
	// 	fmt.Printf("User %s has permission %s? %t", newUser.UserId, unassignedPermission.PermissionId, hasPermission)
	// }

	// // Remove test-role from user
	// err = client.RemoveRoleFromUser(newUser.UserId, testRole.RoleId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Removed role %s from user %s\n", testRole.RoleId, newWarrant.Subject.ObjectId)
	// }

	// // Remove view-self-service-dashboard from test-role
	// err = client.RemovePermissionFromRole(testRole.RoleId, "view-self-service-dashboard")
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Removed permission %s from role %s\n", "view-self-service-dashboard", testRole.RoleId)
	// }

	// // Remove assigned-permission from test-role
	// err = client.RemovePermissionFromRole(testRole.RoleId, assignedPermission.PermissionId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Removed permission %s from role %s\n", assignedPermission.PermissionId, testRole.RoleId)
	// }

	// // Delete assigned-permission
	// err = client.DeletePermission(assignedPermission.PermissionId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Deleted permission %s\n", assignedPermission.PermissionId)
	// }

	// // Delete unassigned-permission
	// err = client.DeletePermission(unassignedPermission.PermissionId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Deleted permission %s\n", unassignedPermission.PermissionId)
	// }

	// // Delete test-role
	// err = client.DeleteRole(testRole.RoleId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Deleted role %s\n", testRole.RoleId)
	// }

	// // Remove user from tenant
	// err = client.RemoveUserFromTenant(newTenant.TenantId, newUser.UserId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Removed user %s from tenant %s\n", newUser.UserId, newTenant.TenantId)
	// }

	// // Delete user
	// err = client.DeleteUser(newUser.UserId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Deleted user %s\n", newUser.UserId)
	// }

	// // Delete tenant
	// err = client.DeleteTenant(newTenant.TenantId)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// } else {
	// 	fmt.Printf("Deleted tenant %s\n", newTenant.TenantId)
	// }
}
