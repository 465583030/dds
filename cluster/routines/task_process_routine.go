package routines

/*
 * task routine
 * this routine do three things
 * 1. fetch task from friends and put into channel
 * 2. multi routine will process these tasks
 * 3. submit task to owner
 */

import (
	"github.com/riclava/dds/cluster/config"
	"github.com/riclava/dds/cluster/friends"
	"github.com/riclava/dds/cluster/routines/task"
)

// TaskProcessRoutine running many daemon thread to fetch and process tasks
func TaskProcessRoutine(cfg *config.Config, myself *friends.Friend, frands *friends.Friends) {
	go task.FetchTask(myself, frands)
	for i := 0; i < 4; i++ { // 4 routine to process
		go task.ProcessHTTPTask(myself)
	}
}
