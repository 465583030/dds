package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/riclava/dds/cluster/config"
	"github.com/riclava/dds/cluster/constants"
	"github.com/riclava/dds/cluster/tasks"
	"github.com/riclava/dds/cluster/utils"

	restful "github.com/emicklei/go-restful"
	"github.com/riclava/dds/api/models"
)

// ListTask retrieve task list
func ListTask(request *restful.Request, response *restful.Response) {
	resp := models.Response{}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = request.Request.RequestURI
	response.WriteHeaderAndEntity(http.StatusOK, resp)
}

// AddTask put a new task to cluster to process
func AddTask(request *restful.Request, response *restful.Response, config *config.Config) {
	var usrHTTPTask tasks.UserHTTPTask
	err := request.ReadEntity(&usrHTTPTask)
	if err != nil {
		handleError(response, err)
		return
	}

	payload, err := json.Marshal(usrHTTPTask)
	if err != nil {
		handleError(response, err)
		return
	}

	endpoint := fmt.Sprintf("%s:%d", config.Host, config.RPCPort)
	r, err := utils.GRPCall(endpoint, config.Host, constants.CidHTTPTaskAdd, string(payload))
	if err != nil {
		log.Println("1")
		handleError(response, err)
		return
	}

	var taskResponse tasks.TaskResponse
	err = json.Unmarshal([]byte(r.Payload), &taskResponse)
	if err != nil {
		log.Println("2")
		handleError(response, err)
		return
	}

	resp := models.Response{}
	resp.Code = taskResponse.Code
	if taskResponse.Code == 0 {
		resp.Message = "success"
	} else {
		resp.Message = "failed"
	}

	resp.Data = taskResponse.Data
	response.WriteHeaderAndEntity(http.StatusOK, resp)
}
