# Dates For Humans
This library is a helper to transform natural language date time info into real times.

### Usage
You can take any base time (e.g. `time.Now()`) and transform it using natural language. For example:

- 10 minutes ago
- 3 hours from now
- next week
- next monday
- last September
- this wednesday
- in 15 seconds
- next year
- next monday at 10pm
- today at noon

Example usage (see it in `example/main.go`):

```go
now := time.Now() //e.g. 2023-07-15 17:03:00
t := datesforhumans.ParseDate(now, "next monday").At("10pm").Time() //t is a standard time.Time
fmt.Println(t.Format(time.DateTime)) //prints 2023-07-17 22:00:00

//alternatively
t2 := datesforhumans.ParseDate(now, "next monday at 10pm").Time() //t is a standard time.Time
fmt.Println(t.Format(time.DateTime)) //prints 2023-07-17 22:00:00
```

This pattern can also be repeated until a given date:

```go
repeated := datesforhumans.ParseDate(now, "next monday").At("10pm").Repeat(datesforhumans.ParseDate(now, "next August").Time())
for _, r := range repeated {
    fmt.Println(r.Time().Format(time.DateTime))
}
// Prints:
// 2023-07-17 22:00:00
// 2023-07-24 22:00:00
// 2023-07-31 22:00:00
```

It's also possible to create date ranges, for example:

```go
now := time.Now() //e.g. 2023-07-15 17:03:00
r := ParseRange(now, "next monday at 10pm", "next tuesday at 11pm")
fmt.Println(r.Start.Time().Format(time.DateTime), r.End.Time().Format(time.DateTime)) //prints 2023-07-17 22:00:00
```

The ranges can also automatically be repeated until a given date:

```go
repeated := r.Repeat(datesforhumans.ParseDate(now, "next August").Time())
for _, r := range repeated {
    fmt.Println(r.Start.Time().Format(time.DateTime), r.End.Time().Format(time.DateTime))
}
// Prints:
// 2023-07-17 22:00:00 2023-07-18 23:00:00
// 2023-07-24 22:00:00 2023-07-25 23:00:00
// 2023-07-31 22:00:00 2023-08-01 23:00:00
```

### Caveats
- Note that if it's July 17 and you request "next August", you'll get August 1. However if you say "in 1 month" you'll get August 17. It's worth noting there are edge cases where it's January 30 and you say "in 1 month" (returns March 2nd).
- Results are unpredictable if you give it nonsensical strings like "in 2 weeks ago" or "next potato at 3pm".