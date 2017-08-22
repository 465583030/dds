package routines

import (
	"fmt"
	"log"

	"github.com/riclava/dds/cluster/config"
	"github.com/riclava/dds/cluster/friends"
	"github.com/riclava/dds/cluster/server"
)

// MainRoutine running a grpc server with params
func MainRoutine(friends *friends.Friends, config *config.Config) {
	if err := server.Serve(fmt.Sprintf("%s:%d", config.Host, config.RPCPort), friends, config); err != nil {
		log.Fatal(err)
	}
}
