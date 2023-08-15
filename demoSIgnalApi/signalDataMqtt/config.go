package signalDataMqtt

type ConfMqtt struct {
	Broker string
	Port   int
	Usr    string
	Psw    string
}

type ConfApiSignal struct {
	Mqtt   ConfMqtt
	Topics []string
}
