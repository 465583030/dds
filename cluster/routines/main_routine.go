package routines

import (
	"fmt"
	"log"

	"github.com/riclava/dds/cluster/config"
	"github.com/riclava/dds/cluster/server"
)

// MainRoutine running a thrift server with params
func MainRoutine(cfg *config.Config) {
	if err := server.Serve(fmt.Sprintf("%s:%d", cfg.Host, cfg.RPCPort)); err != nil {
		log.Fatal(err)
	}
}
