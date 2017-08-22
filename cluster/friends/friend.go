package friends

import (
	"encoding/json"
	"io/ioutil"
)

// Friend friend of someone
type Friend struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Token    string `json:"token"`
}

// Friends is a set of friend
type Friends struct {
	Friends  map[string]Friend `json:"friends"`
	Location string            `json:"location"`
}

// FriendPair friend pair
type FriendPair struct {
	You    Friend `json:"you"`
	Myself Friend `json:"myself"`
}

//ReadConfig read friends config from disk
func (frands *Friends) ReadConfig() error {
	readBytes, err := ioutil.ReadFile(frands.Location)
	if err != nil {
		return err
	}
	err = json.Unmarshal(readBytes, frands)
	if err != nil {
		return err
	}
	return nil
}

//WriteConfig write mem config to disk
func (frands *Friends) WriteConfig() error {
	payload, err := json.Marshal(frands)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(frands.Location, payload, 0644)
	if err != nil {
		return err
	}
	return nil
}

// AddNewFriend add a new friend to friend list and flush disk
func (frands *Friends) AddNewFriend(friend *Friend) {
	frands.Friends[friend.Username] = (*friend)
	frands.WriteConfig()
}

// DeleteFriend delete a friend from friend list and flush disk
func (frands *Friends) DeleteFriend(friend *Friend) {
	delete(frands.Friends, friend.Username)
	frands.WriteConfig()
}
