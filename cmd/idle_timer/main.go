package main

import (
	"github.com/mmcclimon/go-trace-examples/internal/util"
)

func main() {
	util.Main("idle-timer", 8, util.RandSleep)
}
