# Warrant Go Library

Use [Warrant](https://warrant.dev/) in server-side Go projects.

[![Slack](https://img.shields.io/badge/slack-join-brightgreen)](https://join.slack.com/t/warrantcommunity/shared_invite/zt-12g84updv-5l1pktJf2bI5WIKN4_~f4w)

## Installation

```shell
go get github.com/warrant-dev/warrant-go
```

## Usage

Instantiate the Warrant client with your API key to get started:
```go
import "github.com/warrant-dev/warrant-go"

client := warrant.NewClient(warrant.ClientConfig{
    ApiKey: "api_test_f5dsKVeYnVSLHGje44zAygqgqXiLJBICbFzCiAg1E=",
})
```

### `CreateUserWithGeneratedId()`

This method creates a user entity in Warrant with a Warrant-generated id.
```go
user, err := client.CreateUserWithGeneratedId()
```

### `CreateUser(user User)`

This method creates a user entity in Warrant with the specified userId.
```go
user, err := client.CreateUser(warrant.User{
    UserId: "userId",
})
```

### `CreateWarrant(warrantToCreate Warrant)`

This method creates a warrant which specifies that the provided `user` (or `userset`) has `relation` on the object of type `objectType` with id `objectId`.
```go
// Create a warrant allowing user1 to "view" the store with id store1
warrant, err := client.createWarrant(warrant.Warrant{
		ObjectType: "store",
		ObjectId:   "store1",
		Relation:   "viewer",
		User: warrant.WarrantUser{
			UserId: "user1",
		},
	})
```

### `CreateSession(userId string)`

This method creates a session in Warrant for the user with the specified `userId` and returns a session token which can be used to make authorized requests to the Warrant API only for the specified user. This session token can safely be used to make requests to the Warrant API's authorization endpoint to determine user access in web and mobile client applications.

```go
// Creates a session token scoped to the specified userId
// Return this token to your client application to allow
// it to make requests for the given user.
token, err := client.CreateSession(userId)
```

### `IsAuthorized(warrant Warrant)`

This method returns `true` or `false` depending on whether the user with the specified `userId` has the specified `relation` to the object of type `objectType` with id `objectId` and `false` otherwise.

```go
//
// Example Scenario:
// An e-commerce website where Store Owners can edit store info
//
isAuthorized, err := client.IsAuthorized(warrant.Warrant{
		ObjectType: "store",
		ObjectId:   "store1",
		Relation:   "editor",
		User: warrant.WarrantUser{
			UserId: "user1", // store owner
		},
	})
```

We’ve used a random API key in these code examples. Replace it with your
[actual publishable API keys](https://app.warrant.dev) to
test this code through your own Warrant account.

For more information on how to use the Warrant API, please refer to the
[Warrant API reference](https://docs.warrant.dev).

Note that we may release new [minor and patch](https://semver.org/) versions of this library with small but backwards-incompatible fixes to the type declarations. These changes will not affect Warrant itself.

## Warrant Documentation

- [Warrant Docs](https://docs.warrant.dev/)
