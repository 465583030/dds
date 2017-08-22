package handler

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
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

// key => unixnano()
// if task is timed out, we will push back to HTTPTaskChannel
var fetchedTasks = make(map[int64]tasks.HTTPTask)

// key => filename | mark when file is finished
// value => block range
var files = make(map[string]map[int64]int64)

// TaskHandler handle all task request from node
type TaskHandler struct {
	Friands *friends.Friends
	Config  *config.Config
}

// HandleFetch handle fetch task request
func (handler *TaskHandler) HandleFetch(in *ddservice.DDSRequest) ddservice.DDSResponse {

	// check if some sended out task is timedout
	for timestamp, task := range fetchedTasks {
		distance := time.Now().UnixNano() - timestamp
		if distance > 300*1000*1000*1000 {
			//expired
			HTTPTaskChannel <- task
			delete(fetchedTasks, timestamp)
		}
	}

	if len(HTTPTaskChannel) == 0 {
		return *(makeDDSResponse(-1, "I have not task to do"))
	}
	// TODO add a lock for channel
	httpTask := <-HTTPTaskChannel
	timeKey, err := strconv.ParseInt(httpTask.Filename[strings.LastIndex(httpTask.Filename, "|")+1:len(httpTask.Filename)], 10, 0)
	if err != nil {
		return *(makeDDSResponse(-1, "can not process task"))
	}
	fetchedTasks[timeKey] = httpTask
	payload, _ := json.Marshal(httpTask)
	return *(makeDDSResponse(0, string(payload)))
}

// HandleAdd handle put task request
func (handler *TaskHandler) HandleAdd(in *ddservice.DDSRequest) ddservice.DDSResponse {
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
	filename = fmt.Sprintf("%s|%d", filename[1:len(filename)], time.Now().UnixNano())
	endpoint := fmt.Sprintf("%s:%d", handler.Config.Host, handler.Config.RPCPort)
	ranges := make(map[int64]int64)
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
		ranges[i] = end
		HTTPTaskChannel <- httpTask
	}
	files[filename] = ranges

	log.Println("Enqueue a new request [", usrTask.URL, "]")
	return *(makeDDSResponse(0, "success"))
}

// HandleSubmit handle submit task request
func (handler *TaskHandler) HandleSubmit(in *ddservice.DDSRequest) ddservice.DDSResponse {
	payload := in.Payload
	var block tasks.HTTPTaskBlock
	err := json.Unmarshal([]byte(payload), &block)
	if err != nil {
		return *(makeDDSResponse(-1, "parse block failed"))
	}
	ranges := files[block.Filename]
	if _, ok := ranges[block.RangeStart]; !ok {
		return *(makeDDSResponse(-2, "task may already submitted"))
	}

	tmpFilename := block.Filename[0:strings.LastIndex(block.Filename, "|")]
	// write block
	if runtime.GOOS == "windows" {
		tmpFilename = fmt.Sprintf("%s\\%s", handler.Config.Directory, tmpFilename)
	} else {
		tmpFilename = fmt.Sprintf("%s/%s", handler.Config.Directory, tmpFilename)
	}

	rawBlock, err := base64.StdEncoding.DecodeString(block.Block)
	if err != nil {
		return *(makeDDSResponse(-3, "decode block failed"))
	}

	if err := utils.WriteFileBlock(tmpFilename, rawBlock, block.RangeStart, block.RangeEnd); err != nil {
		return *(makeDDSResponse(-4, "write block to disk failed"))
	}

	timeKey, err := strconv.ParseInt(block.Filename[strings.LastIndex(block.Filename, "|")+1:len(block.Filename)], 10, 0)
	delete(fetchedTasks, timeKey)

	delete(ranges, block.RangeStart)
	if len(ranges) == 0 {
		// file is done ;-)
		delete(files, block.Filename)
		log.Println("Download file [", block.URL, "] success")
	}

	return *(makeDDSResponse(0, "success"))
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
