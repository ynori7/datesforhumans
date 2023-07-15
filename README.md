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

Example usage (see it in `example/main.go`):

```go
now := time.Now() //e.g. 2023-07-15 17:03:00
t := datesforhumans.ParseDate(now, "next monday").At("10pm").Time() //t is a standard time.Time
fmt.Println(t.Format(time.DateTime)) //prints 2023-07-17 22:00:00
```

Note that if it's July 17 and you request "next August", you'll get August 17.

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
// 2023-08-07 22:00:00
// 2023-08-14 22:00:00
```

It's also possible to create date ranges, for example:

```go
now := time.Now() //e.g. 2023-07-15 17:03:00
r := ParseRange(now, "next monday", "10pm", "next tuesday", "11pm")
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
// 2023-08-07 22:00:00 2023-08-08 23:00:00
// 2023-08-14 22:00:00 2023-08-15 23:00:00
```