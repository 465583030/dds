package task

const (
	HTTP_CID = 1
)

type HTTPTask struct {
	URL        string `json:"url"`
	Endpoint   string `json:"endpoint"`
	RangeStart int64  `json:"range_start"`
	RangeEnd   int64  `json:"range_end"`
	Filename   string `json:"filename"`
}

// HTTPTaskChannel is the global channel for HTTPTasks
var HTTPTaskChannel = make(chan HTTPTask, 512)
