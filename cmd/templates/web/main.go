package main

import (
	"flag"
	"fmt"
	"net"
	"xframe/cmd"
	"xframe/cmd/templates/web/handler"
	"xframe/config"
	"xframe/log"
	"xframe/metric"
	"xframe/server"
)

var (
	confFile = flag.String("c", "", "configuration file, json format")
)

func main() {
	cmd.ParseCommand()
	cmd.DumpCommand()
	option, err := cmd.GetCommand("c")
	if err != nil {
		log.ERROR(err)
		panic(err)
	}
	if err := config.LoadConfigFromFile(option); err != nil {
		log.ERROR(err)
		panic(err)
	}
	config.DumpConfigContent()
	pprofAddr, _ := config.GetConfigByKey("server.pprofAddr")
	pprofPort, _ := config.GetConfigByKey("server.pprofPort")

	go func() {
		pprof := fmt.Sprintf("%s:%d", pprofAddr, pprofPort)
		log.INFOF("start to initialize metric at %s", pprof)
		metric.InitPprof(pprof)
	}()

	prometheusNet, _ := config.GetConfigByKey("metric.prometheusNetType")
	prometheusAddr, _ := config.GetConfigByKey("metric.prometheusAddr")
	if l, err := net.Listen(prometheusNet.(string), prometheusAddr.(string)); err != nil {
		panic(err)
	} else {
		metric.ServePrometheus(l)
	}

	addr, _ := config.GetConfigByKey("server.ip")
	port, _ := config.GetConfigByKey("server.port")

	mux := handler.RegisterHandler()
	if err := server.RunHTTPMux(addr.(string), int(port.(float64)), mux); err != nil {
		log.ERROR(err)
		panic(err)
	}
}
