package apns

/*
{
     "aps" : {
         "alert" : {   // string or dictionary
			"title"          :   "string",
			"subtitle"       :   "string",
            "body"           :   "string",
            "title-loc-key"  :   "string or null",
            "title-loc-args" :   "array of strings or null",
            "action-loc-key" :   "string or null",
            "loc-key"        :   "string",
            "loc-args"       :   "array of strings",
            "launch-image"   :   "string"
         },
          "badge"             :  number,
          "sound"             :  "string",
		  "content-available" :  number,
		  "mutable-content"   :  number,
          "category"          :  "string"
     },
}

aps：推送消息必须有的key
alert：推送消息包含此key值，系统就会根据用户的设置展示标准的推送信息
badge：在app图标上显示消息数量，缺少此key值，消息数量就不会改变，消除标记时把此key对应的value设置为0
sound：设置推送声音的key值，系统默认提示声音对应的value值为default
content-available：此key值设置为1，系统接收到推送消息时就会调用不同的回调方法，iOS7之后配置后台模式
category：UIMutableUserNotificationCategory's identifier 可操作通知类型的key值
title：简短描述此调推送消息的目的，适用系统iOS8.2之后版本
body：推送的内容
title-loc-key：功能类似title，附加功能是国际化，适用系统iOS8.2之后版本
title-loc-args：配合title-loc-key字段使用，适用系统iOS8.2之后版本
action-loc-key：可操作通知类型key值，不详细叙述
loc-key：参考title-loc-key
loc-args：参考title-loc-args
launch-image：点击推送消息或者移动事件滑块时，显示的图片。如果缺少此key值，会加载app默认的启动图片。

// 官方 json payload 定义
https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/creating_a_remote_notification_payload#2943363
*/
type MessagePayload struct {
	Aps    *MessageAps            `json:"aps"`
	Custom map[string]interface{} `json:"custom,omitempty"`
}

type MessageAps struct {
	Alert            *MessageAlert `json:"alert"`
	Badeg            int64         `json:"badeg,omitempty"`
	Sound            string        `json:"sound,omitempty"`
	ContentAvailable int           `json:"content-available,omitempty"`
	MutableContent   int           `json:"mutable-content,omitempty"`
	Category         string        `json:"category,omitempty"`
}

type MessageAlert struct {
	Title        string `json:"title"`
	SubTitle     string `json:"subtitle,omitempty"`
	Body         string `json:"body"`
	TitleLocKey  string `json:"title-loc-key,omitempty"`
	TitleLocArgs string `json:"title-loc-args,omitempty"`
	ActionLocKey string `json:"action-loc-key,omitempty"`
	LocKey       string `json:"loc-key,omitempty"`
	LocArgs      string `json:"loc-args,omitempty"`
	LaunchImage  string `json:"launch-image,omitempty"`
}
