package task

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/riclava/dds/cluster/constants"
	"github.com/riclava/dds/cluster/ddservice"
	"github.com/riclava/dds/cluster/friends"
	"github.com/riclava/dds/cluster/tasks"
	"google.golang.org/grpc"
)

// HTTPTaskProcessChannel is the global channel for HTTPTasks in processing
var HTTPTaskProcessChannel = make(chan tasks.HTTPTask, 4)

func doFetch(endpoint string, myself *friends.Friend) {
	// Set up a connection to the server.
	clientConn, err := grpc.Dial(endpoint, grpc.WithInsecure())
	if err != nil {
		log.Printf("did not connect: %v", err)
		return
	}
	defer clientConn.Close()
	c := ddservice.NewDDServiceClient(clientConn)

	// Contact the server and print out its response.
	r, err := c.Call(context.Background(), &ddservice.DDSRequest{Ip: myself.Host, Time: time.Now().Unix(), Cid: constants.CidHTTPTaskFetch, Payload: "I want a task to do ;-)"})
	if err != nil {
		log.Println("could not fetch a task from friend", err)
		return
	}

	var taskResponse tasks.TaskResponse

	err = json.Unmarshal([]byte(r.Payload), &taskResponse)
	if err != nil {
		log.Println("can not parse task response", err)
		return
	}
	var httpTask tasks.HTTPTask
	err = json.Unmarshal([]byte(taskResponse.Data), &httpTask)
	if err != nil {
		log.Println("can not parse task", err)
		return
	}
	HTTPTaskProcessChannel <- httpTask
}

// FetchTask fetch task from friends
func FetchTask(myself *friends.Friend, frands *friends.Friends) {
	for {
		for _, friend := range *frands {
			doFetch(fmt.Sprintf("%s:%d", friend.Host, friend.Port), myself)
		}
		time.Sleep(500 * time.Millisecond)
	}
}
