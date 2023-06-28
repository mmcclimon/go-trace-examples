package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/mmcclimon/go-trace-examples/internal/server"
	"github.com/mmcclimon/go-trace-examples/internal/tracer"
	"github.com/mmcclimon/go-trace-examples/internal/workloads"
)

const addr = "localhost:11999"

func main() {
	server := server.ServeHTTP(addr, responder)
	defer server.Shutdown()

	tracer.Main("network-ok", 8, workloads.Network("http://"+addr))
}

func responder(w http.ResponseWriter, r *http.Request) {
	// simulate network latency
	sleepMS := time.Duration(rand.Intn(150)+50) * time.Millisecond
	time.Sleep(sleepMS)

	fmt.Fprintln(w, "ok")
}
