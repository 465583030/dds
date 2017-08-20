package friends

// Friend friend of someone
type Friend struct {
	User
	Token string `json:"token"`
}

// Friends is a set of friend
type Friends map[string]Friend

// FriendPair friend pair
type FriendPair struct {
	You    Friend `json:"you"`
	Myself Friend `json:"myself"`
}
