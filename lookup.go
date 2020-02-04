package main

import "C"
import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

//export CallIPLookup
func CallIPLookup(IPAddress *C.char) *C.char {
	// this is an exportable funciton
	/*
		this is a comment
	*/
	fmt.Println(C.GoString(IPAddress))
	return IPAddress

}

func IPLookUPCall() bool {

	url := "https://nodered.fierylab.io/go-or-no-go"
	method := "GET"
	client := &http.Client{}
	payload := strings.NewReader("")
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer key3456789876545678")

	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	fmt.Println(string(body))
	return true
}

//export BlehBleh
func BlehBleh() {

}

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

	result := IPLookUPCall()
	fmt.Println(result)

}
