package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Basic time api")

	p := fmt.Println

	// get current time
	now := time.Now()
	p(now)

	then := time.Date(2016, 11, 13, 12, 15, 00, 00, time.UTC)
	p(then)

	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())
	p(then.Location())

	p(then.Weekday())

	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	diff := now.Sub(then)
	p(diff)

	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	p(then.Add(diff))
	p(then.Add(-diff))

}
