package main

import "time"

type WebsiteChecker func(string) bool

func CheckWebSites(wc WebsiteChecker, urls []string) map[string]bool {
	resuts := make(map[string]bool)

	// By giving each anonymous function a parameter for the url - u - and
	// then calling the anonymous function with the url as the argument, we
	// make sure that the value of u is fixed as the value of url for the
	// iteration of the loop that we're launching the goroutine in. u is a
	// copy of the value of url, and so can't be changed.
	for _, url := range urls {
		go func(u string) {
			resuts[u] = wc(u)
		}(url)
	}

	time.Sleep(2 * time.Second)

	return resuts
}
