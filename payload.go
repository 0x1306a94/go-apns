package apns

import (
	"encoding/json"
)

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
type Payload struct {
	Aps    *aps                   `json:"aps"`
	Custom map[string]interface{} `json:"custom,omitempty"`
}

type aps struct {
	Alert            *alert      `json:"alert"`
	Badge            interface{} `json:"badge,omitempty"`
	Sound            string      `json:"sound,omitempty"`
	ThreadID         string      `json:"thread-id,omitempty"`
	ContentAvailable int         `json:"content-available,omitempty"`
	MutableContent   int         `json:"mutable-content,omitempty"`
	Category         string      `json:"category,omitempty"`
	URLArgs          []string    `json:"url-args,omitempty"`
}

type alert struct {
	Title           string `json:"title"`
	SubTitle        string `json:"subtitle,omitempty"`
	Body            string `json:"body,omitempty"`
	TitleLocKey     string `json:"title-loc-key,omitempty"`
	TitleLocArgs    string `json:"title-loc-args,omitempty"`
	SubTitleLocKey  string `json:"subtitle-loc-key,omitempty"`
	SubTitleLocArgs string `json:"subtitle-loc-args,omitempty"`
	ActionLocKey    string `json:"action-loc-key,omitempty"`
	LocKey          string `json:"loc-key,omitempty"`
	LocArgs         string `json:"loc-args,omitempty"`
	LaunchImage     string `json:"launch-image,omitempty"`
}

func NewPayload() *Payload {
	return &Payload{
		Aps: &aps{
			Alert: &alert{},
		},
	}
}
func (p *Payload) AddCustom(k string, v interface{}) *Payload {
	if p.Custom == nil {
		p.Custom = make(map[string]interface{})
	}
	p.Custom[k] = v
	return p
}

func (p *Payload) Body(v string) *Payload {
	p.aps().Alert.Body = v
	return p
}

func (p *Payload) Badge(number int64) *Payload {
	p.aps().Badge = number
	return p
}

func (p *Payload) ZeroBadge() *Payload {
	p.aps().Badge = 0
	return p
}

func (p *Payload) UnsetBadge() *Payload {
	p.aps().Badge = nil
	return p
}

func (p *Payload) Sound(v string) *Payload {
	p.aps().Sound = v
	return p
}

func (p *Payload) ThreadID(v string) *Payload {
	p.aps().ThreadID = v
	return p
}

func (p *Payload) ContentAvailable() *Payload {
	p.aps().ContentAvailable = 1
	return p
}

func (p *Payload) UnsetContentAvailable() *Payload {
	p.aps().ContentAvailable = 0
	return p
}

func (p *Payload) MutableContent() *Payload {
	p.aps().MutableContent = 1
	return p
}

func (p *Payload) UnsetMutableContent() *Payload {
	p.aps().MutableContent = 0
	return p
}

func (p *Payload) Category(v string) *Payload {
	p.aps().Category = v
	return p
}

func (p *Payload) UnsetCategory(v string) *Payload {
	p.aps().Category = ""
	return p
}

func (p *Payload) URLArgs(v []string) *Payload {
	p.aps().URLArgs = v
	return p
}

func (p *Payload) Title(v string) *Payload {
	p.aps().Alert.Title = v
	return p
}

func (p *Payload) SubTitle(v string) *Payload {
	p.aps().Alert.SubTitle = v
	return p
}
func (p *Payload) TitleLocKey(v string) *Payload {
	p.aps().Alert.TitleLocKey = v
	return p
}
func (p *Payload) TitleLocArgs(v string) *Payload {
	p.aps().Alert.TitleLocArgs = v
	return p
}
func (p *Payload) SubTitleLocKey(v string) *Payload {
	p.aps().Alert.SubTitleLocKey = v
	return p
}
func (p *Payload) SubTitleLocArgs(v string) *Payload {
	p.aps().Alert.SubTitleLocArgs = v
	return p
}
func (p *Payload) ActionLocKey(v string) *Payload {
	p.aps().Alert.ActionLocKey = v
	return p
}
func (p *Payload) LocKey(v string) *Payload {
	p.aps().Alert.LocKey = v
	return p
}
func (p *Payload) LocArgs(v string) *Payload {
	p.aps().Alert.LocArgs = v
	return p
}
func (p *Payload) LaunchImage(v string) *Payload {
	p.aps().Alert.LaunchImage = v
	return p
}

func (p *Payload) String() string {
	if data, err := json.Marshal(p); err != nil {
		return ""
	} else {
		return string(data)
	}
}

func (p *Payload) aps() *aps {
	return p.Aps
}
