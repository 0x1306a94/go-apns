# go-apns

# APNS
- 相关链接
    * [官方文档 旧 APNS](https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/BinaryProviderAPI.html#//apple_ref/doc/uid/TP40008194-CH13-SW12)
    * [官方文档 新 HTTP/2 APNS](https://developer.apple.com/library/content/documentation/NetworkingInternet/Conceptual/RemoteNotificationsPG/CommunicatingwithAPNs.html#//apple_ref/doc/uid/TP40008194-CH11-SW1)
	* [官方文档](https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server)
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
    
    cer, err := cert.FromP12File(p12FilePath, p12Password)
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

```
* Using authKey

```go
    import (
        "github.com/king129/go-apns"
        "github.com/king129/go-apns/token"
        "fmt"
    )

    token, err := token.NewToken(authKeyPath, teamID, keyID)
	if err != nil {
		t.Error(err)
		return
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

```

![](https://ws1.sinaimg.cn/large/006tKfTcgy1frh1uz69xqj31kw0zkb29.jpg)
![](https://ws1.sinaimg.cn/large/006tKfTcgy1frh1vlf6e2j30ku112h4s.jpg)
