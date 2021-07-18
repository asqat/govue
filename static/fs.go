package static

import (
	"embed"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

//go:embed dist
var dist embed.FS

func WebFS() error {
	e := echo.New()

	staticConf := middleware.StaticConfig{
		Root:       "dist",
		Filesystem: http.FS(dist),
	}

	e.Use(middleware.StaticWithConfig(staticConf))
	e.Logger.Fatal(e.Start(":8888"))
	return nil
}
