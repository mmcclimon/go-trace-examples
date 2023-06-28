package main

import (
	"github.com/mmcclimon/go-trace-examples/internal/util"
)

func main() {
	util.Main("lots-of-timers", 1024, util.RandSleep)
}
