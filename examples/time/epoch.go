package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time epoch example")

	// Getting the secs, millis and nanos since the Unix Epoch

	now := time.Now()
	secs := now.Unix()
	nanos := now.UnixNano()

	fmt.Println(now)

	// there is no unix millis, we will have to convert
	// nanos to millis
	// 1 millisecond = 10^6 nanoseconds
	millis := nanos / 1000000

	fmt.Println(secs)
	fmt.Println(millis)
	fmt.Println(nanos)

	// we can also convert time [secs/nanoseconds/millis] since the epoch
	// into the corresponding time

	fmt.Println(time.Unix(secs, 0))
	fmt.Println(time.Unix(0, nanos))

}
