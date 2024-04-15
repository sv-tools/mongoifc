# mongoifc

[![Code Analysis](https://github.com/sv-tools/mongoifc/actions/workflows/checks.yaml/badge.svg)](https://github.com/sv-tools/mongoifc/actions/workflows/checks.yaml)
[![Go Reference](https://pkg.go.dev/badge/github.com/sv-tools/mongoifc.svg)](https://pkg.go.dev/github.com/sv-tools/mongoifc)
[![codecov](https://codecov.io/gh/sv-tools/mongoifc/branch/main/graph/badge.svg?token=0XVOTDR1CW)](https://codecov.io/gh/sv-tools/mongoifc)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/v/tag/sv-tools/mongoifc?style=flat)](https://github.com/sv-tools/mongoifc/releases)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/8803/badge)](https://www.bestpractices.dev/projects/8803)
[![OpenSSF Scorecard](https://api.securityscorecards.dev/projects/github.com/{owner}/{repo}/badge)](https://securityscorecards.dev/viewer/?uri=github.com/sv-tools/mongoifc)

The Interfaces for the MongoDB driver

## Versioning Policy

The `mongoifc` code is stabilized, so now the version will match the version of the MongoDB driver since `v1.8.0`.

In case of need for bug fixes in `mongoifc`, the version will be in this format `v1.8.1+N`, where `v1.8.1` is the
version of MongoDB driver and `N` is a patch of `mongoifc`. The new version for changes in README.md, tests, examples,
GitHub workflows is not required.

## :bangbang: **Important**

It is not a simple drop in replacement because of the limitations in Go. You should rewrite your code to use this
library instead of mongo driver.

```go
conn := mongoifc.Connect(...)
```

instead of

```go
conn := mongo.Connect(...)
```

or if you have a special code that returns the mongo object then you need to use one of `Wrap` functions to wrap
the `mongo.Client` or `mongo.Database` or `mongo.Collection` or `mongo.Session` objects:

```go
orig := mongo.Connect(...)

...

conn := mongoifc.WrapClient(orig)
```

or

```go
func GetTenantDB(ctx context.Context, tenantID, dbName string) (*mongo.Database, error) {
// a code that returns a special database for a given tenant and database name
}

...

orig, err := GetTenantDB(ctx, tenant, "users")
if err != nil {
...
}
db = mongoifc.WrapDatabase(orig)
```

Now let's dig a bit into the limitations. Assume that you have a function to return a list of admin users, and you
rewrote it using `mongoifc`:

```go
package users

// Original: func GetAdmins(ctx context.Context, db *mongo.Database) ([]*User, error)
func GetAdmins(ctx context.Context, db mongoifc.Database) ([]User, error) {
	var users []User
	cur, err := db.Collection(UsersCollection).Find(ctx, User{
		Active:  true,
		IsAdmin: true,
	})
	if err != nil {
		return nil, err
	}
	if err := cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, err
}
```

and if you pass an object of `*mongo.Database` type instead of `mongoifc.Database`

```go
conn, _ := mongo.Connect(context.Background(), ...)
db := conn.Database(...)

users.GetAdmins(context.Background(), db)
```

then compilation fails with such error:

     cannot use db (type *mongo.Database) as type mongoifc.Database in argument to simple.GetAdmins:
         *mongo.Database does not implement mongoifc.Database (wrong type for Aggregate method)
             have Aggregate(context.Context, interface {}, ...*"go.mongodb.org/mongo-driver/mongo/options".AggregateOptions) (*mongo.Cursor, error)
             want Aggregate(context.Context, interface {}, ...*"go.mongodb.org/mongo-driver/mongo/options".AggregateOptions) (mongoifc.Cursor, error)

This is the main reason of wrapping the original objects and using the `mongoifc` instead.

## Wrapped Interfaces

- [x] Client: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Client
- [x] Database: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Database
- [x] Session: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Session
- [x] ChangeStream: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#ChangeStream
- [x] Cursor: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Cursor
- [x] Collection: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#Collection
- [x] SingleResult: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#SingleResult
- [x] IndexView: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#IndexView
- [x] SessionContext: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#SessionContext
- [x] ClientEncryption: https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#ClientEncryption

## Mocks

The `mocks` folder contains the mocks generated by [mockery](https://github.com/vektra/mockery)
and [gomock](https://github.com/uber-go/mock) tools.

The examples of how to use the mocks can be found in the `examples` folder or check any of the `*_test.go` files as
well.

## Simple Example

### user workflow
1. Create 4 users, with two admins, using `InsertMany` function.
2. Get the admin users only using `Find` function
3. Delete all users using `DeleteMany` function

* [users.go](https://github.com/sv-tools/mongoifc/blob/main/examples/simple/users.go) is a file with a set of functions, like:
  * `Create` to create the users using `InsertMany`
  * `Delete` to delete the users by given IDs
  * `GetAdmins` to return the list of admin users
* [users_test.go](https://github.com/sv-tools/mongoifc/blob/main/examples/simple/users_test.go) is a file with `TestUsersWorkflow` unit tests:
  * `mockery` tests the workflow using `mockery` mocks
  * `gomock` tests the workflow using `gomock` mocks
  * `docker` tests the workflow using real mongo database run by docker

### collection workflow
1. Create a collection with random name.
2. Check that the collection exists.
3. Check that another collection does not exist.
4. Drop collection.
5. Check that the original collection does not exist.

* [collections.go](https://github.com/sv-tools/mongoifc/blob/main/examples/simple/collections.go) is a file with a set of functions, like:
  * `CreateCollection` to create a collection using `CreateCollection`
  * `DropCollection` to delete a collection by given name
  * `CollectionExists` to check that a collection exists
* [collections_test.go](https://github.com/sv-tools/mongoifc/blob/main/examples/simple/collections_test.go) is a file with `TestCollectionsWorkflow` unit tests:
  * `mockery` tests the workflow using `mockery` mocks
  * `gomock` tests the workflow using `gomock` mocks
  * `docker` tests the workflow using real mongo database run by docker
