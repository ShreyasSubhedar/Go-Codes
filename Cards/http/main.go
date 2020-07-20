package main

import (
	"fmt"
	"net/http"
)

func main() {

	resp, _ := http.Get("http://google.com/")
	b := make([]byte, 99999)
	resp.Body.Read(b)
	fmt.Println(string(b))
}
