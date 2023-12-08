package signalmqtt

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/rz1998/invest-stra-signal/types/signalConfig"
	"github.com/zeromicro/go-zero/core/logx"
)

type ApiSignalClientMqtt[S any] struct {
	Config       signalConfig.ConfApiSignal
	FuncOnSignal func(topic string, signals []*S)
	clientMqtt   *mqtt.Client
}

func (api *ApiSignalClientMqtt[S]) Start() {
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
		api.SubSignal(api.Config.Topics)
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

func (api *ApiSignalClientMqtt[S]) Stop() {
	if api.clientMqtt != nil {
		if token := (*api.clientMqtt).Unsubscribe(api.Config.Topics...); token.Wait() && token.Error() != nil {
			logx.Error("clientMqtt error disconnecting", token.Error())
			time.Sleep(2 * time.Second)
			api.Stop()
		}
		(*api.clientMqtt).Disconnect(250)
		api.clientMqtt = nil
	}
}

func (api *ApiSignalClientMqtt[S]) SubSignal(topics []string) {
	logx.Info("clientMqtt", "SubSignal", topics)
	filters := make(map[string]byte)
	for _, topic := range topics {
		filters[topic] = 2
	}
	token := (*api.clientMqtt).SubscribeMultiple(filters, func(client mqtt.Client, message mqtt.Message) {
		// 解析signals
		var signals []*S
		err := json.Unmarshal(message.Payload(), &signals)
		if err != nil {
			logx.Error("handlerSignals parse json error", err)
		}
		if len(signals) == 0 {
			logx.Error("handlerSignals stopped by no signals")
			return
		}
		strSignal := ""
		for _, signal := range signals {
			strSignal += fmt.Sprintf("%+v, ", *signal)
		}
		logx.Info("clientMqtt receiving", strSignal)
		api.OnSignal(message.Topic(), signals)
	})
	token.Wait()
}

func (api *ApiSignalClientMqtt[S]) OnSignal(topic string, signals []*S) {
	api.FuncOnSignal(topic, signals)
}
