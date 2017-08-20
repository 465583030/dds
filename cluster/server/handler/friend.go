package handler

import (
	"github.com/riclava/dds/cluster/ddservice"
	"github.com/riclava/dds/cluster/friends"
)

// FriendHandler handle all task request from node
type FriendHandler struct {
	Users   *friends.Users
	Friends *friends.Friends
}

// HandleAdd handle add friend request
func (handler *FriendHandler) HandleAdd(in *ddservice.DDSRequest) ddservice.DDSResponse {
	response := ddservice.DDSResponse{Payload: "add friend success"}

	return response
}

// HandleDelete handle delete friend request
func (handler *FriendHandler) HandleDelete(in *ddservice.DDSRequest) ddservice.DDSResponse {
	response := ddservice.DDSResponse{Payload: "delete friend success"}

	return response
}
