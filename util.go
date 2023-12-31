/*
 * @Author: rz1998 rz1998@126.com
 * @Date: 2023-08-03 12:38:12
 * @LastEditors: rz1998 rz1998@126.com
 * @LastEditTime: 2023-12-11 13:17:55
 * @FilePath: /signal/util.go
 * @Description:
 *
 */
package signal

import (
	"fmt"

	"github.com/rz1998/invest-stra-signal/types/signalStra"
)

// FromSignal2CSV 信号转换为csv
func FromSignal2CSV(signals []*signalStra.SSignalStra) string {
	var contents string
	// 生成标题行
	signal := &signalStra.SSignalStra{}
	contents = signal.Title() + "\n"
	// 生成内容行
	if len(signals) > 0 {
		for _, signal := range signals {
			contents += fmt.Sprintf("%s\n", signal)
		}
	}
	return contents
}
