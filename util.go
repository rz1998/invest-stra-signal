package signal

import "fmt"
import "github.com/rz1998/invest-stra-signal/types/signalStra"

// FromSignal2CSV 信号转换为csv
func FromSignal2CSV(signals []*signalStra.SSignalStra) string {
	var contents string
	// 生成标题行
	signal := &signalStra.SSignalStra{}
	contents = signal.Title() + "\n"
	// 生成内容行
	if signals != nil && len(signals) > 0 {
		for _, signal := range signals {
			contents += fmt.Sprintf("%s\n", signal)
		}
	}
	return contents
}
