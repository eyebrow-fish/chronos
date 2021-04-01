# chronos

`chronos` is a [Golang](https://golang.org) "batteries included" library for
developing scheduled jobs using
[Cron Notation](https://en.wikipedia.org/wiki/Cron).

# development

All `jobs` should be defined with a unique name and a schedule of your
choosing. At the moment there is no further configuration for jobs, so
all you simply need to do is:

```go
...

chronos.Job("Find Dog Photos", "*/1 * * * *", 
	func(ctx context.Content) error { ... },
)

...
```

Once all jobs are configured in code, the last thing to do is to run the
`scheduler`. The scheduler manages the runtimes of all of the jobs' tasks.
The following should be run after all definitions are configured because it
blocks indefinitely.

```go
...

chronos.Launch(":8080")

...
```

Another import feature of the scheduler is that it also exposes health and
statistics via HTTP. This is why we pass `":8080"` as a parameter to `Launch`.
