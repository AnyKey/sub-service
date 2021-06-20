package grpc

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	proto "sub-github.com/AnyKey/sub-service.git/protocol"
	"sub-github.com/AnyKey/sub-service.git/user"
)

type grpcServer struct {
	usecase user.Usecase
	proto.SubServer
}

// Launch will create new an server object and register in grpc server
func Launch(s *grpc.Server, uuc user.Usecase) {
	proto.RegisterSubServer(s, &grpcServer{
		usecase: uuc,
	})
}

func (s *grpcServer) BroadcastDone(ctx context.Context, in *proto.GetTokenRequest) (*proto.GetTokenResponse, error) {
	log.Debugf("Received: %v", in.GetMessage())
	return &proto.GetTokenResponse{Token: in.GetMessage() + " Received"}, nil
}

