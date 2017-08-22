package handler

import "github.com/riclava/dds/cluster/ddservice"

// EchoHandler used for common check
type EchoHandler struct {
}

// HandleEcho handle echo request
func (handler *EchoHandler) HandleEcho(in *ddservice.DDSRequest) ddservice.DDSResponse {
	return ddservice.DDSResponse{Payload: in.Payload}
}
