package main

import (
	"fmt"

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

	cer, err := cert.FromP12File("xxxx", "xxx")
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

	authKey, err := token.AuthKeyFromFile("xxxx")
	if err != nil {
		fmt.Println(err)
		return
	}

	token := &token.Token{
		AuthKey: authKey,
		TeamID:  "xxxx",
		KeyID:   "xxxx",
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
