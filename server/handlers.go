package server

import (
	"fmt"
	pb "github.com/im-auld/users/proto"
	"golang.org/x/net/context"
	"strconv"
	"time"
)

type UserService struct {
	DB DBClient
}

func (us *UserService) GetUser(ctx context.Context, request *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	userID := strconv.FormatInt(request.UserId, 10)
	user, found := us.DB.GetUser(userID)
	var errorMessage *pb.Error
	if !found {
		user = &pb.User{
			UserId:      request.UserId,
			FirstName:   fmt.Sprintf("User %d", request.UserId),
			LastName:    fmt.Sprintf("LastName%d", request.UserId),
			Email:       fmt.Sprintf("User %d", request.UserId),
			IsActive:    true,
			DateCreated: time.Now().Unix(),
		}
		us.DB.SetUser(user)
	}
	response := &pb.GetUserResponse{User: user, Error: errorMessage}
	return response, nil
}
