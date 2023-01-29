package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
	Fetch2(e Entry, s state)
}

type Entry struct {
	url   string
	depth int
}

type state struct {
	in   chan Entry
	out  chan Entry
	err  chan error
	done chan struct{}
	wg   *sync.WaitGroup
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(e Entry, depth int, fetcher Fetcher, state state) {
	var pending []Entry

	if depth < 0 {
		return
	}
	seen := make(map[string]bool)

	pending = append(pending, e)

	numRunning := 0
	for {
		if len(pending) > 0 {
			entry := pending[0]
			pending = pending[1:]
			if entry.depth < depth {
				go fetcher.Fetch2(entry, state)
				numRunning++
			}
		}

		select {
		case res := <-state.out:
			// if !seen[res.url] {
			if true {
				pending = append(pending, res)
				seen[res.url] = true
			} else {
				fmt.Printf("cached: %s\n", res.url)
			}
		case err := <-state.err:
			fmt.Println(err)
		case <-state.done:
			// fmt.Printf("  done pending:%d running:%d\n", len(pending), numRunning-1)
			if len(pending) == 0 {
				if numRunning--; numRunning == 0 {
					return
				}
			}
		}

	}
}
func main() {
	uin := make(chan Entry)
	uout := make(chan Entry)
	err := make(chan error)
	done := make(chan struct{})
	var wg sync.WaitGroup
	s := state{uin, uout, err, done, &wg}
	Crawl(Entry{"https://golang.org/", 0}, 4, fetcher, s)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch2(e Entry, s state) {
	defer func() { s.done <- struct{}{} }()
	if res, ok := f[e.url]; ok {
		time.Sleep(time.Second)
		fmt.Printf("found: %s %d\n", e.url, e.depth)
		for _, u := range res.urls {
			s.out <- Entry{u, e.depth + 1}
		}
	} else {
		s.err <- fmt.Errorf("not found: %s", e.url)
	}
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		// time.Sleep(time.Second)
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
