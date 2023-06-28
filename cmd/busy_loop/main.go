package main

import (
	"github.com/mmcclimon/go-trace-examples/internal/tracer"
	"github.com/mmcclimon/go-trace-examples/internal/workloads"
)

func main() {
	tracer.Main("busy-loop", 4, workloads.BusyLoop)
}
