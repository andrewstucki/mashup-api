package worker

import (
  "github.com/photoionized/goworker"
  "github.com/mashup-cms/mashup-api/services/github"
  "github.com/mashup-cms/mashup-api/connections"
	"github.com/mashup-cms/mashup-api/commands"

  "log"
  "flag"
)

var SyncCmd = &commands.Command{
	Name:    "worker:sync",	
	Usage:   "",
	Summary: "Sync stuff",
	Help:    `create extended help here...`,
	Run:     syncRun,
}

func syncRun(cmd *commands.Command, args ...string) {
  connections.SetupPostgres()
  goworker.Register("SynchronizeUser", github.SyncGithubAccount)
  flag.Parse()
  log.Printf("%s", flag.Args())
  worker := &goworker.Worker{
    Queues: "github_sync",
    IntervalFloat: 5.0,
    Concurrency: 25,
    Connections: 2,
    Uri: "redis://localhost:6379/",
    Namespace: "",
    ExitOnComplete: false,
  }
  if err := worker.Work(); err != nil {
    log.Printf("Error: %s", err.Error())
  }
}