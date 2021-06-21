package main

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	return url != "waat://furhurterwe.geds"
}

func slowStubWebsitesChecker(_ string) bool {
	time.Sleep(20 * time.Second)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) { // go test -bench=.
	urls := make([]string, 100)
	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	for i := 0; i < b.N; i++ {
		CheckWebSites(slowStubWebsitesChecker, urls)
	}
}

func TestChechWebSites(t *testing.T) {
	websites := []string{
		"http://google.com",
		"https://petrostrak.netlify.app/",
		"waat://furhurterwe.geds",
	}

	want := map[string]bool{
		"http://google.com":               true,
		"https://petrostrak.netlify.app/": true,
		"waat://furhurterwe.geds":         false,
	}

	got := CheckWebSites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, but got %v", want, got)
	}
}
