package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-hello-world/pkg/gtls"
	pb "grpc-hello-world/proto"
	"log"
)

const (
	// Address gRPC服务地址
	Address = "127.0.0.1:50059"
)

type Auth struct {
	AppKey    string
	AppSecret string
}

func (a *Auth) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{"app_key": a.AppKey, "app_secret": a.AppSecret}, nil
}

func (a *Auth) RequireTransportSecurity() bool {
	return true
}

func main() {
	tlsClient := gtls.Client{
		ServerName: "grpc-hello-world",
		CertFile:   "../../conf/server/server.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}

	auth := Auth{
		AppKey:    "ftd",
		AppSecret: "20181005",
	}
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(c), grpc.WithPerRPCCredentials(&auth))

	defer conn.Close()

	client := pb.NewSearchServiceClient(conn)
	resp, err := client.Search(context.Background(), &pb.SearchRequest{
		Request: "gRPC",
	})
	if err != nil {
		log.Fatalf("client.Search err: %v", err)
	}

	log.Printf("resp: %s", resp.GetResponse())
}
