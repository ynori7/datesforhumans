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