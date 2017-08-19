package handler

import (
	"fmt"
	"log"
	"net/http"
	"time"

	restful "github.com/emicklei/go-restful"
	"github.com/riclava/dds/api/controller"
	"github.com/riclava/dds/api/models"
	"github.com/riclava/dds/cluster/config"
)

const (
	// RequestLogString is a template for request log message.
	RequestLogString = "[%s] Incoming %s %s %s request from: %s"

	// ResponseLogString is a template for response log message.
	ResponseLogString = "[%s] Outcoming response to %s with %d status code"
)

// APIHandler is a representation of API handler
type APIHandler struct {
	token string
}

// CreateAPIHandler create an API handler for restful API
func CreateAPIHandler(cfg *config.Config) (http.Handler, error) {

	container := restful.NewContainer()
	apiHandler := APIHandler{
		token: cfg.Token,
	}

	webService := new(restful.WebService)
	webService.Filter(logRequestAndResponse)

	webService.Path("/api/v1").
		Consumes(restful.MIME_JSON).
		Produces(restful.MIME_JSON)
	container.Add(webService)

	webService.Route(
		webService.GET("/").To(apiHandler.handlerGetRoot).Writes(models.Response{}))

	return container, nil
}

// logRequestAndReponse is a web-service filter function used for request and response logging.
func logRequestAndResponse(request *restful.Request, response *restful.Response, chain *restful.FilterChain) {
	log.Printf(formatRequestLog(request))
	chain.ProcessFilter(request, response)
	log.Printf(formatResponseLog(response, request))
}

// formatRequestLog formats request log string.
func formatRequestLog(request *restful.Request) string {
	uri := ""
	if request.Request.URL != nil {
		uri = request.Request.URL.RequestURI()
	}

	return fmt.Sprintf(RequestLogString, time.Now().Format(time.RFC3339), request.Request.Proto,
		request.Request.Method, uri, request.Request.RemoteAddr)
}

// formatResponseLog formats response log string.
func formatResponseLog(response *restful.Response, request *restful.Request) string {
	return fmt.Sprintf(ResponseLogString, time.Now().Format(time.RFC3339),
		request.Request.RemoteAddr, response.StatusCode())
}

func (apiHandler *APIHandler) handlerGetRoot(request *restful.Request, response *restful.Response) {
	controller.Index(request, response)
}
