package main


import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

import (
	"dubbo.apache.org/dubbo-go/v3/config"
	"github.com/iotcenter/golange/erp/app"
	"dubbo.apache.org/dubbo-go/v3/common/logger"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	_ "github.com/iotcenter/golange/erp/filter"
)

var (
	survivalTimeout = int(3e9)
)

func main() {
	logger.SetLoggerLevel("debug")
	initBaseConfig()
	// initConfigCenter()
	// select {}
	initSignal()
}

func initBaseConfig() {
	config.SetProviderService(new(app.MeterialProvider))
	// path := "/Users/chenweijiang/Documents/workspaces/msquare/iotcenter/golange/erp/conf/dubbogo.yml"
	if err := config.Load(); err != nil {
		logger.Error("init error ", err)
		panic(err)
	}
}

const configCenterZKServerConfig = `# set in config center, group is 'dubbogo', dataid is 'dubbo-go-samples-configcenter-zookeeper-server', namespace is default
dubbo:
  registries:
    demoZK:
      protocol: zookeeper
      address: 127.0.0.1:2181
  protocols:
    triple:
      name: tri
      port: 20000
  provider:
    services:
			MeterialProvider:
        interface: com.apache.dubbo.sample.basic.IGreeter # must be compatible with grpc or dubbo-java`

func initConfigCenter() {
	dynamicConfig, err := config.NewConfigCenterConfigBuilder().
		SetProtocol("zookeeper").
		SetAddress("127.0.0.1:2181").
		Build().
		CreateDynamicConfiguration()
		if err != nil {
			panic(err)
		}
		if err := dynamicConfig.PublishConfig("dubbo-go-samples-configcenter-zookeeper-server", "dubbogo", configCenterZKServerConfig); err != nil {
			panic(err)
		}
		time.Sleep(time.Second * 10)

		config.SetProviderService(&app.MeterialProvider{})
		rootConfig := config.NewRootConfigBuilder().
		SetConfigCenter(config.NewConfigCenterConfigBuilder().
			SetProtocol("zookeeper").SetAddress("127.0.0.1:2181").
			SetDataID("dubbo-go-samples-configcenter-zookeeper-server").
			SetGroup("dubbogo").
			Build()).
		Build()

	if err := config.Load(config.WithRootConfig(rootConfig)); err != nil {
		panic(err)
	}
}


func initSignal() {
	signals := make(chan os.Signal, 1)
	// It is not possible to block SIGKILL or syscall.SIGSTOP
	signal.Notify(signals, os.Interrupt, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)
	for {
		sig := <-signals
		logger.Infof("get signal %s", sig.String())
		switch sig {
		case syscall.SIGHUP:
			// reload()
		default:
			time.AfterFunc(time.Duration(survivalTimeout), func() {
				logger.Warnf("app exit now by force...")
				os.Exit(1)
			})

			// The program exits normally or timeout forcibly exits.
			fmt.Println("provider app exit now...")
			return
		}
	}
}