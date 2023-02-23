package util

import (
	"fmt"
	"time"
)

type timer struct {
	totalDuration time.Duration
	lastTime      time.Time
}

func (t *timer) count(s string, i int) {
	currentTime := time.Now()

	if t.lastTime.IsZero() {
		t.lastTime = currentTime
	}

	elapsedTime := currentTime.Sub(t.lastTime)
	t.totalDuration += elapsedTime
	avgDuration := t.totalDuration / time.Duration(i+1)

	fmt.Printf("%s — Total: %v — avg/item: %v\n", s, t.totalDuration, avgDuration)

	t.lastTime = currentTime
}
