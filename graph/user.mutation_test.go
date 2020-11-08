// +build integration

package graph_test

import (
	"context"
	"database/sql"
	"log"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/alextanhongpin/go-graphql-template/entity"
	"github.com/alextanhongpin/go-graphql-template/graph/usergraph"
	"github.com/alextanhongpin/go-graphql-template/model"
	"github.com/alextanhongpin/go-graphql-template/pkg/database"
	"github.com/alextanhongpin/go-graphql-template/resolver"
)

func TestCreateUser(t *testing.T) {
	Convey("feature: create user", t, func() {
		db, err := database.NewTestDB()
		if err != nil {
			log.Fatal(err)
		}
		Reset(func() {
			db.Close()
		})

		mutator := createUserMutator(db)
		Convey("scenario: user created", func() {
			user, err := createUser(context.TODO(), mutator)
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)
			So(user.Name(), ShouldEqual, "John Doe")
			So(user.Email(), ShouldEqual, "john.doe@mail.com")
		})

		Convey("scenario: input is empty", func() {
			user, err := mutator.CreateUser(context.TODO(), usergraph.CreateUserArgs{})
			So(err, ShouldNotBeNil)
			So(user, ShouldBeNil)
		})
	})
}

func TestUpdateUser(t *testing.T) {
	Convey("feature: update user", t, func() {
		db, err := database.NewTestDB()
		if err != nil {
			log.Fatal(err)
		}
		Reset(func() {
			db.Close()
		})
		mutator := createUserMutator(db)

		Convey("scenario: user updated", func() {
			user, err := createUser(context.TODO(), mutator)
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)

			updatedUser, err := mutator.UpdateUser(context.TODO(), usergraph.UpdateUserArgs{
				Input: usergraph.UpdateUserInput{
					ID:   user.ID(),
					Name: "John Doe (edited)",
				},
			})
			So(err, ShouldBeNil)
			So(updatedUser.Name(), ShouldEqual, "John Doe (edited)")
		})

		Convey("scenario: input invalid ", func() {
			user, err := createUser(context.TODO(), mutator)
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)

			updatedUser, err := mutator.UpdateUser(context.TODO(), usergraph.UpdateUserArgs{
				Input: usergraph.UpdateUserInput{
					ID:   "xyz",
					Name: "John Doe (edited)",
				},
			})
			So(err, ShouldNotBeNil)
			So(updatedUser, ShouldBeNil)
		})

		Convey("scenario: input is empty", func() {
			user, err := createUser(context.TODO(), mutator)
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)

			updatedUser, err := mutator.UpdateUser(context.TODO(), usergraph.UpdateUserArgs{})
			So(err, ShouldNotBeNil)
			So(updatedUser, ShouldBeNil)
		})
	})
}

func TestDeleteUser(t *testing.T) {
	Convey("feature: delete user", t, func() {
		db, err := database.NewTestDB()
		if err != nil {
			log.Fatal(err)
		}
		Reset(func() {
			db.Close()
		})
		mutator := createUserMutator(db)

		Convey("scenario: user deleted", func() {
			user, err := createUser(context.TODO(), mutator)
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)

			deletedUser, err := mutator.DeleteUser(context.TODO(), usergraph.DeleteUserArgs{
				Input: usergraph.DeleteUserInput{
					ID: user.ID(),
				},
			})
			So(err, ShouldBeNil)
			So(deletedUser, ShouldNotBeNil)
		})

		Convey("scenario: input is invalid", func() {
			user, err := createUser(context.TODO(), mutator)
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)

			deletedUser, err := mutator.DeleteUser(context.TODO(), usergraph.DeleteUserArgs{
				Input: usergraph.DeleteUserInput{
					ID: "xyz",
				},
			})
			So(err, ShouldNotBeNil)
			So(deletedUser, ShouldBeNil)
		})

		Convey("scenario: input is empty", func() {
			user, err := createUser(context.TODO(), mutator)
			So(err, ShouldBeNil)
			So(user, ShouldNotBeNil)

			deletedUser, err := mutator.DeleteUser(context.TODO(), usergraph.DeleteUserArgs{})
			So(err, ShouldNotBeNil)
			So(deletedUser, ShouldBeNil)
		})
	})
}

func createUserMutator(db *sql.DB) *usergraph.Mutation {
	repo := entity.New(db)
	return usergraph.NewMutation(model.NewResolverContext(repo))
}

func createUser(ctx context.Context, mutator *usergraph.Mutation) (*resolver.UserResolver, error) {
	user, err := mutator.CreateUser(
		ctx,
		usergraph.CreateUserArgs{
			Input: usergraph.CreateUserInput{
				Name:  "John Doe",
				Email: "john.doe@mail.com",
			},
		},
	)
	return user, err
}
