// Go routine and channels website checker

package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	urls := []string{
		"http://google.com",
		"http://stackoverflow.com",
		"http://netflix.com",
		"http://shreyassubhedar.rf.gd",
		"http://facebook.com",
	}
	c := make(chan string)
	for _, site := range urls {
		go websiteChecker(site, c)
	}
	//  Now, Other can understand the workflow
	for l := range c {
		//time.Sleep(time.Second)  Bad practices to sleep the thread because it throttle the other go-routines output.
		go func(link string) {
			time.Sleep(time.Second)
			websiteChecker(link, c) //Always pass the argument to other routine
		}(l)
	}
}
func websiteChecker(url string, c chan string) {

	_, err := http.Get(url)
	if err != nil {
		fmt.Println("There might be a downtime at: ", url)
		c <- url
		return
	}
	fmt.Println(url, ": Looks Perfect")
	c <- url
}
