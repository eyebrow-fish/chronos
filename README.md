# chronos

`chronos` is a [Golang](https://golang.org) library for developing schedules jobs which are even on a micro scale. The
concept is similar to a cron job, in that a job will be trigger on a specified schedule, but there are major elements
which define a chronos job:

- Schedules are more simple
- Minimal overhead
- `Down to the second` granularity

# example

In this case, maybe your self-esteem isn't doing so hot. Maybe a scheduled reminder will help? In this example, we
create a new `Job` and schedule it to run hourly.

```go
package main

import (
	"context"
	"fmt"
	"github.com/eyebrow-fish/chronos"
)

func main() {
	chronos.
		NewJob(func(ctx context.Context) error {
			fmt.Println("You are awesome!")
			return nil
		}).
		Hourly().
		Run()
}
```
