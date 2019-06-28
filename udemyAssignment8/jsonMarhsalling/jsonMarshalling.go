package main

import (
	"encoding/json"
	"fmt"
	"time"
)


//Struct for Deserialization
//Contains fields from Json and tags to
//map json fields to the differing struct fields
type JsonPayload struct {
	Id string `json:"id"`
	DetailType string `json:"detail-type"`
	Source string `json:"source"`
	Account string `json:"account"`
	Time time.Time `json:"time"`
	Region string `json:"region"`
	Resources []string `json:"resources"`
}

func main() {

	//Json Marshalling
	rawByteSlice := []byte(`{
  "id": "cdc73f9d-aea9-11e3-9d5a-835b769c0d9c",
  "detail-type": "Scheduled Event",
  "source": "aws.events",
  "account": "00000000000000",
  "time": "1970-01-01T00:00:00Z",
  "region": "us-east-1",
  "resources": [
    "arn:aws:events:us-east-1:123456789012:rule/ExampleRule"
  ]}`)

	jsonArray, err := json.Marshal(rawByteSlice)
	if err != nil {
		fmt.Println("Something went wrong during marshalling: ", err)
	}

	//Dump jsonArray
	fmt.Println("Here is the jsonArray in bytes: ", jsonArray)

	//Json Unmarshalling
	var jsonPayload JsonPayload
	err = json.Unmarshal(rawByteSlice,&jsonPayload)
		if err != nil {
			fmt.Println("Something went wrong during the unmarshalling process: ", err)
		}

	//Dump the jsonPayload object
	fmt.Println("Here is the jsonPayload and a few values.")
	fmt.Println("Payload: ", jsonPayload)
	fmt.Println("Id:", jsonPayload.Id)
	fmt.Println("Detail-Type: ", jsonPayload.DetailType)
	
}
