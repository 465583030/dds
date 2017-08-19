package friends

// User common user
type User struct {
	Username string `json:"username"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
}

// Users is a set of user
type Users map[string]User
