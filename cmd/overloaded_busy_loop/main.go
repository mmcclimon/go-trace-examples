package main

import (
	"github.com/mmcclimon/go-trace-examples/internal/tracer"
	"github.com/mmcclimon/go-trace-examples/internal/workloads"
)

func main() {
	tracer.Main("overloaded-busy-loop", 32, workloads.BusyLoop)
}
