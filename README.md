# go trace examples

This is just a place for me to dump examples of what the go execution tracer
looks like under various kinds of workload. Everything just uses the standard
library.

To run:

1. `go run cmd/idle_timer` (or whatever)
2. This will output some stuff to the terminal, ending with a line like
   `output written to traces/idle-timer.trace`
3. Start the trace viewer by running `go tool trace traces/idle-timer.trace` (the
   filename from the previous line).
4. This will open a browser, where you can click around.


The execution tracer is [effectively undocumented](https://github.com/golang/go/issues/16526),
and so the best way to learn it is by searching around for various other
explanations. Here are some things I've found.

- [Rhys Hiltner - An Introduction to "go tool trace"](https://www.youtube.com/watch?v=V74JnrGTwKA) (video)
- [Understanding Go Execution Tracer by Example](https://www.sobyte.net/post/2022-03/go-execution-tracer-by-example/) (article)


