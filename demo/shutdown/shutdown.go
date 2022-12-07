/*
demo of shutdown package to support graceful server shutdown.

Once running the application responds to http://:55555/ping with:

	{"message":"pong"}

[gin]: https://gin-gonic.com/docs/
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/madkins23/gin-utils/pkg/handler"
	"github.com/madkins23/gin-utils/pkg/shutdown"
)

func main() {
	// Initialize for graceful shutdown.
	graceful := &shutdown.Graceful{}
	graceful.Initialize()
	defer graceful.Close()

	router := gin.Default()
	router.GET("/ping", handler.Ping)
	router.GET("/exit", handler.Exit)

	fmt.Println("Ping:  http://localhost:55555/exit")
	fmt.Println("Ping:  http://localhost:55555/ping")
	if err := graceful.Serve(router, 55555); err != nil {
		fmt.Printf("Server failure: %s\n", err)
	}

	// At this point either a <ctrl>-C (interrupt) or /exit will shut down gracefully.
}
