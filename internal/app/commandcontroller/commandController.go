package commandcontroller

import (
	"flag"
	"log"
)

type CommandController struct {
	Action string
	Keys   []struct {
		Key   string
		Value string
	}
}

func ParseCliParams() CommandController {
	var stop bool
	var start bool
	var createJob bool
	var restart bool

	flag.BoolVar(&stop, "stop", false, "set this param to stopping demon")
	flag.BoolVar(&start, "start", false, "set this param from start program")
	flag.BoolVar(&createJob, "create-job", false, "set this param from start program")
	flag.BoolVar(&restart, "restart", false, "set this param from start program")

	flag.Parse()

	var commandController CommandController
	commandController.Action = ""
	if stop {
		commandController.Action = "stop"
	}
	if createJob && "" == commandController.Action {
		commandController.Action = "create"
	}
	if restart && "" == commandController.Action {
		commandController.Action = "restart"
	}
	if (start && commandController.Action == "") || (!start && commandController.Action == "") {
		commandController.Action="start"
	}
	if commandController.Action == "" {
		errorReturn()
	}

	return commandController
}

func errorReturn() {
	log.Fatalf("U must select only one from keys: -stop -start -create-job -restart")
}
