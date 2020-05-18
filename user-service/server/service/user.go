package service

import (
	"github.com/hydaizd/zdmicroservice/user-service/pb"
	"golang.org/x/net/context"
)

// Server is used to implement helloworld.GreeterServer.
type Server struct{}

// Register implements pb.GreeterServer
func (s *Server) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterReply, error) {
	return &pb.RegisterReply{Message: "Hello " + in.Username}, nil
}
