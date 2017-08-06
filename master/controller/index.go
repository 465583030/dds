package controller

import (
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/riclava/distribute-download-system/master/models"
)

// Index of page
func Index(request *restful.Request, response *restful.Response) {
	resp := models.Response{}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = request.Request.RequestURI
	response.WriteHeaderAndEntity(http.StatusOK, resp)
}
