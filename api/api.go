package api

import (
	"net/http"
	"translation/model"
	"translation/tools/postgres"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartAPI() error {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route => handler
	e.GET("/source_systems", func(c echo.Context) error {
		r, err := postgres.SelectSourceSystems()
		if err != nil {
			return err
		}

		return c.JSONPretty(http.StatusOK, r, "  ")
	})

	// Route => handler
	e.POST("/source_system", func(c echo.Context) error {
		var ss model.SourceSystem

		if err := c.Bind(&ss); err != nil {
			return err
		}
		_, err := postgres.InsertSourceSystem(ss)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "source system added")
	})

	// Route => handler
	e.POST("/source_system_application", func(c echo.Context) error {
		var ssa model.SourceSystemApplication

		if err := c.Bind(&ssa); err != nil {
			return err
		}
		_, err := postgres.InsertSourceSystemApplication(ssa)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "source system application added")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
