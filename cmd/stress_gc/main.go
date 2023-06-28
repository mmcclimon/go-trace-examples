package main

import (
	"github.com/mmcclimon/go-trace-examples/internal/tracer"
	"github.com/mmcclimon/go-trace-examples/internal/workloads"
)

func main() {
	tracer.Main("stress-gc", 4, workloads.StressGC(10))
}
