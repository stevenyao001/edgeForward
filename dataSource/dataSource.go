package dataSource

import (
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/stevenyao001/edgeCommon/mqtt"
	"math/rand"
	"time"
)

/* define new struct dataSource */
type DataSource struct {
	GlobalCount int64
}

/* define new struct collectData for marshal */

type DD struct {
	Properties *CollectData `json:"properties"`
	Ts         int          `json:"ts"`
}
type CollectData struct {
	MegnetStatus bool `json:"megnet_status"`
	//MegnetStatusCnt int64 `json:""`
	//Status          int32 `json:""`
	Ia int `json:"ia"`
	//Threadhold      int32 `json:""`
	Ep int64 `json:"ep"`
	//InstantEp       int32 `json:""`
}

func (d *DataSource) Run() {
	/**
	 * @Descripton: your method
	 * @Params:
	 * @Return:
	 * @Date: 2022/3/4 下午2:26
	 */
	go d.createData()

	return
}

func (d *DataSource) createData() {
	/**
	 * @Description: your method
	 * @Params:
	 * @Return:
	 * @Date: 2022/3/4 下午2:28
	 */
	cData := d.initData()

	mqttClient := mqtt.GetClient("rootcloud")

	msgSend := d.msgNew()
	msgSend.DeviceId = "18"
	msgSend.Cmd = 101001
	_, _ = mqttClient.Publish("$ROOTEDGE/thing/model/18", msgSend, 2, false)
	time.Sleep(time.Second)
	for true {
		if rand.Intn(2) == 1 {
			cData.MegnetStatus = true
		} else {
			cData.MegnetStatus = false
		}
		cData.Ia = rand.Intn(60)
		cData.Ep = d.GlobalCount
		d.GlobalCount += 1

		dd := DD{
			Properties: cData,
			Ts:         123456,
		}
		jsons, _ := jsoniter.Marshal(dd)
		fmt.Println("new data: ", string(jsons))
		mqttClient := mqtt.GetClient("rootcloud")

		msgSend := d.msgNew()
		msgSend.DeviceId = "18"

		json.Unmarshal(jsons, &msgSend.Content)

		_, _ = mqttClient.Publish("$ROOTEDGE/datasource/rawdata/18", msgSend, 2, false)
		time.Sleep(60 * time.Second)
	}
	return
}

func (d *DataSource) msgNew() mqtt.Msg {
	/**
	 * @Description: your method
	 * @Params:
	 * @Return:
	 * @Date: 2022/3/4 下午3:28
	 */
	return mqtt.Msg{
		MsgId:    0,
		TraceId:  "",
		DeviceId: "",
		Version:  "",
		Source:   0,
		Mold:     0,
		Cmd:      0,
		Content:  make(map[string]interface{}),
	}
}

func (d *DataSource) initData() (c *CollectData) {
	/**
	 * @Description: your method
	 * @Params:
	 * @Return:
	 * @Date: 2022/3/4 下午2:48
	 */
	return &CollectData{
		MegnetStatus: true,
		//MegnetStatusCnt: 0,
		Ia: 0,
		//Threadhold: 40,
		//Status: 0,
		Ep: 0,
		//InstantEp: 0,
	}
}

func (d *DataSource) Destroy() (e error) {
	/**
	 * @Description: your method
	 * @Params:
	 * @Return:
	 * @Date: 2022/3/4 下午2:29
	 */

	return nil
}

func DataSourceNew() (*DataSource, error) {
	/*
	 * @Description: new method for MethodNew
	 * @Params:
	 * @Return: new pointer for Struct
	 * @Date: 2022/3/4 下午2:22
	 */
	return &DataSource{
		GlobalCount: 0,
	}, nil
}
