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
			return echo.ErrInternalServerError
		}

		return c.JSONPretty(http.StatusOK, r, "  ")
	})

	e.GET("/standard_codes", func(c echo.Context) error {
		ct := c.QueryParam("codetype")

		r, err := postgres.SelectStandardCodes(ct)
		if err != nil {
			return echo.ErrInternalServerError
		}

		return c.JSONPretty(http.StatusOK, r, "  ")
	})

	e.POST("/client_code", func(c echo.Context) error {
		var cc model.ClientCode

		if err := c.Bind(&cc); err != nil {
			c.Logger().Error(err)
			return echo.ErrBadRequest
		}

		_, err := postgres.InsertClientCode(cc)
		if err != nil {
			return echo.ErrInternalServerError
		}
		return c.String(http.StatusOK, "client code added")
	})

	e.POST("/code_type", func(c echo.Context) error {
		var ct model.CodeType

		if err := c.Bind(&ct); err != nil {
			return echo.ErrBadRequest
		}
		_, err := postgres.InsertCodeType(ct)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "code type added")
	})

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

	e.POST("/standard_code", func(c echo.Context) error {
		var sc model.StandardCode

		if err := c.Bind(&sc); err != nil {
			return err
		}
		_, err := postgres.InsertStandardCode(sc)
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "standard code added")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))

	return nil
}
