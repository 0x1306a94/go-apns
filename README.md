# go-apns

# APNS
- 相关链接
    * [官方文档 旧 APNS](https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/BinaryProviderAPI.html#//apple_ref/doc/uid/TP40008194-CH13-SW12)
    * [官方文档 新 HTTP/2 APNS](https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CommunicatingwithAPNs.html#//apple_ref/doc/uid/TP40008194-CH11-SW1)
    * [浅谈基于HTTP2推送消息到APNs](http://www.linkedkeeper.com/detail/blog.action?bid=167)
    * [基于HTTP/2与Token的APNs新协议](http://www.skyfox.org/apple-push-with-auth-key-token.html)
    * [苹果服务器是如何承载全球移动设备Push请求的](https://www.zhihu.com/question/33181208?sort=created)

#### Example
* First installation

```go
go get github.com/king129/go-apns
```

* Using p12 certificate

```go
import (
   "github.com/king129/go-apns"
   "github.com/king129/go-apns/cer"
   "fmt"
)
    
cer, err := cert.FromP12File("xxxx", "xxx")
if err != nil {
	fmt.Println(err)
	return
}

client := apns.NewClientWithCer(cer)

m := &apns.Message{
	Topic: "me.kinghub.apns.demo",
	DeviceToken: "xxxxxxxxxxxxxxxxx",
	Payload: &apns.MessagePayload{
		Aps: &apns.MessageAps{
			Alert: &apns.MessageAlert{
				Title: "test",
				Body: "xxxxxx",
			},
			Badeg: 100,
			Sound: apns.MessageAPNSSoundDefault,
		},
	},
}

client.Push(m)

```
* Using authKey

```go
import (
	"github.com/king129/go-apns"
	"github.com/king129/go-apns/token"
	"fmt"
)

authKey, err := token.AuthKeyFromFile("xxxx")
if err != nil {
	fmt.Println(err)
	return
}
	
token := &token.Token{
	AuthKey:authKey,
	TeamID: "xxxx",
	KeyID: "xxxx",
}
	
client := apns.NewClientWithToken(token)
	
m := &apns.Message{
	Topic: "me.kinghub.apns.demo",
	DeviceToken: "xxxxxxxxxxxxxxxxx",
	Payload: &apns.MessagePayload{
		Aps: &apns.MessageAps{
			Alert: &apns.MessageAlert{
				Title: "test",
				Body: "xxxxxx",
			},
			Badeg: 100,
			Sound: apns.MessageAPNSSoundDefault,
		},
	},
}
client.Push(m)

```

![](https://ws1.sinaimg.cn/large/006tKfTcgy1frh1uz69xqj31kw0zkb29.jpg)
![](https://ws1.sinaimg.cn/large/006tKfTcgy1frh1vlf6e2j30ku112h4s.jpg)
