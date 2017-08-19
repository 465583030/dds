package server

import (
	"net"

	"github.com/riclava/dds/cluster/ddservice"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

// GRPCServer is used to implement ddservice.CallServer.
type GRPCServer struct{}

// Call implement ddservice.CallServer interface
func (s *GRPCServer) Call(ctx context.Context, in *ddservice.DDSRequest) (*ddservice.DDSResponse, error) {
	return &ddservice.DDSResponse{Payload: in.Payload}, nil
}

// Serve start server of grpc
func Serve(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	ddservice.RegisterDDServiceServer(s, &GRPCServer{})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return err
	}
}
