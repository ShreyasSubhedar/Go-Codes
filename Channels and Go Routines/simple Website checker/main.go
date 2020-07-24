// Simple website checker

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

	for _, site := range urls {
		websiteChecker(site)
	}
}
func websiteChecker(url string) {

	_, err := http.Get(url)
	if err != nil {
		fmt.Println("There might be a downtime at: ", url)
		return
	}
	fmt.Println("Looks Perfect")

}
