package grpccon

import (
	"errors"
	"log"
	"net"

	"github.com/itzaddddd/game-shop/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	authPb "github.com/itzaddddd/game-shop/service/auth/proto"
	itemPb "github.com/itzaddddd/game-shop/service/item/proto"
	playerPb "github.com/itzaddddd/game-shop/service/player/proto"
)

type (
	GrpcClientFactoryHandler interface {
		Auth() authPb.AuthGrpcServiceClient
		Item() itemPb.ItemGrpcServiceClient
		Player() playerPb.PlayerGrpcServiceClient
	}

	grpcClientFactory struct {
		client *grpc.ClientConn
	}
)

func (c *grpcClientFactory) Auth() authPb.AuthGrpcServiceClient {
	return authPb.NewAuthGrpcServiceClient(c.client)
}

func (c *grpcClientFactory) Item() itemPb.ItemGrpcServiceClient {
	return itemPb.NewItemGrpcServiceClient(c.client)
}

func (c *grpcClientFactory) Player() playerPb.PlayerGrpcServiceClient {
	return playerPb.NewPlayerGrpcServiceClient(c.client)
}

func NewGrpcClient(host string) (GrpcClientFactoryHandler, error) {
	opts := make([]grpc.DialOption, 0)

	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.NewClient(host, opts...)

	if err != nil {
		log.Printf("Error: Grpc client connection failed: %s", err.Error())
		return nil, errors.New("error: grpc client connection failed")
	}

	return &grpcClientFactory{
		client: conn,
	}, nil
}

func NewGrpcServer(jwtCfg *config.Jwt, host string) (*grpc.Server, net.Listener) {
	opts := make([]grpc.ServerOption, 0)

	grpcServer := grpc.NewServer(opts...)

	lis, err := net.Listen("tcp", host)

	if err != nil {
		log.Fatalf("Error: failed to listen: %v", err)
	}

	return grpcServer, lis
}
