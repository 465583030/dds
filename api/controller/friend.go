package controller

import (
	"errors"
	"fmt"
	"net/http"

	restful "github.com/emicklei/go-restful"
	"github.com/riclava/dds/api/models"
	"github.com/riclava/dds/cluster/constants"
	"github.com/riclava/dds/cluster/friends"
	"github.com/riclava/dds/cluster/utils"
)

// AddFriend add a new friend
func AddFriend(request *restful.Request, response *restful.Response, frands *friends.Friends) {
	var friend friends.Friend
	err := request.ReadEntity(&friend)
	if err != nil {
		handleError(response, err)
		return
	}

	// if alive
	// 	add to user list as well as friend list
	// else
	// 	return error

	// check if user is alive
	payload := "are you ok?"
	r, err := utils.GRPCall(fmt.Sprintf("%s:%d", friend.Host, friend.Port), friend.Host, constants.CidEcho, payload)
	if err != nil {
		handleError(response, err)
		return
	}
	if r.Payload != payload {
		handleError(response, errors.New("echo test failed"))
		return
	}
	frands.AddNewFriend(&friend)

	resp := models.Response{}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = request.Request.RequestURI
	response.WriteHeaderAndEntity(http.StatusOK, resp)
}

// DeleteFriend delete a friend
func DeleteFriend(request *restful.Request, response *restful.Response, frands *friends.Friends) {

	var friend friends.Friend
	err := request.ReadEntity(&friend)
	if err != nil {
		handleError(response, err)
		return
	}

	frands.DeleteFriend(&friend)

	resp := models.Response{}
	resp.Code = 0
	resp.Message = "success"
	resp.Data = request.Request.RequestURI
	response.WriteHeaderAndEntity(http.StatusOK, resp)
}
