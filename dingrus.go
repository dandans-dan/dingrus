// logrus 的 Hook 插件, 把日志通过钉钉机器人发送到群组
package dingrus

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/sirupsen/logrus"
)

type DingResponse struct {
	ErrCode int64  `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type DingHook struct {
	Webhook *url.URL // 钉钉机器人的 Webhook URL
	client  *http.Client
}

func NewDingHook(webhook string, client *http.Client) (*DingHook, error) {
	wh, err := url.Parse(webhook)
	if err != nil {
		return nil, errors.New("Parse webhook to url.URL error: " + err.Error())
	}

	dh := &DingHook{Webhook: wh}
	if client != nil {
		dh.client = client
	} else {
		dh.client = &http.Client{}
	}

	return dh, err
}

func (dh *DingHook) Levels() []logrus.Level {
	return []logrus.Level{
		logrus.WarnLevel,
		logrus.ErrorLevel,
		logrus.FatalLevel,
		logrus.PanicLevel,
	}
}

func (dh *DingHook) Fire(entry *logrus.Entry) error {
	b, err := json.Marshal(entry.Data)
	if err != nil {
		return errors.New("Marshal Fields to JSON error: " + err.Error())
	}

	body := ioutil.NopCloser(bytes.NewBuffer(b))
	request := &http.Request{
		Method:     "POST",
		URL:        dh.Webhook,
		Proto:      "HTTP/1.1",
		ProtoMajor: 1,
		ProtoMinor: 1,
		Header:     make(http.Header),
		Body:       body,
		Host:       dh.Webhook.Host,
	}
	request.Header.Set("Content-Type", "application/json; charset=utf-8")

	response, err := dh.client.Do(request)
	if err != nil {
		return errors.New("Send to DingTalk error: " + err.Error())
	}
	defer response.Body.Close()

	rb, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.New("Read DingTalk response error: " + err.Error())
	}

	dr := &DingResponse{}
	err = json.Unmarshal(rb, dr)
	if err != nil {
		return errors.New("Unmarshal DingTalk response to JSON error: " + err.Error())
	}

	if dr.ErrCode != 0 {
		return errors.New("DingTalk return error message: " + dr.ErrMsg)
	}

	return nil
}
