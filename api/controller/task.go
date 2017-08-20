package controller

import (
	"net/http"

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

// PutTask put a new task to cluster to process
func PutTask(request *restful.Request, response *restful.Response) {
	resp := models.Response{}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = request.Request.RequestURI
	response.WriteHeaderAndEntity(http.StatusOK, resp)
}
