package runner

import (
	"bashCommandCyclicRunner/internal/app/fileutils"
	"encoding/json"
	"fmt"
	"log"
)

type JsonEncoder struct {
	Interval int    `json:"interval"`
	Command   string `json:"command"`
}

func GetConfigJson() []JsonEncoder {
	var myJson []JsonEncoder
	jsonResp := fileutils.ReadFile("command.json")
	err:= json.Unmarshal([]byte(jsonResp),&myJson)
	if err!= nil {
		fmt.Println(err)
	}
	return myJson
}

func AddCommand2Config(newConfig []JsonEncoder)  {
	j, err:= json.Marshal(newConfig)
	if err != nil {
		log.Fatalln("Error convertation to json command collection")
	}
	fileutils.RewriteFile(string(j),"command.json")
}