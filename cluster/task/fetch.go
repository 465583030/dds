package task

import (
	"context"
	"log"

	"github.com/riclava/dds/cluster/ddservice"
	"google.golang.org/grpc"
)

func FetchTask(endpoints map[string]string) {
	// Set up a connection to the server.
	clientConn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer clientConn.Close()
	c := ddservice.NewDDServiceClient(clientConn)

	// Contact the server and print out its response.
	r, err := c.Call(context.Background(), &ddservice.DDSRequest{Ip: "172.23.8.99", Time: 12434334, Cid: 1, Payload: "你好!"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Response: %s", r.Payload)
}
