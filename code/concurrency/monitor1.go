package main

import (
	"flag"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Parse all arguments
	flag.Parse()

	// flag.Args contains all non-flag arguments
	sites := flag.Args()

	// Lets keep a reference to when we started
	start := time.Now()

	for _, site := range sites {
		// start a timer for this request
		begin := time.Now()
		// Retreive the site
		_, err := http.Get(site)
		if err != nil {
			fmt.Println(site, err)
		} else {
			fmt.Printf("Site %q took %s to retrieve.\n", site, time.Since(begin))
		}
	}

	fmt.Printf("Entire process took %s\n", time.Since(start))
}
