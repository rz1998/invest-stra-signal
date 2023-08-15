package signalStraMqtt

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/rz1998/invest-stra-signal/types/signalStra"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ApiSignalStraMqtt struct {
	Config       ConfApiSignal
	FuncOnSignal func(topic string, signals []*signalStra.SSignalStra)
	clientMqtt   *mqtt.Client
}

func (api *ApiSignalStraMqtt) Start() {
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

func (api *ApiSignalStraMqtt) Stop() {
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

func (api *ApiSignalStraMqtt) SubSignal(topics []string) {
	logx.Info("clientMqtt", "SubSignal", topics)
	filters := make(map[string]byte)
	for _, topic := range topics {
		filters[topic] = 2
	}
	token := (*api.clientMqtt).SubscribeMultiple(filters, func(client mqtt.Client, message mqtt.Message) {
		// 解析signals
		var signals []*signalStra.SSignalStra
		err := json.Unmarshal(message.Payload(), &signals)
		if err != nil {
			logx.Error("handlerSignals parse json error", err)
		}
		if signals == nil || len(signals) == 0 {
			fmt.Println("handlerSignals stopped by no signals")
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

func (api *ApiSignalStraMqtt) OnSignal(topic string, signals []*signalStra.SSignalStra) {
	api.FuncOnSignal(topic, signals)
}
