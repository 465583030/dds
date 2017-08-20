package handler

import (
	"encoding/json"

	"github.com/riclava/dds/cluster/ddservice"
	"github.com/riclava/dds/cluster/friends"
)

// FriendHandler handle all task request from node
type FriendHandler struct {
	Myself  *friends.Friend
	Users   *friends.Users
	Friends *friends.Friends
}

// HandleAdd handle add friend request
func (handler *FriendHandler) HandleAdd(in *ddservice.DDSRequest) ddservice.DDSResponse {
	// add user to my friends list, provide his/her's username & token and my username & token
	payload := in.Payload
	var friendPair friends.FriendPair
	err := json.Unmarshal([]byte(payload), &friendPair)
	if err != nil {
		return *(makeDDSResponse(-1, "can not parse friend pair"))
	}
	if !(friendPair.Myself.Username == handler.Myself.Username && friendPair.Myself.Token == handler.Myself.Token) {
		return *(makeDDSResponse(-2, "you've not provide valid information"))
	}
	(*handler.Friends)[friendPair.You.Username] = friendPair.You
	// TODO flush to disk

	return *(makeDDSResponse(0, "success"))
}

// HandleDelete handle delete friend request
func (handler *FriendHandler) HandleDelete(in *ddservice.DDSRequest) ddservice.DDSResponse {
	// delete a friend from my friend list
	payload := in.Payload
	var friendPair friends.FriendPair
	err := json.Unmarshal([]byte(payload), &friendPair)
	if err != nil {
		return *(makeDDSResponse(-1, "can not parse friend pair"))
	}
	if !(friendPair.Myself.Username == handler.Myself.Username && friendPair.Myself.Token == handler.Myself.Token) {
		return *(makeDDSResponse(-2, "you've not provide valid information"))
	}
	delete(*handler.Friends, friendPair.You.Username)
	// TODO flush to disk

	return *(makeDDSResponse(0, "success"))
}
