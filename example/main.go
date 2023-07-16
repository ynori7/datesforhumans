package main

import (
	"fmt"
	"time"

	"github.com/ynori7/datesforhumans"
)

func main() {
	fmt.Println("Basic")
	now := time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC)        //2023-07-15 17:03:00
	t := datesforhumans.ParseDate(now, "next monday").At("10pm").Time() //t is a standard time.Time
	fmt.Println(t.Format(time.DateTime))                                //prints 2023-07-17 22:00:00
	t2 := datesforhumans.ParseDate(now, "next monday at 10pm").Time()   //t is a standard time.Time
	fmt.Println(t2.Format(time.DateTime))                               //prints 2023-07-17 22:00:00

	fmt.Println("\nRepeat")
	repeated := datesforhumans.ParseDate(now, "next monday").At("10pm").Repeat(datesforhumans.ParseDate(now, "next August").Time())
	for _, r := range repeated {
		fmt.Println(r.Time().Format(time.DateTime))
	}

	fmt.Println("\nRepeat with time")
	repeated2 := datesforhumans.ParseDate(now, "next monday at 10pm").Repeat(datesforhumans.ParseDate(now, "next August").Time())
	for _, r := range repeated2 {
		fmt.Println(r.Time().Format(time.DateTime))
	}

	fmt.Println("\nRange")
	r := datesforhumans.ParseRange(now, "next monday at 10pm", "next tuesday at 11pm")
	fmt.Println(r.Start.Time().Format(time.DateTime), r.End.Time().Format(time.DateTime)) //prints 2023-07-17 22:00:00 2023-07-18 23:00:00

	fmt.Println("\nRepeat")
	repeatedRange := r.Repeat(datesforhumans.ParseDate(now, "next August").Time())
	for _, r := range repeatedRange {
		fmt.Println(r.Start.Time().Format(time.DateTime), r.End.Time().Format(time.DateTime))
	}

	fmt.Println("\nParse")
	now2 := time.Date(2023, time.January, 30, 17, 4, 0, 0, time.UTC)
	t3 := datesforhumans.ParseDate(now2, "next February").Time()
	fmt.Println(t3.Format(time.DateTime)) //prints february 1st
	t3 = datesforhumans.ParseDate(now2, "in 1 month").Time()
	fmt.Println(t3.Format(time.DateTime)) //prints march 1st
}
