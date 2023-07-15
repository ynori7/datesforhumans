package main

import (
	"fmt"
	"time"

	"github.com/ynori7/datesforhumans"
)

func main() {
	now := time.Date(2023, time.July, 15, 17, 4, 0, 0, time.UTC)        //2023-07-15 17:03:00
	t := datesforhumans.ParseDate(now, "next monday").At("10pm").Time() //t is a standard time.Time
	fmt.Println(t.Format(time.DateTime))                                //prints 2023-07-17 22:00:00
}
