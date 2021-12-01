package main

import (
	"fmt"
	"log"

	"dubbo.apache.org/dubbo-go/v3/common/logger"
	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/iotcenter/golange/api/gobal"
	"github.com/iotcenter/golange/api/router"
	"github.com/opentracing/opentracing-go"
	zipkinot "github.com/openzipkin-contrib/zipkin-go-opentracing"
	"github.com/openzipkin/zipkin-go"
	zipkinhttp "github.com/openzipkin/zipkin-go/reporter/http"
)

type ErrorMsg struct {
	Code    int
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

func initZipkin() {
	reporter := zipkinhttp.NewReporter("http://localhost:9411/api/v2/spans")
	endpoint, err := zipkin.NewEndpoint("dobbugoZipkinTracingService", "myservice.mydomain.com:80")
	if err != nil {
		logger.Errorf("unable to create local endpoint: %+v\n", err)
	}
	nativeTracer, err := zipkin.NewTracer(reporter, zipkin.WithLocalEndpoint(endpoint))
	if err != nil {
		logger.Errorf("unable to create tracer: %+v\n", err)
	}
	tracer := zipkinot.Wrap(nativeTracer)
	opentracing.SetGlobalTracer(tracer)
}

func main() {
	initZipkin()
	initDubbo()

	//init the fiber app and router
	errorHandler := func(ctx *fiber.Ctx, err error) error {
		fmt.Println(err)

		return ctx.Status(500).JSON(&ErrorMsg{Code: -1, Message: err.Error()})
	}
	app := fiber.New(fiber.Config{
		ErrorHandler: errorHandler,
	})
	app.Use(recover.New())
	router.SetRouter(app)
	log.Fatal(app.Listen(":3000"))
	// init the dubbo
}
