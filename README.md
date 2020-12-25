# chronos

`chronos` is a [Golang](https://golang.org) library for developing scheduled jobs which may even be on the micro scale.
The concept is similar to a cron job, in that a job will be trigger on a specified schedule, but there are major
elements which distinguish a chronos job:

- Schedules are more simple
- Minimal overhead
- `Down to the second` granularity

# example

In our example your self-esteem isn't at it's peak. Maybe an hourly reminder about how awesome you are will help!

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
