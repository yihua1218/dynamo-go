package dynamo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func (db DB) save(res interface{}, fileName string) {
	b, err := json.Marshal(res)
	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}
	err = ioutil.WriteFile(fileName, b, 0644)
	if err != nil {
		panic(err)
	}
}
