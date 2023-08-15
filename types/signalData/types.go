package signalData

// SSignalDataUpdated 数据更新信号
type SSignalDataUpdated struct {
	UcData    string `json:"ucData"`
	Timestamp int64  `json:"timestamp"`
	Num       int64  `json:"num"`
}
