package users

import (
	"context"
	"log"

	users "github.com/ishank838/go-users-grpc/proto/user"
	store "github.com/ishank838/go-users-grpc/services/users/store"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	users.UnimplementedUsersServer
}

func (s *UserService) GetUser(ctx context.Context, userId *users.UserId) (*users.User, error) {
	log.Println(userId)
	if userId == nil {
		return nil, status.Error(codes.InvalidArgument, "userId cann't be empty")
	}

	user := store.GetUserById(userId.UserId)
	log.Println(user.Married)
	if user == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid userId")
	}
	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context, void *users.Void) (*users.UserList, error) {
	/*users := store.GetAllUsers()
	log.Println(users)
	list := *[]users.UserList{}
	return nil, nil*/
	return nil, status.Error(codes.InvalidArgument, "api getAllUsers called") //errors.New("api getAllUsers called")
}
