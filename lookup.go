package main

import "C"
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type JsonResponse struct {
	Data IPDetails
}

type IPDetails struct {
	Ip           string
	Type         string
	Domain       string
	Organization string
	Route        string
	Valid_host   bool

	//"ip": "103.36.149.7",
	//"type": "IPv4",
	//"domain": "efi.com",
	//"organization": "Electronics For Imaging India Private Limited",
	//"route": "103.36.149.0/24",
	// "valid_host": true}

}

// //export CallIPLookup
// func CallIPLookup(IPAddress *C.char) *C.char {
// 	// this is an exportable funciton
// 	/*
// 		this is a comment
// 	*/
// 	fmt.Println(C.GoString(IPAddress))
// 	return IPAddress

// }

//export IsValidEnvironment
func IsValidEnvironment() C.int {

	url := "https://quiet-garden-38373.herokuapp.com/validate"
	method := "GET"
	client := &http.Client{}
	payload := strings.NewReader("")
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpYXQiOjE1ODE0ODI5MjksIm5iZiI6MTU4MTQ4MjkyOSwianRpIjoiNTU4ZGVkMTQtMjQxMi00NjU4LWE4YjYtNzRlM2I1ZTI1MjVhIiwiaWRlbnRpdHkiOiJ5b2RlYnVAZ21haWwuY29tMTIiLCJmcmVzaCI6ZmFsc2UsInR5cGUiOiJhY2Nlc3MifQ.1xijizY99LXaw0NmWIBmhd8AL_gGmCTj4B1xNL178MQ")

	res, httpcallerr := client.Do(req)
	if httpcallerr != nil {
		fmt.Println("false")
		return 0
	}
	defer res.Body.Close()
	//body, err := ioutil.ReadAll(res.Body)
	var JsonStore JsonResponse
	errors := json.NewDecoder(res.Body).Decode(&JsonStore)
	if errors != nil {
		fmt.Println(errors)
	}

	fmt.Println(JsonStore)

	if JsonStore.Data.Valid_host == true {
		return 1
	} else {
		return 0
	}

	//fmt.Println(string(body))
	//fmt.Println(jsonbody)
	//return true
}

// //export BlehBleh
// func BlehBleh() {

// }

func main() {
	// os.Setenv("FOO", "1")
	// fmt.Println("FOO:", os.Getenv("FOO"))
	// fmt.Println("BAR:", os.Getenv("BAR"))

	// fmt.Println()
	// for _, e := range os.Environ() {
	//     pair := strings.SplitN(e, "=", 2)
	//     fmt.Println(pair)
	// }
	// CallIPLookup(C.CString("abcd"))

	// ...

	// req, err := http.NewRequest("GET", "http://example.com", nil)
	// // ...
	// req.Header.Add("If-None-Match", `W/"wyzzy"`)
	// resp, err := client.Do(req)

	result := IsValidEnvironment()
	fmt.Println(result)

}
