package runner

import (
	"bashCommandCyclicRunner/internal/app/fileutils"
	"encoding/json"
	"fmt"
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