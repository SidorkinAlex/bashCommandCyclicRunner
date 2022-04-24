package main

import (
	"bashCommandCyclicRunner/internal/app/commandcontroller"
	"bashCommandCyclicRunner/internal/app/fileutils"
	"bashCommandCyclicRunner/internal/app/jobcreator"
	"bashCommandCyclicRunner/internal/app/runner"
	"fmt"
	"github.com/sevlyar/go-daemon"
	"log"
	"os"
	"os/exec"
	"time"
)

func main() {

	comCont :=commandcontroller.ParseCliParams()
	if comCont.Action == "stop" {
		stoppingApp()
	} else if comCont.Action == "start" {
		demonise()
		log.Println("Programm bashCommandCyclicRunner has been success running")
	} else if comCont.Action == "restart" {
		restartApp()
	} else if comCont.Action == "create" {
		createJob()
	} else {
		log.Fatalln("Action not set Error")
	}
}

func createJob() {
	jobcreator.CreateJob()
}

func restartApp() {
	stoppingApp()
	demonise()
	log.Println("App is restarted success")
}
func demonise()  {
	cntxt := &daemon.Context{
		PidFileName: ".sample.pid.log",
		PidFilePerm: 0644,
		LogFileName: ".sample.log",
		LogFilePerm: 0640,
		WorkDir:     "./",
		Umask:       027,
		Args:        []string{"[go-daemon sample]"},
	}

	d, err := cntxt.Reborn()
	if err != nil {
		log.Fatal("Unable to run: ", err)
	}
	if d != nil {
		return
	}
	defer cntxt.Release()

	log.Print("- - - - - - - - - - - - - - -")
	log.Print("daemon started")
	runCyclicCommand()
	runControllerPrograms()
}

func stoppingApp() {
	if !fileutils.HasFile(".sample.pid.log"){
		log.Println("App not running")
	} else {
		fmt.Printf("%s", "stopping app in progress")
		fileutils.WriteFile("", ".stop.log")
		// TODO сделать проверку на запущенный процесс с пидом указанным в файлике старта
		for {
			isStoped := fileutils.ReadFile(".stop.log")
			fmt.Printf("%s", ".")
			if isStoped == "stoped" {
				fmt.Println("")
				log.Println("app is stoped")
				break
			}
			time.Sleep(1 * time.Second)
		}
	}
}

func runCommand(command runner.JsonEncoder)  {
	for  {
		 exec.Command("/bin/sh","-c",command.Command).Output()
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
	log.Println("Programm bashCommandCyclicRunner has been success running")
	os.Remove(".stop.log")
	for {
		if _, err := os.Stat(".stop.log"); err == nil {
			fileutils.WriteFile("stoped",".stop.log")
			fmt.Println("")
			log.Fatalln("servis-stoped")
		}
		time.Sleep(5 * time.Second)
	}
}