package tasks

// UserHTTPTask frontend user put task
type UserHTTPTask struct {
	URL string `json:"url"`
}

// HTTPTask http type of task
type HTTPTask struct {
	URL        string `json:"url"`
	Endpoint   string `json:"endpoint"`
	RangeStart int64  `json:"range_start"`
	RangeEnd   int64  `json:"range_end"`
	Filename   string `json:"filename"`
}

// HTTPTaskBlock http task block
type HTTPTaskBlock struct {
	HTTPTask
	Block string `json:"block"`
}

// TaskResponse task request exec response
type TaskResponse struct {
	Code int    `json:"code"`
	Data string `json:"data"`
}

// HTTPTasks http task list
type HTTPTasks []HTTPTask
