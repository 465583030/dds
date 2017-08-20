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
	Users   *friends.Users
	Friends *friends.Friends
	Config  *config.Config
}

// Call implement ddservice.CallServer interface
// 1 => Fetch Task
// 2 => Put Task
// 3 => Submit Task
// 3 => Add Friend
// 4 => Delete Friend
func (s *GRPCServer) Call(ctx context.Context, in *ddservice.DDSRequest) (*ddservice.DDSResponse, error) {
	log.Println("Recieved a request from ", in.Ip, ", cid: ", in.Cid)
	var response ddservice.DDSResponse
	switch in.Cid {
	case constants.CidHTTPTaskFetch:
		hdl := handler.TaskHandler{}
		response = hdl.HandleFetch(in)
		break
	case constants.CidHTTPTaskPut:
		hdl := handler.TaskHandler{
			Friends: s.Friends,
			Config:  s.Config,
		}
		response = hdl.HandlePut(in)
		break
	case constants.CidHTTPTaskSubmit:
		hdl := handler.TaskHandler{
			Friends: s.Friends,
			Config:  s.Config,
		}
		response = hdl.HandleSubmit(in)
		break
	case constants.CidFriendAdd:
		hdl := handler.FriendHandler{
			Users:   s.Users,
			Friends: s.Friends,
		}
		response = hdl.HandleAdd(in)
		break
	case constants.CidFriendDelete:
		hdl := handler.FriendHandler{
			Users:   s.Users,
			Friends: s.Friends,
		}
		response = hdl.HandleDelete(in)
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
func Serve(addr string, users *friends.Users, friends *friends.Friends, config *config.Config) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	s := grpc.NewServer()
	ddservice.RegisterDDServiceServer(s, &GRPCServer{
		Users:   users,
		Friends: friends,
		Config:  config,
	})
	// Register reflection service on gRPC server.
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		return err
	}
	return nil
}
