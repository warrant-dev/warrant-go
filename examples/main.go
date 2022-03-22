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

	// Create a session for newly created user
	token, err := client.CreateAuthorizationSession(warrant.Session{
		UserId:   newUser.UserId,
		TenantId: newTenant.TenantId,
	})
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Created session token for new user %s\n", token)
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
		fmt.Println(isAuthorized)
	}
}
