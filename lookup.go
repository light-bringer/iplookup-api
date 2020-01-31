package main

import "C"
import (
	"fmt"
	"strings"
	"net/http"
)


func callIPLookup() {
	
}

func main() {
	client := &http.Client{
		//CheckRedirect: redirectPolicyFunc,
	}
	
	resp, err := client.Get("http://example.com")
	// ...
	
	req, err := http.NewRequest("GET", "http://example.com", nil)
	// ...
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)

}