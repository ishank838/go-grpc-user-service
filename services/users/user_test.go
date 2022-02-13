package users_test

import (
	"context"
	"testing"

	"github.com/go-playground/assert/v2"
	user "github.com/ishank838/go-users-grpc/proto/user"
	users "github.com/ishank838/go-users-grpc/services/users"
	"github.com/ishank838/go-users-grpc/services/users/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userStruct struct {
	id      int64
	fname   string
	city    string
	phone   string
	height  float32
	married bool
}

func TestGetUser(t *testing.T) {
	store.InituserData()
	svc := users.UserService{}
	userData := userStruct{id: 1, fname: "Abey", city: "Houxiang", phone: "4824120772", height: 5.7, married: false}

	userId := int64(1)
	u, err := svc.GetUser(context.Background(), &user.UserId{UserId: &userId})

	assert.Equal(t, err, nil)
	assert.Equal(t, u.Id, userData.id)
	assert.Equal(t, u.FirstName, userData.fname)
	assert.Equal(t, u.City, userData.city)
	assert.Equal(t, u.Phone, userData.phone)
	assert.Equal(t, u.Height, userData.height)
	assert.Equal(t, u.Married, userData.married)
}

func TestGetUserEmpty(t *testing.T) {
	store.InituserData()
	svc := users.UserService{}

	u, err := svc.GetUser(context.Background(), nil)

	assert.Equal(t, u, nil)
	assert.Equal(t, err, status.Errorf(codes.InvalidArgument, "userId cann't be empty"))
}

func TestGetUserInvalidUserId(t *testing.T) {
	store.InituserData()
	svc := users.UserService{}

	userId := int64(0)
	u, err := svc.GetUser(context.Background(), &user.UserId{UserId: &userId})

	assert.Equal(t, u, nil)
	assert.Equal(t, err, status.Errorf(codes.InvalidArgument, "invalid userId"))
}

func TestGetAllUsers(t *testing.T) {
	store.InituserData()
	svc := users.UserService{}
	resData := store.UserDb
	list, err := svc.GetAllUsers(context.Background(), nil)
	assert.Equal(t, err, nil)
	assert.Equal(t, len(list.AllUser), len(resData))

	for i, _ := range list.AllUser {
		assert.Equal(t, list.AllUser[i].Id, resData[i].Id)
		assert.Equal(t, list.AllUser[i].FirstName, resData[i].FirstName)
		assert.Equal(t, list.AllUser[i].City, resData[i].City)
		assert.Equal(t, list.AllUser[i].Phone, resData[i].Phone)
		assert.Equal(t, list.AllUser[i].Height, resData[i].Height)
		assert.Equal(t, list.AllUser[i].Married, resData[i].Married)
	}
}
