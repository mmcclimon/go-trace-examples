package main

import (
	"context"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/mmcclimon/go-trace-examples/internal/util"
)

func main() {
	util.Main("idle-timer", 8, hangAround)
}

func hangAround(ctx context.Context, wg *sync.WaitGroup, i int) {
	defer wg.Done()

	interval := time.Duration(rand.Intn(2500)+500) * time.Millisecond
	ticker := time.NewTicker(interval)

	for {
		select {
		case <-ticker.C:
			log.Printf("routine %d: tick after %s", i, interval)
		case <-ctx.Done():
			return
		}
	}
}
