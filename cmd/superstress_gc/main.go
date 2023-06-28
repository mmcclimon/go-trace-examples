package main

import (
	"fmt"
	"strings"

	"github.com/mmcclimon/go-trace-examples/internal/tracer"
	"github.com/mmcclimon/go-trace-examples/internal/workloads"
)

var warning = `
NOTE: this might produce an unparseable trace, in which case you'll see

    failed to parse trace: time stamps out of order

This is a bug in go itself: see https://github.com/golang/go/issues/16755
It's fixed (or at least less bad) in gotip (the newest version of Go).

Good luck!
`

func main() {
	fmt.Println(strings.TrimSpace(warning) + "\n")

	tracer.Main("superstress-gc", 8, workloads.StressGC(5000))
}
