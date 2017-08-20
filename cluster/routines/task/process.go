package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"github.com/riclava/dds/cluster/constants"
	"github.com/riclava/dds/cluster/ddservice"
	"github.com/riclava/dds/cluster/friends"
	"github.com/riclava/dds/cluster/tasks"
	"github.com/riclava/dds/cluster/utils"
)

// ProcessHTTPTask process an http task
func ProcessHTTPTask(myself *friends.Friend) {
	for {
		doProcess(myself)
		time.Sleep(1 * time.Second)
	}
}

func doProcess(myself *friends.Friend) {
	httpTask := <-HTTPTaskProcessChannel

	log.Println("processing task ", httpTask.Filename, httpTask.RangeStart, httpTask.RangeEnd)

	headers := make(map[string]string)
	headers["Range"] = fmt.Sprintf("bytes=%d-%d", httpTask.RangeStart, httpTask.RangeEnd)

	body, err := utils.Get(httpTask.URL, headers)
	if err != nil {
		log.Println(err)
		return
	}

	httpTaskBlock := tasks.HTTPTaskBlock{}
	httpTaskBlock.HTTPTask = httpTask
	httpTaskBlock.Block = body

	payload, err := json.Marshal(httpTaskBlock)
	if err != nil {
		log.Println(err)
		return
	}
	doSubmit(httpTask.Endpoint, myself, string(payload))
}

func doSubmit(endpoint string, myself *friends.Friend, payload string) {
	// Set up a connection to the server.
	clientConn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer clientConn.Close()
	c := ddservice.NewDDServiceClient(clientConn)

	// Contact the server and print out its response.
	r, err := c.Call(context.Background(), &ddservice.DDSRequest{Ip: myself.Host, Time: time.Now().Unix(), Cid: constants.CidHTTPTaskSubmit, Payload: payload})
	if err != nil {
		log.Println("could not fetch a task from friend", err)
		return
	}
	var httpTask tasks.HTTPTask
	err = json.Unmarshal([]byte(r.Payload), &httpTask)
	if err != nil {
		log.Println("can not submit task ", err)
	}
	log.Println("submit a task to " + endpoint + ", filename: " + httpTask.Filename + " success")
}
