package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/riclava/dds/cluster/config"
	"github.com/riclava/dds/cluster/ddservice"
	"github.com/riclava/dds/cluster/friends"
	"github.com/riclava/dds/cluster/tasks"
	"github.com/riclava/dds/cluster/utils"
)

// HTTPTaskChannel is the global channel for HTTPTasks
var HTTPTaskChannel = make(chan tasks.HTTPTask, 10240)

// key => unixnano() | filename | timeout sec
// if task is timed out, we will push back to HTTPTaskChannel
var fetchedTasks = make(map[string]tasks.HTTPTask)

// key => unixnano() | filename | block number
// value => blocks
var files = make(map[string]map[string]tasks.HTTPTaskBlock)

// TaskHandler handle all task request from node
type TaskHandler struct {
	Friends *friends.Friends
	Config  *config.Config
}

// HandleFetch handle fetch task request
func (handler *TaskHandler) HandleFetch(in *ddservice.DDSRequest) ddservice.DDSResponse {

	response := ddservice.DDSResponse{Payload: "fetch task success"}

	return response
}

// HandlePut handle put task request
func (handler *TaskHandler) HandlePut(in *ddservice.DDSRequest) ddservice.DDSResponse {
	payload := in.Payload
	var usrTask tasks.UserHTTPTask
	err := json.Unmarshal([]byte(payload), &usrTask)
	if err != nil {
		return *(makeDDSResponse(-1, "parse user task failed"))
	}

	// split
	// head file information
	headers := make(map[string]string)
	respHeaders, err := utils.Head(usrTask.URL, headers)
	if err != nil {
		return *(makeDDSResponse(-1, "url point resource is not valid"))
	}
	contentLength := respHeaders.Get("Content-Length")
	iContentLength, err := strconv.ParseInt(contentLength, 10, 0)
	if err != nil {
		return *(makeDDSResponse(-1, "resource server content length is not valid"))
	}
	var blockSize int64 = 512 * 1024 // 512 KB
	var i int64
	filename := usrTask.URL[strings.LastIndex(usrTask.URL, "/"):len(usrTask.URL)]
	filename = filename[1:len(filename)]
	filename = fmt.Sprintf("%d%s%s", time.Now().UnixNano(), "|", filename)
	endpoint := fmt.Sprintf("%s:%d", handler.Config.Host, handler.Config.RPCPort)
	for i = 0; i < iContentLength; i += blockSize {
		end := i + blockSize - 1
		if end > iContentLength {
			end = iContentLength
		}
		httpTask := tasks.HTTPTask{
			URL:        usrTask.URL,
			Endpoint:   endpoint,
			RangeStart: i,
			RangeEnd:   end,
			Filename:   filename,
		}
		log.Println(httpTask)
		HTTPTaskChannel <- httpTask
	}

	return *(makeDDSResponse(0, "success"))
}

// HandleSubmit handle submit task request
func (handler *TaskHandler) HandleSubmit(in *ddservice.DDSRequest) ddservice.DDSResponse {
	response := ddservice.DDSResponse{Payload: "submit task success"}

	return response
}

func makeDDSResponse(code int, data string) *ddservice.DDSResponse {
	resp := &tasks.TaskResponse{
		Code: code,
		Data: data,
	}
	respPayload, _ := json.Marshal(resp)
	response := ddservice.DDSResponse{Payload: string(respPayload)}
	return &response
}
