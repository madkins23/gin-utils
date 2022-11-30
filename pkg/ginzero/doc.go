// Package ginzero provides tools for using [zerolog] from within [gin] applications.
// This includes the ginzero.Logger [gin] Middleware to dump request data into [zerolog] and
// the ginzero.Writer IO.Writer to trap low-level [gin] messaging.
//
// [gin]: https://github.com/gin-gonic/gin
// [zerolog]: https://github.com/rs/zerolog
//
// # Import Packages
//
//  import (
//      "github.com/gin-gonic/gin"
//      "github.com/rs/zerolog"
//      "github.com/rs/zerolog/log"
//      "github.com/madkins23/gin-utils/pkg/ginzero"
//  )
//
// # Configure Router
//
// There is some gin logging of non-request issues that just goes to the default logging streams.
// This mostly happens at startup.
// These streams can be replaced with any IO.Writer entity.
// Trap and redirect these streams to zerolog using the following:
//
//  gin.DefaultWriter = ginzero.NewWriter(zerolog.InfoLevel)
//  gin.DefaultErrorWriter = ginzero.NewWriter(zerolog.ErrorLevel)
//
// The basic logging for request traffic in gin is generally handled via middleware.
// The existing default middleware sends request data to the default
// logging streams with some formatting.
// Add the ginzero logger using the following:
//
//  router := gin.New() // not gin.Default()
//  router.Use(ginzero.Logger())
//
// Use gin.New() instead of gin.Default() in this case.
// The latter adds its own logging middleware
// which would conflict with the ginzero middleware.
//
// Add routing configuration after these statements.
// Actual router configuration will depend on the application.
// After configuration run the server.
package ginzero
