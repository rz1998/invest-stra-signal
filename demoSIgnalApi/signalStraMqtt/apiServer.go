/*
 *                   ___====-_  _-====___
 *             _--^^^#####//      \\#####^^^--_
 *          _-^##########// (    ) \\##########^-_
 *         -############//  |\^^/|  \\############-
 *       _/############//   (@::@)   \############\_
 *      /#############((     \\//     ))#############\
 *     -###############\\    (oo)    //###############-
 *    -#################\\  / VV \  //#################-
 *   -###################\\/      \//###################-
 *  _#/|##########/\######(   /\   )######/\##########|\#_
 *  |/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
 *  `  |/  V  V  `   V  \#\| |  | |/#/  V   '  V  V  \|  '
 *     `   `  `      `   / | |  | | \   '      '  '   '
 *                      (  | |  | |  )
 *                     __\ | |  | | /__
 *                    (vvv(VVV)(VVV)vvv)
 *
 *      ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~
 *
 *                神兽保佑            永无BUG
 *
 * @Author: rz1998 rz1998@126.com
 * @Date: 2023-09-13 21:01:08
 * @LastEditors: rz1998 rz1998@126.com
 * @LastEditTime: 2023-09-13 21:06:09
 * @FilePath: /signal/demoSIgnalApi/signalStraMqtt/apiServer.go
 * @Description:
 *
 */

package signalStraMqtt

import (
	"encoding/json"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/rz1998/invest-stra-signal/types/signalStra"
	"github.com/zeromicro/go-zero/core/logx"
)

type ApiSignalServerStraMqtt struct {
	Config     ConfApiSignal
	clientMqtt *mqtt.Client
}

func (api *ApiSignalServerStraMqtt) Start() {
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

func (api *ApiSignalServerStraMqtt) Stop() {
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

func (api *ApiSignalServerStraMqtt) PubSignal(topic string, signals []*signalStra.SSignalStra) {
	if api.clientMqtt != nil {
		jsonData, _ := json.Marshal(&signals)
		(*api.clientMqtt).Publish(topic, 2, false, string(jsonData))
	}
}
