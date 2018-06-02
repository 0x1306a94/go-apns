package apns

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func TestNewPayload(t *testing.T) {

	payload := NewPayload()
	payload.Body("测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试")
	payload.Title("标题")
	payload.SubTitle("子标题")
	payload.Badge(100)
	payload.Sound(MessageAPNSSoundDefault)
	payload.AddCustom("time", time.Now())
	fmt.Println(payload)
	data, err := json.Marshal(payload)
	if err != nil {
		t.Error(err)
	} else {
		fmt.Println(string(data))
	}

	var p2 Payload
	if err = json.Unmarshal(data, &p2); err != nil {
		t.Error(err)
	} else {
		fmt.Println(&p2)
	}
}
