package apns

import (
	"fmt"
	"os"
	"testing"

	"github.com/king129/go-apns/cer"
	"github.com/king129/go-apns/token"
)

func TestClient_UsignCer(t *testing.T) {

	p12Path := os.Getenv("P12PATH")
	p12Passwd := os.Getenv("P12PASSWD")

	cer, err := cert.FromP12File(p12Path, p12Passwd)
	if err != nil {
		t.Error(err)
	}

	client := NewClientWithCer(cer)

	payload := NewPayload()
	payload.Body("测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试")
	payload.Title("标题")
	payload.SubTitle("子标题")
	payload.Badge(100)
	payload.Sound(MessageAPNSSoundDefault)
	m := &Message{
		Topic:       "com.taihe.test.moblink",
		DeviceToken: "ac8a3585166d60abd44867195b8f3edb63f4a2d1b2bb77c896913b7dd68716b8",
		Payload:     payload,
	}

	resp, err := client.Push(m)
	if err != nil {
		t.Error(err)
	}
	if !resp.Success() {
		t.Error(resp.Reason)
	} else {
		fmt.Println("发送成功 apns-id:", resp.ApnsID)
	}

}

func TestClient_UsignToken(t *testing.T) {

	authKeyPath := os.Getenv("AUTHKEY")

	authKey, err := token.AuthKeyFromFile(authKeyPath)
	if err != nil {
		fmt.Println(err)
		return
	}

	token := &token.Token{
		AuthKey: authKey,
		TeamID:  "xxxx",
		KeyID:   "xxxx",
	}

	client := NewClientWithToken(token)

	payload := NewPayload()
	payload.Body("测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试测试")
	payload.Title("标题")
	payload.SubTitle("子标题")
	payload.Badge(100)
	payload.Sound(MessageAPNSSoundDefault)
	m := &Message{
		Topic:       "com.taihe.test.moblink",
		DeviceToken: "ac8a3585166d60abd44867195b8f3edb63f4a2d1b2bb77c896913b7dd68716b8",
		Payload:     payload,
	}

	resp, err := client.Push(m)
	if err != nil {
		t.Error(err)
	}
	if !resp.Success() {
		t.Error(resp.Reason)
	} else {
		fmt.Println("发送成功 apns-id:", resp.ApnsID)
	}
}
