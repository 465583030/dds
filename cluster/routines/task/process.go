package task

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/riclava/dds/cluster/constants"
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

	headers := make(map[string]string)
	headers["Range"] = fmt.Sprintf("bytes=%d-%d", httpTask.RangeStart, httpTask.RangeEnd)

	body, err := utils.Get(httpTask.URL, headers)
	if err != nil {
		log.Println(err)
		return
	}

	httpTaskBlock := tasks.HTTPTaskBlock{}
	httpTaskBlock.HTTPTask = httpTask
	httpTaskBlock.Block = base64.StdEncoding.EncodeToString([]byte(body))

	payload, err := json.Marshal(httpTaskBlock)
	if err != nil {
		log.Println(err)
		return
	}
	doSubmit(httpTask.Endpoint, myself, string(payload))
}

func doSubmit(endpoint string, myself *friends.Friend, payload string) {

	resp, err := utils.GRPCall(endpoint, myself.Host, constants.CidHTTPTaskSubmit, payload)
	if err != nil {
		log.Println("error call grpc", err)
		return
	}

	var httpTask tasks.HTTPTask
	err = json.Unmarshal([]byte(resp.Payload), &httpTask)
	if err != nil {
		log.Println("submit task failed ", err)
	}
}
