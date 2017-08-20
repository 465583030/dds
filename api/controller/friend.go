package controller

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/riclava/dds/api/models"
)

// AddFriend add a new friend
func AddFriend(request *restful.Request, response *restful.Response) {
	resp := models.Response{}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = request.Request.RequestURI
	response.WriteHeaderAndEntity(http.StatusOK, resp)
}

// DeleteFriend delete a friend
func DeleteFriend(request *restful.Request, response *restful.Response) {
	resp := models.Response{}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = request.Request.RequestURI
	response.WriteHeaderAndEntity(http.StatusOK, resp)
}
