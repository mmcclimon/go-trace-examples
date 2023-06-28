package main

import (
	"github.com/mmcclimon/go-trace-examples/internal/tracer"
	"github.com/mmcclimon/go-trace-examples/internal/workloads"
)

func main() {
	tracer.Main("lots-of-timers", 1024, workloads.RandSleep)
}
