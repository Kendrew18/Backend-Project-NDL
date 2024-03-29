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
	PO := e.Group("/PO")
	Template := e.Group("/Template")

	//NDL
	NDL.POST("/read-excel", controllers.ReadEXCEL)

	NDL.POST("/input-ndl", controllers.InputNDL)

	NDL.GET("/NDL", controllers.ReadNDL)

	NDL.GET("/page-ndl", controllers.PageNo)

	NDL.PUT("/update-ndl", controllers.UpdateNDL)

	NDL.GET("/Read-NDL-wsno", controllers.ReadNDLwsno)

	//Rekap
	Rekap.GET("/Rekap", controllers.ReadRekap)

	Rekap.PUT("/update-status-rkp", controllers.UpdateStatusRekap)

	//PO-supplier
	PO.POST("/input-PO-supplier", controllers.InputPOsupplier)

	PO.GET("/read-PO", controllers.ReadPO)

	PO.GET("/read-layer-PO", controllers.LyrPO)

	//Template
	Template.PUT("/template-update", controllers.UpdateTemplate)

	return e
}
