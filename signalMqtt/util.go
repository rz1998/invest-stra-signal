/*
 * @Author: rz1998 rz1998@126.com
 * @Date: 2023-12-11 13:18:02
 * @LastEditors: rz1998 rz1998@126.com
 * @LastEditTime: 2023-12-11 13:18:04
 * @FilePath: /signal/signalMqtt/util.go
 * @Description:
 *
 */
package signalMqtt

import (
	"fmt"
	"strings"
)

const TopicSignalDataUpdated = "signal/dataUpdated/"

// 唯一识别转换为topic
func FromUC2Topic(ucData string) string {
	return strings.ReplaceAll(ucData, ".", "/")
}

// 根据唯一识别，生成数据更新信号的topic
func TopicSignalDataUpdate(ucData string) string {
	return fmt.Sprintf("%s%s", TopicSignalDataUpdated, strings.ReplaceAll(ucData, ".", "/"))
}
