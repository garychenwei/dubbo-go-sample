package router

import "github.com/gofiber/fiber/v2"
import "github.com/iotcenter/golange/api/controller/erp"

func SetRouter(app *fiber.App) {
	SetErpRouter(app)
}

func SetErpRouter(app *fiber.App) {
	erpGroup := app.Group("/erp")
	meterialGroup := erpGroup.Group("/meterial")
	meterialGroup.Post("/get", erp.MeterialControllerImpl.GetMeterial)
}