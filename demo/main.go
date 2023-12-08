/*
 * @Author: rz1998 rz1998@126.com
 * @Date: 2023-12-08 10:25:48
 * @LastEditors: rz1998 rz1998@126.com
 * @LastEditTime: 2023-12-08 10:38:05
 * @FilePath: /signal/demo/main.go
 * @Description:
 *
 */
package main

import (
	"fmt"
	"time"

	signalmqtt "github.com/rz1998/invest-stra-signal/signalMqtt"
	"github.com/rz1998/invest-stra-signal/types/signalConfig"
	"github.com/rz1998/invest-stra-signal/types/signalStra"
)

func main() {
	client := signalmqtt.ApiSignalClientMqtt[signalStra.SSignalStra]{
		Config: signalConfig.ConfApiSignal{
			PMqtt: signalConfig.ConfMqtt{
				Broker: "101.132.144.250",
				Port:   1883,
				Usr:    "zhangyu",
				Psw:    "zhangyum123",
			},
			Topics: []string{"signal/stra/test"},
		},
		FuncOnSignal: func(topic string, signals []*signalStra.SSignalStra) {
			fmt.Printf("on signal %s\n", topic)
			for _, signal := range signals {
				fmt.Printf("signal %+v\n", signal)
			}
		},
	}
	client.Start()
	time.Sleep(5 * time.Second)
	server := signalmqtt.ApiSignalServerMqtt[signalStra.SSignalStra]{
		Config: signalConfig.ConfApiSignal{
			PMqtt: signalConfig.ConfMqtt{
				Broker: "101.132.144.250",
				Port:   1883,
				Usr:    "zhangyu",
				Psw:    "zhangyum123",
			},
			Topics: []string{"signal/stra/test"},
		}}
	server.Start()
	server.PubSignal("signal/stra/test", []*signalStra.SSignalStra{{
		NameStra: "test",
	}})
	select {}
}
