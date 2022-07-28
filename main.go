package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("url", "", "url parameter")
	flag.Parse()

	if *url == "" {
		fmt.Println("url is required")
		return
	}

	fmt.Println("url:", *url)
}
