// Go routine website checker

package main

import (
	"fmt"
	"net/http"
)

func main() {

	urls := []string{
		"http://google.com",
		"http://amazon.com",
		"http://netflix.com",
		"http://shreyassubhedar.rf.gd",
		"http://facebook.com",
	}
	c := make(chan string)
	for _, site := range urls {
		go websiteChecker(site, c)
	}
	i := 0
	for i < len(urls) {
		fmt.Println(<-c)
		i++
	}
}
func websiteChecker(url string, c chan string) {

	_, err := http.Get(url)
	if err != nil {
		//fmt.Println("There might be a downtime at: ", url)
		c <- "There might be a downtime at: " + url
		return
	}
	//fmt.Println(url, ": Looks Perfect")
	c <- url + ": Looks Perfect"
}
