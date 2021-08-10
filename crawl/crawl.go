package main

import (
	"fmt"
	"time"
)

type Fetcher interface {
	Fetch2(e Entry, s state)
}

type Entry struct {
	url   string
	depth int
}

type state struct {
	out  chan Entry
	err  chan error
	done chan struct{}
}

func Crawl(e Entry, depth int, fetcher Fetcher, state state) {
	var pending []Entry

	if depth < 0 {
		return
	}
	seen := make(map[string]bool)

	pending = append(pending, e)

	ngo := 0
	for {
		if len(pending) > 0 {
			entry := pending[0]
			pending = pending[1:]
			if entry.depth < depth {
				if !seen[entry.url] {
					go fetcher.Fetch2(entry, state)
					seen[entry.url] = true
					ngo++
				} else {
					fmt.Printf("cached: %s\n", entry.url)
				}
			}
		}

		select {
		case res := <-state.out:
			pending = append(pending, res)
		case err := <-state.err:
			fmt.Println(err)
		case <-state.done:
			ngo--
			// fmt.Printf("  done pending:%d running:%d\n", len(pending), ngo-1)
			if len(pending) == 0 && ngo == 0 {
				return
			}
		}
	}
}
func main() {
	uout := make(chan Entry)
	err := make(chan error)
	done := make(chan struct{})
	s := state{uout, err, done}
	Crawl(Entry{"https://golang.org/", 0}, 4, fetcher, s)
}

type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch2(e Entry, s state) {
	defer func() { s.done <- struct{}{} }()
	if res, ok := f[e.url]; ok {
		time.Sleep(time.Second)
		fmt.Printf("found: %s %d %s\n", e.url, e.depth, res.body)
		for _, u := range res.urls {
			s.out <- Entry{u, e.depth + 1}
		}
	} else {
		s.err <- fmt.Errorf("not found: %s", e.url)
	}
}

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
