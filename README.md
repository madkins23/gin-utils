# gin-utils
Tools and template for constructing applications using [`gin`](https://github.com/gin-gonic/gin).

Can provide:

* Graceful shutdown of [`gin`](https://github.com/gin-gonic/gin) server
* Redirection of [`gin`](https://github.com/gin-gonic/gin) log messages to [`zerolog`](https://github.com/rs/zerolog)
* Template application using [`gin`](https://github.com/gin-gonic/gin) and [`zerolog`](https://github.com/rs/zerolog)

[![Go Report Card](https://goreportcard.com/badge/github.com/madkins23/gin-zerolog)](https://goreportcard.com/report/github.com/madkins23/gin-zerolog)
![GitHub](https://img.shields.io/github/license/madkins23/gin-zerolog)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/madkins23/gin-zerolog)
[![Go Reference](https://pkg.go.dev/badge/github.com/madkins23/gin-zerolog.svg)](https://pkg.go.dev/github.com/madkins23/gin-zerolog)

# Graceful Shutdown

# Logging via `zerolog`

## Usage

Import packages using:

    import (
        "github.com/gin-gonic/gin"
        "github.com/rs/zerolog"
        "github.com/rs/zerolog/log"

        "github.com/madkins23/gin-utils/pkg/ginzero"
    )

## Tools

There is a demo program located in `demo/ginzerolog/ginzerolog.go`.

### Middleware

The basic logging for request traffic in `gin` is generally handled via middleware.
The existing default middleware sends request data to the default
logging streams with some formatting.

Add the `ginzero` logger using the following:

    router := gin.New() // not gin.Default()
    router.Use(ginzero.Logger())

Add routing configuration after these statements.

Use `gin.New()` instead of `gin.Default()`.
The latter adds its own logging middleware
which would conflict with the `ginzero` middleware.

### IO Writer

There is some `gin` logging of non-request issues that just goes to
the default logging streams.
This mostly happens at startup.
These streams can be replaced with any `IO.Writer` entity.

Trap and redirect these streams to `zerolog` using the following:

    gin.DefaultWriter = ginzero.NewWriter(zerolog.InfoLevel)
    gin.DefaultErrorWriter = ginzero.NewWriter(zerolog.ErrorLevel)
    router := gin.New() // or gin.Default() if not using ginzero.Logger()


## Application Template
