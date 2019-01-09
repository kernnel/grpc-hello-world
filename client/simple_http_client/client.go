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
	Address = "127.0.0.1:5058"
)

func main() {
	tlsClient := gtls.Client{
		ServerName: "grpc-hello-world",
		CertFile:   "../../conf/server/server.pem",
	}
	c, err := tlsClient.GetTLSCredentials()
	if err != nil {
		log.Fatalf("tlsClient.GetTLSCredentials err: %v", err)
	}

	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}
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
