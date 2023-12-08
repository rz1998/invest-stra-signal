/*
 *           佛曰:
 *                   写字楼里写字间，写字间里程序员；
 *                   程序人员写程序，又拿程序换酒钱。
 *                   酒醒只在网上坐，酒醉还来网下眠；
 *                   酒醉酒醒日复日，网上网下年复年。
 *                   但愿老死电脑间，不愿鞠躬老板前；
 *                   奔驰宝马贵者趣，公交自行程序员。
 *                   别人笑我忒疯癫，我笑自己命太贱；
 *                   不见满街漂亮妹，哪个归得程序员？
 *
 * @Author: rz1998 rz1998@126.com
 * @Date: 2023-08-08 11:17:38
 * @LastEditors: rz1998 rz1998@126.com
 * @LastEditTime: 2023-12-08 10:33:08
 * @FilePath: /signal/types/signalConfig/config.go
 * @Description:
 *
 */

package signalConfig

type ConfMqtt struct {
	Broker string
	Port   int
	Usr    string
	Psw    string
}

type ConfApiSignal struct {
	PMqtt  ConfMqtt
	Topics []string `json:",optional"`
}
