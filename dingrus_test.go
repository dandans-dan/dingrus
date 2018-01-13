package dingrus

import (
	"flag"
	"testing"

	"github.com/sirupsen/logrus"
)

var test_webhook = flag.String("webhook", "", "DingTalk Webhook")

func TestDingrusHook(t *testing.T) {
	dh, err := NewDingHook(*test_webhook, nil)
	t.Log(*test_webhook, err)
	if err != nil {
		t.Fatal(err)
	}

	logrus.AddHook(dh)

	logrus.WithFields(logrus.Fields{
		"msgtype": "markdown",
		"markdown": map[string]string{
			"title": "苍老师结婚啦!",
			"text":  "#### 苍老师结婚啦~~~  \n> 只是发个测试  \n> ![screenshot](https://gss3.bdstatic.com/-Po3dSag_xI4khGkpoWK1HF6hhy/baike/c0%3Dbaike92%2C5%2C5%2C92%2C30/sign=b17bc49a71f40ad101e9cfb136457aba/e4dde71190ef76c61cef08899716fdfaaf516757.jpg)\n> *测试结束*",
		},
	}).Error()
}
