package users

import (
	"context"

	user "github.com/ishank838/go-users-grpc/proto/user"
	store "github.com/ishank838/go-users-grpc/services/users/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	user.UnimplementedUsersServer
}

func (s *UserService) GetUser(ctx context.Context, userId *user.UserId) (*user.User, error) {
	if userId == nil {
		return nil, status.Errorf(codes.InvalidArgument, "userId cann't be empty")
	}

	user := store.GetUserById(*userId.UserId)
	if user == nil {
		return nil, status.Errorf(codes.InvalidArgument, "invalid userId")
	}
	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, void *user.Void) (*user.UserList, error) {
	users := store.GetAllUsers()
	if users == nil {
		return nil, status.Errorf(codes.Internal, "unable to get user list")
	}

	list := user.UserList{}

	for i, _ := range users {
		list.AllUser = append(list.AllUser, &users[i])
	}

	return &list, nil
}
