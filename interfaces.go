package signal

import "github.com/rz1998/invest-stra-signal/types/signalStra"

type ApiSignal interface {
	Start()
	Stop()
	SubSignal(topics []string)
	OnSignal(topic string, signals []*signalStra.SSignalStra)
}
