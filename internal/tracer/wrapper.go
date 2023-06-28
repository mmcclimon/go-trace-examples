package tracer

import "log"

func Main(filename string, numRoutines int, action TracerFunc) {
	tracer := StartTracing(filename)
	defer tracer.Stop()

	tracer.RunRoutines(numRoutines, action)

	log.Printf("output written to %s", tracer.Filename)
}
