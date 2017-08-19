package routines

import (
	"log"

	"github.com/riclava/dds/cluster/config"
	"github.com/riclava/dds/cluster/friends"
)

// TaskRoutine running many daemon thread to fetch and process tasks
func TaskRoutine(cfg *config.Config) {
	var usrs friends.Users

	log.Println(usrs)

}
