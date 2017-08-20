package utils

import (
	"context"
	"time"

	"github.com/riclava/dds/cluster/ddservice"
	"google.golang.org/grpc"
)

// GRPCall call a grpc server
func GRPCall(endpoint string, host string, cid int32, payload string) (*ddservice.DDSResponse, error) {
	// Set up a connection to the server.
	clientConn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer clientConn.Close()
	c := ddservice.NewDDServiceClient(clientConn)

	// Contact the server and print out its response.
	r, err := c.Call(context.Background(), &ddservice.DDSRequest{Ip: host, Time: time.Now().Unix(), Cid: cid, Payload: payload})
	if err != nil {
		return nil, err
	}
	return r, nil
}
