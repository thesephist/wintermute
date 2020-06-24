package main

import (
	"log"
	"time"
)

type timer struct {
	start time.Time
}

func newTimer() timer {
	return timer{
		start: time.Now(),
	}
}

func (t timer) finish(label string) {
	end := time.Now()
	log.Printf("[timing] %s: %dus\n",
		label,
		end.Sub(t.start)/time.Microsecond)
}
