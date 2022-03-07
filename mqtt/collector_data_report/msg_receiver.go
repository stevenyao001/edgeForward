package collector_data_report

import (
	"fmt"
	mqtt2 "github.com/eclipse/paho.mqtt.golang"
	"github.com/stevenyao001/edgeCommon/mqtt"
)

func MsgNew() mqtt.Msg {
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

//消息接收者
var CollectorDataReport = func(client mqtt2.Client, msg mqtt2.Message) {
	fmt.Println("forward receive topic: ", msg.Topic())
	fmt.Println("forward receive msg: ", string(msg.Payload()))

	mqttClient := mqtt.GetClient("rootcloud")

	msgSend := MsgNew()
	msgSend.Content["testRoot"] = string(msg.Payload())

	_, _ = mqttClient.Publish("$ROOTEDGE/thing/upload", msgSend, 2, false)

	//msgEntity := mqtt2.MsgPool.Get().(mqtt2.Msg)
	//defer mqtt2.MsgPool.Put(msgEntity)
	//
	//err := json.Unmarshal(msg.Payload(), &msgEntity)
	//if err != nil {
	//	logger.ErrorLog("MsgReceiver-ReceiveMsg", "消息解析失败", "", err)
	//	return
	//}
	//if msgEntity.DeviceId == "" {
	//	logger.ErrorLog("MsgReceiver-ReceiveMsg", "设备id不能为空", "", err)
	//	return
	//}
	//
	//if msgEntity.Cmd == mqtt2.CollectDeviceRegister {
	//	msgHandlerM.newMsgHandler(msgEntity.DeviceId)
	//	return
	//}
	//if msgEntity.Cmd == mqtt2.CollectDeviceDel {
	//	msgHandlerM.delMsgHandler(msgEntity.DeviceId)
	//	return
	//}
	//
	//go msgHandlerM.msgPutQueue(msgEntity)
}
