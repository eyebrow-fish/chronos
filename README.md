# chronos

`chronos` is a [Golang](https://golang.org) library for developing schedules jobs of
all forms.

# example

In this case, maybe your self-esteem isn't doing so hot. Maybe a scheduled reminder 
will help? In this example, we create a new `Job` and schedule it to run hourly.

```go
package main

import (
	"fmt"
	"github.com/eyebrow-fish/chronos"
)

func main() {
	chronos.
		NewJob(func() {
			fmt.Println("You are awesome!")
		}).
		Hourly().
		Run()
}
```
