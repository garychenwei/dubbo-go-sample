package main

import (
	"fmt"
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iotcenter/golange/api/router"
	"github.com/iotcenter/golange/api/gobal"
	"dubbo.apache.org/dubbo-go/v3/config"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "dubbo.apache.org/dubbo-go/v3/filter/filter_impl"
	_ "github.com/iotcenter/golange/api/filter"
)
// ErrorMsg msg
type ErrorMsg struct {
	Code int
	Message string
}



func initDubbo() {
	config.SetConsumerService(gobal.MaterialServiceClientImpl)
	// path := "/Users/chenweijiang/Documents/workspaces/msquare/iotcenter/golange/api/conf/dubbogo.yml"
	logger.Info("start to test dubbo")
	if err := config.Load(); err != nil {
		panic(err)
	}
}

func main() {
	initDubbo()
	//init the fiber app and router
	errorHandler := func(ctx *fiber.Ctx, err error) error {
		fmt.Println(err)
		
		return ctx.Status(500).JSON(&ErrorMsg{ Code:-1, Message: err.Error()})
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	app.Use(recover.New())
  router.SetRouter(app)
	log.Fatal(app.Listen(":3000"))
	// init the dubbo
}