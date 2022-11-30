// Package shutdown provides a mechanism for graceful shutdown of [gin].
//
// Support graceful shutdown of [gin] during an interrupt signal.
// This tool captures the Linux SIGINT and SIGKILL signals,
// so it won't work (completely) with Apple or Windows.
//
// [gin]: https://github.com/gin-gonic/gin
//
// # Import Packages
//
//  import (
//      "github.com/gin-gonic/gin"
//      "github.com/madkins23/gin-utils/pkg/shutdown"
//  )
//
// # Initialize Server
//
//  graceful := &shutdown.Graceful{}
//  graceful.Initialize()
//  defer graceful.Close()
//
// The defer statement will automatically shut down and cleanup after the server
// when the enclosing scope is exited.
//
// # Configure Router
//
//  router := gin.Default()
//  router.GET("/exit", handler.Exit)
//
// Actual router configuration will depend on the application.
//
// # Run Server:
//
//  if err := graceful.Serve(router, port); err != nil {
//      log.Fatal().Err(err).Msg("Running gin server")
//  }
//
// The server will shut down gracefully if the process receives either a
// SIGINT or SIGKILL signal.
package shutdown
