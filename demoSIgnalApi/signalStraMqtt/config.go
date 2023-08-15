package signalStraMqtt

type ConfMqtt struct {
	Broker string
	Port   int
	Usr    string
	Psw    string
}

type ConfApiSignal struct {
	PMqtt  ConfMqtt
	Topics []string
}
