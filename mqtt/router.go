package mqtt

import (
	"edgeForward/mqtt/collector_data_report"
	"github.com/stevenyao001/edgeCommon/mqtt"
)

var Subscribes = map[string][]mqtt.SubscribeOpts{
	"rootcloud": {
		//新采集数据
		{
			Topic:    "$ROOTCLOUD/thing/realtimedata/+",
			Qos:      0,
			Callback: collector_data_report.CollectorDataReport,
		},
	},
}
