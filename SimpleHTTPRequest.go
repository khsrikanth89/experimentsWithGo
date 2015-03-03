package main

import (
	"fmt"
	"net/http"
	"bytes"
	"io/ioutil"
)

func main(){

resp, err := http.Post("url","application/json",bytes.NewBuffer([]byte(`{"test":"test"}`)))
if err != nil {
	// handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
fmt.Printf("%s %s",body,err)
}

