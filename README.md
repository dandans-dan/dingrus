# dingrus
钉钉机器人的 [logrus](https://github.com/sirupsen/logrus) Hook.

## Useage
``` go
func main() {
	dh, _ := NewDingHook("https://oapi.dingtalk.com/robot/send?access_token=xxxxxx", nil)
	logrus.AddHook(dh)
	logrus.WithFields(logrus.Fields{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": "dingrus 测试",
			"text": "# Hello World",
		},
	}).Warn()
}
```
