package main

import (
	"fmt"
	"time"
)

func expensiveCall() {
	fmt.Println("working for some time...")
	time.Sleep(time.Second*1)
	fmt.Println("done")
}

func main() {
	t0 := time.Now()
	expensiveCall()
	t1 := time.Now()
	fmt.Printf("The call took %v to run.\n", t1.Sub(t0))


	d, err := time.ParseDuration("1h15m30.918273645s")
	if err != nil {
		panic(err)
	}

	fmt.Println(d)

	round := []time.Duration{
		time.Nanosecond,
		time.Microsecond,
		time.Millisecond,
		time.Second,
		2 * time.Second,
		time.Minute,
		10 * time.Minute,
		time.Hour,
	}

	for _, r := range round {
		fmt.Printf("d.Round(%6s) = %s\n", r, d.Round(r).String())
	}
	// Output:
	// d.Round(   1ns) = 1h15m30.918273645s
	// d.Round(   1µs) = 1h15m30.918274s
	// d.Round(   1ms) = 1h15m30.918s
	// d.Round(    1s) = 1h15m31s
	// d.Round(    2s) = 1h15m30s
	// d.Round(  1m0s) = 1h16m0s
	// d.Round( 10m0s) = 1h20m0s
	// d.Round(1h0m0s) = 1h0m0s


	_, month, day := time.Now().Date()
	if month == time.June && day == 17 {
		fmt.Println("Today is the Day!")
	}

	t := time.Date(2009, time.November, 10, 23, 0, 0, 0, time.UTC)
	fmt.Printf("Go launched at %s\n", t.Local())
	fmt.Printf("Go launched at %s\n", t)
	fmt.Printf("If you compare time in different timezones the Equal() says equal: %v\n", t.Equal(t.Local()))

	fmt.Println(t.Add(time.Second).After(t.UTC()))
	fmt.Println(t.Before(t))

	// Output: Go launched at 2009-11-10 15:00:00 -0800 PST




	// Parse a time value from a string in the standard Unix format.
	t, err = time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}
	// time.Time's Stringer method is useful without any format.
	fmt.Println("default format:", t)

	// Predefined constants in the package implement common layouts.
	fmt.Println("Unix format:", t.Format(time.UnixDate))


	/*
		If you have worked with time/date formatting/parsing in other languages you might have noticed that the other
		languages use special placeholders for time/date formatting. For eg ruby language uses

		%d for day
		%Y for year
		etc
		Golang, instead of using codes such as above, uses date and time format placeholders that look like date and time only.
		Go uses standard time, which is:
		Mon Jan 2 15:04:05 MST 2006  (MST is GMT-0700)
		or
		01/02 03:04:05PM '06 -0700

		So if you notice go uses
		01 for day of the month ,
		02 for the month
		03 for hours ,
		04 for minutes
		05 for second
		and so on
	*/
	// the cusrom layout.
	const longForm = "Jan 2, 2006 at 3:04pm (MST)"
	t, _ = time.Parse(longForm, "Feb 3, 2013 at 7:54pm (PST)")
	fmt.Println(t)

	// shortForm is another way the reference time would be represented
	// in the desired layout; it has no time zone present.
	// Note: without explicit zone, returns time in UTC.
	const shortForm = "2006-Jan-02"
	t, _ = time.Parse(shortForm, "2013-Feb-03")
	fmt.Println(t)


	loc, _ := time.LoadLocation("Europe/Berlin")
	// This will look for the name CEST in the Europe/Berlin time zone.
	t, _ = time.ParseInLocation(longForm, "Jul 9, 2012 at 5:02am (CEST)", loc)
	fmt.Println(t)
	// Note: without explicit zone, returns time in given location.
	t, _ = time.ParseInLocation(shortForm, "2012-Jul-09", loc)
	fmt.Println(t)

	fmt.Println(time.Now().Truncate(time.Hour*24))
	fmt.Println(time.Now().Truncate(time.Hour*24).UTC())
}

