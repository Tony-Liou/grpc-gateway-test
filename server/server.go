package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	servicepb "mypackage.com/gateway/gen/go/your/service/v1"
)

const port = ":9090"

type server struct {
	servicepb.UnimplementedYourServiceServer
}

func (s *server) Echo(ctx context.Context, msg *servicepb.StringMessage) (*servicepb.StringMessage, error) {
	log.Printf("Received: %v\n", msg.Value)
	return msg, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	servicepb.RegisterYourServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
