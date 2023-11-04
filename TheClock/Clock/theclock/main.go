package main

import (
	"os"
	"time"

	theclock "BadassStuff.com/TheClock/Clock"
)

func main() {
	tm := time.Now()
	theclock.BuildClock(os.Stdout, tm)
}
