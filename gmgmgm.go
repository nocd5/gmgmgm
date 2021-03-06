package gmgmgm

import (
	"github.com/koron/gomigemo/embedict"
	"github.com/koron/gomigemo/migemo"
	"regexp"
	"runtime"
	"sort"
	"sync"
)

func Match(pat string, src []string, i bool, f bool) []string {
	// Load embedded dictionary.
	d, err := embedict.Load()
	if err != nil {
		panic(err)
	}
	// Get regex pattern with migemo
	p, err := migemo.Pattern(d, pat)
	if err != nil {
		panic(err)
	}

	if f {
		p = "^" + p
	}
	if i {
		p = "(?i)" + p
	}

	re, err := regexp.Compile(p)
	if err != nil {
		panic(err)
	}

	n := runtime.NumCPU()
	s := make(chan bool, n)
	q := make(chan string, 1)

	var wg sync.WaitGroup
	for _, e := range src {
		wg.Add(1)
		go func(_e string) {
			s <- true
			defer func() { <-s }()

			if re.MatchString(_e) {
				q <- _e
			} else {
				wg.Done()
			}
		}(e)
	}

	var l []string
	go func() {
		for t := range q {
			l = append(l, t)
			wg.Done()
		}
	}()

	wg.Wait()
	sort.Strings(l)
	return l
}
