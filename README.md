# Warrant Go Library

Use [Warrant](https://warrant.dev/) in server-side Go projects.

[![Slack](https://img.shields.io/badge/slack-join-brightgreen)](https://join.slack.com/t/warrantcommunity/shared_invite/zt-12g84updv-5l1pktJf2bI5WIKN4_~f4w)

## Installation

```shell
go get github.com/warrant-dev/warrant-go
```

## Usage

You can use the Warrant SDK with or without a client. Instantiating a client allows you to create different client instances each with their own config (API key, API endpoint, etc).

### Without a Client

```go
import "github.com/warrant-dev/warrant-go"

// Setup
warrant.ApiKey = "api_test_f5dsKVeYnVSLHGje44zAygqgqXiLJBICbFzCiAg1E="

// Create warrant
warrant, err := warrant.Create(&warrant.WarrantParams{})

// Create tenant
tenant, err := tenant.Create(&tenant.TenantParams{})
```

### With a Client

Instantiate the Warrant client with your API key to get started:
```go
import "github.com/warrant-dev/warrant-go"

client := warrant.NewClient(config.ClientConfig{
    ApiKey: "api_test_f5dsKVeYnVSLHGje44zAygqgqXiLJBICbFzCiAg1E=",
	ApiEndpoint: "https://api.warrant.dev",
	AuthorizeEndpoint: "https://api.warrant.dev",
	SelfServiceDashEndpoint: "https://self-serve.warrant.dev"
})
```

## Examples

### Users

```go
// Create
createdUser, err := user.Create(&warrant.UserParams{
    UserId: "userId",
})

// Get
user, err := user.Get("userId")


// Delete
err = user.Delete("userId")
```

### Warrants

```go

// Create
createdWarrant, err := warrant.Create(&warrant.WarrantParams{
	ObjectType: "tenant",
	ObjectId:   "1",
	Relation:   "member",
	Subject: warrant.Subject{
		ObjectType: "user",
		ObjectId:   "1",
	},
})

// Delete
err = warrant.Delete(&warrant.WarrantParams{
	ObjectType: "tenant",
	ObjectId:   "1",
	Relation:   "member",
	Subject: warrant.Subject{
		ObjectType: "user",
		ObjectId:   "1",
	},
})

// Check access
isAuthorized, err := warrant.Check(&warrant.WarrantCheckParams{
	Object: &warrant.WarrantObject{
		ObjectType: "tenant",
		ObjectId:   "1",
	},
	Relation: "member",
	Subject: &warrant.Subject{
		ObjectType: "user",
		ObjectId:   "1",
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
