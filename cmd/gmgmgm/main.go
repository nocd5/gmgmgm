package main

import (
	"flag"
	"fmt"
	"github.com/nocd5/gmgmgm"
	"os"
	"strings"
)

var (
	ignoreCaseOpt   = flag.Bool("i", false, "ignore case")
	forwardMatchOpt = flag.Bool("f", false, "forward match")
)

func main() {
	flag.Parse()
	a := flag.Args()

	if len(a) < 2 {
		fmt.Fprintf(os.Stderr, "Invalid argments\n")
		return
	}

	pat := a[0]
	src := a[1:]

	l := gmgmgm.Match(pat, src, *ignoreCaseOpt, *forwardMatchOpt)

	fmt.Println(strings.Join(l, "\n"))
}
