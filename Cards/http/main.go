package main

import (
	"io"
	"net/http"
	"os"
)

func main() {

	resp, _ := http.Get("http://google.com/")
	//b := make([]byte, 99999)
	//resp.Body.Read(b)
	//fmt.Println(string(b))
	// Another method of reading the body
	io.Copy(os.Stdout, resp.Body)

}
