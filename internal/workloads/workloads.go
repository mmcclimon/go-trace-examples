package workloads

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/mmcclimon/go-trace-examples/internal/tracer"
)

func RandSleep(ctx context.Context, i int) {
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

func BusyLoop(ctx context.Context, i int) {
	sum := 0

	for {
		select {
		case <-ctx.Done():
			log.Printf("routine %d: got up to %e", i, float64(sum))
			return
		default:
			for i := 0; i < 10_000; i++ {
				sum++
			}
		}
	}
}

func StressGC(sliceLen int) tracer.TracerFunc {
	return func(ctx context.Context, i int) {
		appends := 0

		for {
			select {
			case <-ctx.Done():
				log.Printf("routine %d: appended %e ", i, float64(appends))
				return
			default:
				s := make([]int, 0)
				for i := 0; i < sliceLen; i++ {
					s = append(s, i)
				}

				appends += len(s)
			}
		}
	}
}
