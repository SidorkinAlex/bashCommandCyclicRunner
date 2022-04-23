package jobcreator

import (
	"bashCommandCyclicRunner/internal/app/runner"
	"bufio"
	"fmt"
	"log"
	"os"
)

func CreateJob()  {
	var interval int ;
	var command string;
	log.Println("write interval in sec after they running command")
	fmt.Scanf("%d",&interval)
	log.Println("set command,they must cyclic running")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan() // use `for scanner.Scan()` to keep reading
	command = scanner.Text()
	oldConfig:= runner.GetConfigJson();
	var toAddComand runner.JsonEncoder
	toAddComand.Command = command
	toAddComand.Interval = interval
	newConfig :=append(oldConfig,toAddComand)
	runner.AddCommand2Config(newConfig)
	log.Println("your command added success")
}
