package main

import (
	"edgeForward/dataSource"
	"edgeForward/global"
	mqtt2 "edgeForward/mqtt"
	"github.com/stevenyao001/edgeCommon"
	"github.com/stevenyao001/edgeCommon/mqtt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	//core.Run()
	edge := edgeCommon.New()

	filePath, _ := os.Getwd()
	edge.RegisterConfig(filePath+"/conf/local.yaml", &global.Conf)

	edge.RegisterLogger(global.Conf.Log.MainPath)

	mqttConfs := make([]mqtt.Conf, 0)
	for k := range global.Conf.Mqtt {
		mqttConfs = append(mqttConfs, mqtt.Conf{
			InsName:  global.Conf.Mqtt[k].InsName,
			ClientId: global.Conf.Mqtt[k].ClientId,
			Username: global.Conf.Mqtt[k].Username,
			Password: global.Conf.Mqtt[k].Password,
			Addr:     global.Conf.Mqtt[k].Addr,
			Port:     global.Conf.Mqtt[k].Port,
		})
	}
	edge.RegisterMqtt(mqttConfs, mqtt2.Subscribes)

	ds, _ := dataSource.DataSourceNew()
	ds.Run()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	_ = <-quit

}
