package web

import (
	"embed"
	"fmt"
	"github.com/kaiaverkvist/dployr/internal/logging"
	"github.com/kaiaverkvist/dployr/internal/utils"
	"github.com/kaiaverkvist/dployr/internal/web/context"
	"github.com/kaiaverkvist/dployr/internal/web/routes"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"strings"
	"time"
)

//go:embed assets
var embeddedFiles embed.FS

func embedFS(fs embed.FS) http.FileSystem {
	return http.FS(fs)
}

func CreateServer(container *context.WebDataContainer) *echo.Echo {
	// Embedded file system used to render templates and serve assets
	fs := embedFS(embeddedFiles)

	// Setup a template renderer using the embedded filesystem.
	templateRenderer := NewTemplateRenderer("/assets/templates", fs)

	// Initialize Echo and setup for template rendering.
	e := echo.New()
	e.Logger.SetHeader(logging.LogStyle)
	e.HidePort = true
	e.HideBanner = true
	e.Renderer = templateRenderer

	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &context.DployrContext{c, container}
			return next(cc)
		}
	})

	// Initialize a new file server with our embedded resources
	assetHandler := http.FileServer(fs)

	// Include our routes.
	routes.Routes(e)

	e.GET("/assets/*", echo.WrapHandler(assetHandler))

	go func() {
		err := e.Start("")
		if err != nil {
			log.Error("Unable to start web server: ", err.Error())
			return
		}
	}()

	// Wait if the listener is nil.
	// This prevents nil dereference panic and ensures we don't wait for too long once it gets up.
	var waited time.Duration
	for e.Listener == nil {
		wait := 50 * time.Millisecond
		time.Sleep(wait)
		waited += wait

		if waited > time.Second+2 {
			log.Warn("It has taken more than two seconds to start the web server!")
			log.Warn("This could indicate a problem.")
			return nil
		}
	}

	port := strings.ReplaceAll(e.Listener.Addr().String(), "[::]", "")
	browserUrl := fmt.Sprintf("http://localhost%s", port)

	log.Info("Opening web browser @ ", browserUrl)
	_ = utils.OpenUserBrowser(browserUrl)

	return e
}
