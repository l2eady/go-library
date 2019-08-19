package main

import (
	"encoding/json"
	"fmt"
	"log"
	"owner/khttp"
)

func main() {
	// khttp.New(url param, body param,header param)
	// url to connect to server
	// body is interface can be anything such as struct, map[string]interface{}
	// header is map[string]string type such as map[string]string{"request_id":"00001"}
	content, err := khttp.New("http://localhost:8031/ping", nil, nil).GET()
	if err != nil { // this error is server error such as no host, or http timeout
		log.Println("server error", err)
	}
	// return json string format, but if err != nil the content will be empty string
	fmt.Println(content)

	// example to parse response to struct, or map[string]interface{}
	var output = map[string]interface{}{}
	json.Unmarshal([]byte(content), &output)
	fmt.Println(output)
}
