package tracer

import (
	"context"
	"log"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

type Tracer struct {
	Filename string
	C        context.Context
	cancel   context.CancelFunc
	wg       sync.WaitGroup
}

type TracerFunc func(context.Context, *sync.WaitGroup, int)

func StartTracing(traceFilename string) *Tracer {
	err := os.Mkdir("traces", 0755)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	filename := "traces/" + traceFilename + ".trace"
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	return &Tracer{
		Filename: filename,
		C:        ctx,
		cancel:   cancel,
	}
}

func (t *Tracer) Stop() {
	t.cancel()
	trace.Stop()
}

func (t *Tracer) RunRoutines(n int, fn TracerFunc) {
	log.Printf("will run %d routines", n)

	for i := 0; i < n; i++ {
		t.wg.Add(1)
		go fn(t.C, &t.wg, i)
	}

	t.wg.Wait()
	log.Println("routines finished")
}
