## KHTTP Caller Library ##

Package KHTTP Caller provide the function that wrap the http caller to make it easy to use, and handle
support methods `GET,POST,PUT,PATCH,DELETE`

## How to use GET method in KHTTP Caller ##

``` go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"khttp"
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

```

## How to use POST method in KHTTP Caller ##

``` go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"khttp"
)

func main() {
	// khttp.New(url param, body param,header param)
	// url to connect to server
	// body is interface can be anything such as struct, map[string]interface{}
	// header is map[string]string type such as map[string]string{"request_id":"00001"}

	var body = map[string]interface{}{"first_name": "John Ma"}
	content, err := khttp.New("http://localhost:8031/ping", body, nil).POST()
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

```

## How to use PUT method in KHTTP Caller ##

``` go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"khttp"
)

func main() {
	// khttp.New(url param, body param,header param)
	// url to connect to server
	// body is interface can be anything such as struct, map[string]interface{}
	// header is map[string]string type such as map[string]string{"request_id":"00001"}

	var body = map[string]interface{}{"first_name": "John Ma"}
	content, err := khttp.New("http://localhost:8031/ping", body, nil).PUT()
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

```

## How to use PATCH method in KHTTP Caller ##

``` go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"khttp"
)

func main() {
	// khttp.New(url param, body param,header param)
	// url to connect to server
	// body is interface can be anything such as struct, map[string]interface{}
	// header is map[string]string type such as map[string]string{"request_id":"00001"}

	var body = map[string]interface{}{"first_name": "John Ma"}
	content, err := khttp.New("http://localhost:8031/ping", body, nil).PATCH()
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
```


## How to use DELETE method in KHTTP Caller ##

``` go
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"khttp"
)

func main() {
	// khttp.New(url param, body param,header param)
	// url to connect to server
	// body is interface can be anything such as struct, map[string]interface{}
	// header is map[string]string type such as map[string]string{"request_id":"00001"}

	content, err := khttp.New("http://localhost:8031/ping", nil, nil).DELETE()
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
```

## Setup HTTP Client ##

when you want to custom your http client, you also can setup by your self

```go
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"khttp"
)

func main() {
	//initlize new http client
	client := &http.Client{}

	// initlize LHTTP caller
	caller := khttp.New("http://localhost:8031/ping", nil, nil)
	// assign new client to LHTTP Caller
	caller.SetClient(client)

	/** another way
		content,err := khttp.New("http://localhost:8031/ping", nil, nil).SetClient(client).GET()
	**/
	content, err := caller.GET()
	if err != nil {
		fmt.Println("server error", err)
	}
	var m = map[string]interface{}{}
	json.Unmarshal([]byte(content), &m)
	fmt.Println(content)
	fmt.Println(m)

}

```