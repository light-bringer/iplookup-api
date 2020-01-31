package main

import "C"
import (
	"fmt"
	"strings"
	"net/http"
	"os"
)

//export CallIPLookup
func CallIPLookup(myval *C.char)  *C.char{
	// this is an exportable funciton
	/*
	this is a comment
	*/
	fmt.Println(C.GoString(myval))
	return myval

}

//export BlehBleh
func BlehBleh() {
	
}

func main() {
	os.Setenv("FOO", "1")
    fmt.Println("FOO:", os.Getenv("FOO"))
    fmt.Println("BAR:", os.Getenv("BAR"))

    fmt.Println()
    for _, e := range os.Environ() {
        pair := strings.SplitN(e, "=", 2)
        fmt.Println(pair)
	}
	CallIPLookup(C.CString("abcd"))
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}
	
	resp, err := client.Get("http://example.com")
	
	fmt.Println(resp.Body, err)
	// ...
	
	// req, err := http.NewRequest("GET", "http://example.com", nil)
	// // ...
	// req.Header.Add("If-None-Match", `W/"wyzzy"`)
	// resp, err := client.Do(req)

}