/*
 * @Author: rz1998 rz1998@126.com
 * @Date: 2023-12-08 10:34:48
 * @LastEditors: rz1998 rz1998@126.com
 * @LastEditTime: 2023-12-11 11:28:34
 * @FilePath: /signal/signalMqtt/server.go
 * @Description:
 *
 */
package signalMqtt

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/rz1998/invest-stra-signal/types/signalConfig"
	"github.com/zeromicro/go-zero/core/logx"
)

type ApiSignalServerMqtt[S any] struct {
	Config     signalConfig.ConfApiSignal
	clientMqtt *mqtt.Client
}

func (api *ApiSignalServerMqtt[S]) Start() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", api.Config.PMqtt.Broker, api.Config.PMqtt.Port))
	// uuid
	u1, err := uuid.NewUUID()
	if err != nil {
		logx.Error("clientMqtt uuid error", err)
	}
	opts.SetClientID(u1.String())
	opts.SetUsername(api.Config.PMqtt.Usr)
	opts.SetPassword(api.Config.PMqtt.Psw)
	opts.AutoReconnect = true
	opts.OnReconnecting = func(client mqtt.Client, options *mqtt.ClientOptions) {
		logx.Info("clientMqtt reconnecting...")
	}
	opts.OnConnect = func(client mqtt.Client) {
		logx.Info("clientMqtt connect succeed...")
		api.clientMqtt = &client
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		logx.Error("clientMqtt connection error", err)
	}
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		logx.Error("clientMqtt error connecting", token.Error())
		time.Sleep(2 * time.Second)
		api.Start()
	} else {
		api.clientMqtt = &client
	}
}

func (api *ApiSignalServerMqtt[S]) Stop() {
	if api.clientMqtt != nil {
		(*api.clientMqtt).Disconnect(250)
		api.clientMqtt = nil
	}
}

func (api *ApiSignalServerMqtt[S]) PubSignal(topic string, signals []*S) {
	if api.clientMqtt != nil {
		jsonData, _ := json.Marshal(&signals)
		token := (*api.clientMqtt).Publish(topic, 2, false, string(jsonData))
		token.Wait()
	}
}
