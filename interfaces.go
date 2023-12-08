/*
 * @Author: rz1998 rz1998@126.com
 * @Date: 2023-08-03 10:45:38
 * @LastEditors: rz1998 rz1998@126.com
 * @LastEditTime: 2023-12-08 09:35:01
 * @FilePath: /signal/interfaces.go
 * @Description:
 *
 */
package signal

type ApiSignalClient[S any] interface {
	Start()
	Stop()
	SubSignal(topics []string)
	OnSignal(topic string, signals []*S)
}

type ApiSignalServer[S any] interface {
	Start()
	Stop()
	PubSignal(topic string, signals []*S)
}
