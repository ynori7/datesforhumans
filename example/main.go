package main

import (
	"fmt"
	"time"

	"github.com/ynori7/datesforhumans"
)

func main() {
	Basic()
	Repeat()
	Range()
	AdditionalParsingCases()
}

func Basic() {
	now := time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC) //2023-07-15 17:03:00

	fmt.Printf("Basic (from %s)\n", now.Format(time.DateTime))

	t := datesforhumans.ParseDate(now, "next monday").At("10pm").Time()                               //t is a standard time.Time
	fmt.Println(`datesforhumans.ParseDate(now, "next monday").At("10pm") -`, t.Format(time.DateTime)) //prints 2023-07-17 22:00:00
	t2 := datesforhumans.ParseDate(now, "next monday at 10pm").Time()
	fmt.Println(`datesforhumans.ParseDate(now, "next monday at 10pm") -`, t2.Format(time.DateTime)) //prints 2023-07-17 22:00:00
	t3 := datesforhumans.FromNow("next monday at 10pm").Time()
	fmt.Println(`datesforhumans.FromNow("next monday at 10pm") -`, t3.Format(time.DateTime)) //prints 2023-07-17 22:00:00

	fmt.Println("")
}

func Repeat() {
	now := time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC) //2023-07-15 17:03:00

	fmt.Printf("Repeat (from %s)\n", now.Format(time.DateTime))

	repeated := datesforhumans.ParseDate(now, "next monday").At("10pm").Repeat(datesforhumans.ParseDate(now, "next August").Time())
	fmt.Println(`datesforhumans.ParseDate(now, "next monday").At("10pm").Repeat(datesforhumans.ParseDate(now, "next August"): `)
	for _, r := range repeated {
		fmt.Println(r.Time().Format(time.DateTime))
	}

	fmt.Println("")

	fmt.Println("Repeat with time")
	repeated2 := datesforhumans.ParseDate(now, "next monday at 10pm").Repeat(datesforhumans.ParseDate(now, "next August").Time())
	fmt.Println(`datesforhumans.ParseDate(now, "next monday at 10pm").Repeat(datesforhumans.ParseDate(now, "next August"): `)
	for _, r := range repeated2 {
		fmt.Println(r.Time().Format(time.DateTime))
	}

	fmt.Println("")
}

func Range() {
	now := time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC) //2023-07-15 17:03:00

	fmt.Printf("Range (from %s)\n", now.Format(time.DateTime))

	r := datesforhumans.ParseRange(now, "next monday at 10pm", "next tuesday at 11pm")
	fmt.Println(`datesforhumans.ParseRange(now, "next monday at 10pm", "next tuesday at 11pm"):`, r.Start.Time().Format(time.DateTime), "-", r.End.Time().Format(time.DateTime)) //prints 2023-07-17 22:00:00 - 2023-07-18 23:00:00

	fmt.Println("")

	fmt.Println("Repeat")
	repeatedRange := r.Repeat(datesforhumans.ParseDate(now, "next August").Time())
	fmt.Println(`r.Repeat(datesforhumans.ParseDate(now, "next August"): `)
	for _, r := range repeatedRange {
		fmt.Println(r.Start.Time().Format(time.DateTime), "-", r.End.Time().Format(time.DateTime))
	}

	fmt.Println("")
}

func AdditionalParsingCases() {
	now := time.Date(2023, time.January, 30, 17, 4, 0, 0, time.UTC)

	fmt.Printf("Additional Parsing Cases (from %s)\n", now.Format(time.DateTime))

	t3 := datesforhumans.ParseDate(now, "next February").Time()
	fmt.Println(`datesforhumans.ParseDate(now2, "next February") -`, t3.Format(time.DateTime)) //prints february 1st
	t3 = datesforhumans.ParseDate(now, "in 1 month").Time()
	fmt.Println(`datesforhumans.ParseDate(now2, "in 1 month") -`, t3.Format(time.DateTime)) //prints march 1st

	fmt.Println("")
}
