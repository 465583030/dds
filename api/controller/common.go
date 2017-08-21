package controller

import (
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful"
)

func handleError(response *restful.Response, err error) {
	log.Print("API caught an error [", err, "]")
	statusCode := http.StatusInternalServerError
	response.AddHeader("Content-Type", "text/plain")
	response.WriteErrorString(statusCode, err.Error()+"\n")
}
