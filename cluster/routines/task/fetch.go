package task

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/riclava/dds/cluster/constants"
	"github.com/riclava/dds/cluster/friends"
	"github.com/riclava/dds/cluster/tasks"
	"github.com/riclava/dds/cluster/utils"
)

// HTTPTaskProcessChannel is the global channel for HTTPTasks in processing
var HTTPTaskProcessChannel = make(chan tasks.HTTPTask, 4)

func doFetch(endpoint string, myself *friends.Friend) {

	r, err := utils.GRPCall(endpoint, myself.Host, constants.CidHTTPTaskFetch, "I want a task to do ;-)")
	if err != nil {
		// maybe friend is not up or removed
		return
	}

	var taskResponse tasks.TaskResponse
	err = json.Unmarshal([]byte(r.Payload), &taskResponse)
	if err != nil {
		log.Println("can not parse task response", err)
		return
	}
	if taskResponse.Code != 0 {
		if taskResponse.Code == -2 {
			// there is not task to process
		} else {
			log.Println("task is invalid")
		}
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
		for _, friend := range frands.Friends {
			go doFetch(fmt.Sprintf("%s:%d", friend.Host, friend.Port), myself)
			time.Sleep(500 * time.Millisecond)
		}
		time.Sleep(500 * time.Millisecond)
	}

}
