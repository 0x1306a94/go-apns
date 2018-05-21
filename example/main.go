package main

import (
	"fmt"
	"os"

	"github.com/king129/go-apns"
	"github.com/king129/go-apns/cer"
	"github.com/king129/go-apns/token"
)

func main() {

	// Using p12 certificate
	UsingCer()
	// Using authKey
	UsingAuthKey()
}

func UsingCer() {

	p12Path := os.Getenv("P12PATH")
	p12Passwd := os.Getenv("P12PASSWD")

	cer, err := cert.FromP12File(p12Path, p12Passwd)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := apns.NewClientWithCer(cer)

	m := &apns.Message{
		Topic:       "me.kinghub.apns.demo",
		DeviceToken: "xxxxxxxxxxxxxxxxx",
		Payload: &apns.MessagePayload{
			Aps: &apns.MessageAps{
				Alert: &apns.MessageAlert{
					Title: "test",
					Body:  "xxxxxx",
				},
				Badge: 100,
				Sound: apns.MessageAPNSSoundDefault,
			},
		},
	}

	client.Push(m)
}

func UsingAuthKey() {

	authKeyPath := os.Getenv("AUTHKEY")
	teamID := os.Getenv("TEAMID")
	keyID := os.Getenv("KEYID")

	token, err := token.NewToken(authKeyPath, teamID, keyID)
	if err != nil {
		fmt.Println(err)
		return
	}

	client := apns.NewClientWithToken(token)

	m := &apns.Message{
		Topic:       "me.kinghub.apns.demo",
		DeviceToken: "xxxxxxxxxxxxxxxxx",
		Payload: &apns.MessagePayload{
			Aps: &apns.MessageAps{
				Alert: &apns.MessageAlert{
					Title: "test",
					Body:  "xxxxxx",
				},
				Badge: 100,
				Sound: apns.MessageAPNSSoundDefault,
			},
		},
	}
	client.Push(m)
}
