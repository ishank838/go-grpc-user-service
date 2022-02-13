package main

import (
	"context"
	"log"

	user "github.com/ishank838/go-users-grpc/proto/user"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		log.Panic(err)
	}

	defer conn.Close()

	userId := int64(4)
	c := user.NewUsersClient(conn)
	response, err := c.GetUser(context.Background(), &user.UserId{UserId: &userId})
	if err != nil {
		log.Println(err)
	}
	log.Println(response)
}
