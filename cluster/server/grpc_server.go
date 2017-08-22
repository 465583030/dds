package server

import (
	"log"
	"net"

	"github.com/riclava/dds/cluster/config"
	"github.com/riclava/dds/cluster/constants"
	"github.com/riclava/dds/cluster/ddservice"
	"github.com/riclava/dds/cluster/friends"
	"github.com/riclava/dds/cluster/server/handler"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// GRPCServer is used to implement ddservice.CallServer.
type GRPCServer struct {
	Friands *friends.Friends
	Config  *config.Config
}

// Call implement ddservice.CallServer interface
func (s *GRPCServer) Call(ctx context.Context, in *ddservice.DDSRequest) (*ddservice.DDSResponse, error) {
	var response ddservice.DDSResponse
	switch in.Cid {
	case constants.CidHTTPTaskFetch:
		hdl := handler.TaskHandler{}
		response = hdl.HandleFetch(in)
		break
	case constants.CidHTTPTaskAdd:
		hdl := handler.TaskHandler{
			Friands: s.Friands,
			Config:  s.Config,
		}
		response = hdl.HandleAdd(in)
		break
	case constants.CidHTTPTaskSubmit:
		hdl := handler.TaskHandler{
			Friands: s.Friands,
			Config:  s.Config,
		}
		response = hdl.HandleSubmit(in)
		break
	case constants.CidFriendAdd:
		hdl := handler.FriendHandler{
			Friands: s.Friands,
		}
		response = hdl.HandleAdd(in)
		break
	case constants.CidFriendDelete:
		hdl := handler.FriendHandler{
			Friands: s.Friands,
		}
		response = hdl.HandleDelete(in)
		break
	case constants.CidEcho:
		hdl := handler.EchoHandler{}
		response = hdl.HandleEcho(in)
		break
	default:
		msg := "data exchange id (cid) invalid"
		log.Println(msg)
		response = ddservice.DDSResponse{Payload: msg}
		break
	}
	return &response, nil
}

// Serve start server of grpc
func Serve(addr string, friands *friends.Friends, config *config.Config) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	ddservice.RegisterDDServiceServer(s, &GRPCServer{
		Friands: friands,
		Config:  config,
	})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
