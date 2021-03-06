package handlers

import (
	"context"

	pb "github.com/Reasno/jumpstart"
)

// NewService returns a naïve, stateless implementation of Service.
func NewService() pb.AppServer {
	return appService{}
}

type appService struct{}

func (s appService) Create(ctx context.Context, in *pb.UserRequest) (*pb.GenericReply, error) {
	var resp pb.GenericReply
	return &resp, nil
}
