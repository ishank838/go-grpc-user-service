package store_test

import (
	"testing"

	"github.com/go-playground/assert/v2"
	"github.com/ishank838/go-users-grpc/services/users/store"
)

func TestGetUserById(t *testing.T) {
	store.InituserData()

	u := store.GetUserById(1)
	res := int64(1)
	assert.Equal(t, *u.Id, res)
}

func TestGetUserByIdZero(t *testing.T) {
	store.InituserData()
	u := store.GetUserById(0)
	assert.Equal(t, u, nil)
}

func TestGetAllUsers(t *testing.T) {
	store.InituserData()

	userList := store.GetAllUsers()

	resList := store.UserDb

	assert.Equal(t, userList, resList)
}
