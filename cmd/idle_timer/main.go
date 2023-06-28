package main

import (
	"github.com/mmcclimon/go-trace-examples/internal/tracer"
	"github.com/mmcclimon/go-trace-examples/internal/workloads"
)

func main() {
	tracer.Main("idle-timer", 8, workloads.RandSleep)
}
