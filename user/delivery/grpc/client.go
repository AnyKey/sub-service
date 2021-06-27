package grpc

import (
	"context"
	proto "github.com/AnyKey/sub-service/protocol"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"time"
)

type grpcClient struct {
	subClient proto.SubClient
}

func NewClient() *grpcClient {
	conn, err := grpc.Dial(viper.GetString("serviceMap.sub-service.address"), grpc.WithInsecure())
	if err != nil {
		log.Debugf("[NewClient] Error: %v\n", err)
	}
	// defer conn.Close()

	return &grpcClient{subClient: proto.NewSubClient(conn)}
}

func (nc *grpcClient) GetToken(token string) (*string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	data := &proto.GetTokenRequest{Token: token}
	resp, err := nc.subClient.GetToken(ctx, data)
	if err != nil {
		log.Errorln("[GetToken] Error:", err)
	}

	return &resp.Message, err
}
