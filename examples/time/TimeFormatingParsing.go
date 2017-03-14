package main

import (
	"fmt"
	"time"
)

func main() {
	p := fmt.Println
	p("Time formatting")

	t := time.Now()
	p(t.Format(time.RFC3339))

	t1, _ := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1)

	p(t.Format("3:04PM"))

	// Layouts must use the reference time Mon Jan 2 15:04:05 MST 2006
	// to show the pattern with which to format/parse a given time/string.
	// The example time must be exactly as shown: the year 2006, 15 for the hour,
	// Monday for the day of the week, etc.
	p(t.Format("Mon Jan _2 15:04:05 2006"))

	// Mon - day of the week
	// 01 - Month
	// 02 - date
	// 2006 - year
	// 03  - hours [15 - in case you want to print in 24 hr format]
	// 04 - Min
	// 05 - sec
	// 999999 - nanoseconds [The number of 9s will decide the number of digits to be displayed in nanosecond]
	// -7:00 - Time offset from UTC time [It is 5:30 hrs ahead of UTC time. Hence we will see +05:30]
	p(t.Format("Mon 2006-01-02T15:04:05.999999-07:00"))

	form := "3 04 PM"
	t2, _ := time.Parse(form, "8 41 PM")
	p(t2)

	// For purely numeric representations you can also use standard string formatting
	// with the extracted components of the time value.
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e := time.Parse(ansic, "8:41PM")
	p(e)

}
