package signalDataMqtt

import (
	"encoding/json"
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/rz1998/invest-stra-signal/types/signalData"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type ApiSignalDataMqtt struct {
	Config       ConfApiSignal
	FuncOnSignal func(topic string, signal *signalData.SSignalDataUpdated)
	clientMqtt   *mqtt.Client
}

func (api *ApiSignalDataMqtt) Start() {
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", api.Config.Mqtt.Broker, api.Config.Mqtt.Port))
	// uuid
	u1, err := uuid.NewUUID()
	if err != nil {
		logx.Error("clientMqtt uuid error", err)
	}
	opts.SetClientID(u1.String())
	opts.SetUsername(api.Config.Mqtt.Usr)
	opts.SetPassword(api.Config.Mqtt.Psw)
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

func (api *ApiSignalDataMqtt) Stop() {
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

func (api *ApiSignalDataMqtt) SubSignal(topics []string) {
	logx.Info("clientMqtt", "SubSignal", topics)
	if len(topics) == 0 {
		return
	}
	filters := make(map[string]byte)
	for _, topic := range topics {
		filters[topic] = 2
	}
	token := (*api.clientMqtt).SubscribeMultiple(filters, func(client mqtt.Client, message mqtt.Message) {
		// 解析signals
		var signal *signalData.SSignalDataUpdated
		err := json.Unmarshal(message.Payload(), &signal)
		if err != nil {
			logx.Error("handlerSignals parse json error", err)
		}
		if signal == nil {
			fmt.Println("handlerSignals stopped by no signals")
			return
		}
		strSignal := ""
		strSignal += fmt.Sprintf("%+v, ", *signal)
		logx.Info("clientMqtt receiving", strSignal)
		api.OnSignal(message.Topic(), signal)
	})
	token.Wait()
}

func (api *ApiSignalDataMqtt) OnSignal(topic string, signal *signalData.SSignalDataUpdated) {
	api.FuncOnSignal(topic, signal)
}

func (api *ApiSignalDataMqtt) PubSignal(topic string, signal *signalData.SSignalDataUpdated) {
	if signal == nil {
		return
	}
	jsonData, _ := json.Marshal(signal)
	token := (*api.clientMqtt).Publish(
		topic,
		2, false, string(jsonData))
	token.Wait()
}
