package main

import (
	"fmt"
	"time"

	"github.com/warrant-dev/warrant-go"
)

const API_KEY = "YOUR_API_KEY"

func main() {
	client := warrant.NewClient(warrant.ClientConfig{
		ApiKey: API_KEY,
	})

	createUserAndSession(client)
	createWarrants(client)
	checkAuthorization(client)
}

func createUserAndSession(client warrant.WarrantClient) {
	// Create a new user with Warrant generated IDs
	newUser, err := client.CreateUserWithGeneratedId()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created new user with generated id " + newUser.UserId)
	}

	// Create a session for newly created user
	token, err := client.CreateSession(newUser.UserId)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Created session token for new user " + token)
	}

	// Create a new user with provided IDs
	newUser, err = client.CreateUser(warrant.User{
		UserId: fmt.Sprintf("newUser%d", time.Now().Unix()),
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Created new user with provided id " + newUser.UserId)
	}
}

func createWarrants(client warrant.WarrantClient) {
	// Create a warrant that establishes "userId" as a member of "store1"
	// Note: object types for store must exist
	resp, err := client.CreateWarrant(warrant.Warrant{
		ObjectType: "store",
		ObjectId:   "store1",
		Relation:   "member",
		User: warrant.WarrantUser{
			UserId: "user1",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}

	// Create a warrant on a userset (instead of specific userId)
	resp, err = client.CreateWarrant(warrant.Warrant{
		ObjectType: "report",
		ObjectId:   "report_45",
		Relation:   "viewer",
		User: warrant.WarrantUser{
			Userset: &warrant.Userset{
				ObjectType: "store",
				ObjectId:   "store1",
				Relation:   "member",
			},
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(resp)
	}
}

func checkAuthorization(client warrant.WarrantClient) {
	// Check authorization (store membership) for a user
	isAuthorized, err := client.IsAuthorized(warrant.Warrant{
		ObjectType: "store",
		ObjectId:   "store1",
		Relation:   "member",
		User: warrant.WarrantUser{
			UserId: "user1",
		},
	})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(isAuthorized)
	}
}
