package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"grpc-hello-world/pkg/gtls"
	pb "grpc-hello-world/proto"
	"log"
	"net"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50059"
)

type SearchService struct {
	auth *Auth
}

func (s *SearchService) Search(ctx context.Context, r *pb.SearchRequest) (*pb.SearchResponse, error) {
	if err := s.auth.Check(ctx); err != nil {
		return nil, err
	}
	return &pb.SearchResponse{Response: r.GetRequest() + " Token Server"}, nil
}

type Auth struct {
	appKey    string
	appSecret string
}

func (a *Auth) Check(ctx context.Context) error {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return status.Errorf(codes.Unauthenticated, "自定义认证 Token 失败")
	}

	var (
		appKey    string
		appSecret string
	)
	if value, ok := md["app_key"]; ok {
		appKey = value[0]
	}
	if value, ok := md["app_secret"]; ok {
		appSecret = value[0]
	}

	if appKey != a.GetAppKey() || appSecret != a.GetAppSecret() {
		return status.Errorf(codes.Unauthenticated, "自定义认证 Token 无效")
	}

	return nil
}

func (a *Auth) GetAppKey() string {
	return "ftd"
}

func (a *Auth) GetAppSecret() string {
	return "20181005"
}
func main() {
	tlsServer := gtls.Server{
		CertFile: "../../conf/server/server.pem",
		KeyFile:  "../../conf/server/server.key",
	}
	c, err := tlsServer.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsServer.GetTLSCredentials err: %v", err)
	}

	server := grpc.NewServer(grpc.Creds(c))
	pb.RegisterSearchServiceServer(server, &SearchService{})

	lis, err := net.Listen("tcp", Address)
	if err != nil {
		log.Fatalf("net.Listen err: %v", err)
	}

	println("Listen on " + Address)

	server.Serve(lis)
}
