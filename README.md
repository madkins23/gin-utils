# gin-utils
Tools and template for constructing applications using [`gin`](https://github.com/gin-gonic/gin).

Can provide:

* Graceful shutdown of [`gin`](https://github.com/gin-gonic/gin) server
* Redirection of [`gin`](https://github.com/gin-gonic/gin) log messages to [`zerolog`](https://github.com/rs/zerolog)
* Some simple [`gin`](https://github.com/gin-gonic/gin) handlers
* Template application using [`gin`](https://github.com/gin-gonic/gin) and [`zerolog`](https://github.com/rs/zerolog)

[![Go Report Card](https://goreportcard.com/badge/github.com/madkins23/gin-utils)](https://goreportcard.com/report/github.com/madkins23/gin-utils)
![GitHub](https://img.shields.io/github/license/madkins23/gin-utils)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/madkins23/gin-utils)
[![Go Reference](https://pkg.go.dev/badge/github.com/madkins23/gin-utils.svg)](https://pkg.go.dev/github.com/madkins23/gin-utils)

# Graceful Shutdown

Support graceful shutdown of `gin` during an interrupt signal.
This tool captures the Linux `interrupt` and `kill` signals,
so it won't work (completely) with Apple or Windows.

There is a demo program located in `demo/shutdown/shutdown.go`.

See package `shutdown` [documentation](https://pkg.go.dev/github.com/madkins23/gin-utils/pkg/shutdown) for more details.

# Logging via `zerolog`

Support for connecting gin logging to `zerolog`.
This includes request-logging middleware and
the capture and reprocessing of `stderr` and `stdout` streams.

There is a demo program located in `demo/ginzero/ginzero.go`.

See package `ginzero` [documentation](https://pkg.go.dev/github.com/madkins23/gin-utils/pkg/ginzero) for more details.

## Simple Handlers

A small collection of simpler `gin` handlers is provided in the `handler` package.
These include:

* `Ping` handler to return a 200 "Pong!" response.
* `Exit` handler to send a `SIGINT` signal to the current process,
  thereby ending the service.

See package `handler` [documentation](https://pkg.go.dev/github.com/madkins23/gin-utils/pkg/handler) for more details.

## Application Template

There is an executable template application in `cmd/template/template.go`.

The template utilizes:
* `ginzero` to redirect all `gin` messaging through `zerolog`,
* `shutdown` to implement graceful shutdown, and
* `handler` for a few simple handlers (remove later)

This template is intended to be copied into another project as a starting point.
