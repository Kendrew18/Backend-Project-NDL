package routes

import (
	controllers "Backend-Project-NDL/controller"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-NDL")
	})

	NDL := e.Group("/NDL")
	Rekap := e.Group("/Rekap")

	//NDL
	NDL.POST("/read-excel", controllers.ReadEXCEL)

	NDL.POST("/input-ndl", controllers.InputNDL)

	NDL.GET("/NDL", controllers.ReadNDL)

	//Rekap
	Rekap.GET("/Rekap", controllers.ReadRekap)

	return e
}
