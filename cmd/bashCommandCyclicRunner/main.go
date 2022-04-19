package main

import (
	"bashCommandCyclicRunner/internal/app/fileutils"
	"bashCommandCyclicRunner/internal/app/runner"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	if(checkStopApp()){
		stoppingApp()
	} else {
		runCyclicCommand()
		runControllerPrograms()
	}
}

func stoppingApp() {
	fmt.Printf("%s","stopping app in progress")
	fileutils.WriteFile("",".stop.log")
	for {
		isStoped := fileutils.ReadFile(".stop.log")
		if isStoped == "stoped" {
			fmt.Println("app is stoped")
			log.Fatalf("app is stoped")
		}
		fmt.Printf("%s",".")
		time.Sleep(1 * time.Second)
	}
}

func checkStopApp() bool {
	var stop bool

	flag.BoolVar(&stop, "stop", false , "set this param to stopping demon")
	flag.Parse()

	fmt.Println(stop)
	return stop
}

func runCommand(command runner.JsonEncoder)  {
	for  {
		out, err := exec.Command("/bin/sh","-c",command.Command).Output()
		if err != nil {
			fmt.Println(err)
		}
		//fileutils.WriteFile(command.Command + string(out)+ "\n","start_log.log")
		fmt.Println(out)
		time.Sleep(time.Duration(command.Interval) * time.Second)
	}
}

func runCyclicCommand()  {
	var CommandCollection []runner.JsonEncoder
	CommandCollection = runner.GetConfigJson()
	for index, command := range CommandCollection{
		fmt.Println("Command ",index, command.Command)
		fmt.Println("Interval ",index, command.Interval)
		go runCommand(command)
	}
}
func runControllerPrograms(){
	os.Remove(".stop.log")
	for {
		if _, err := os.Stat(".stop.log"); err == nil {
			fileutils.WriteFile("stoped",".stop.log")
			log.Fatalln("servis-stoped")
		}
		time.Sleep(5 * time.Second)
	}
}