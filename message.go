package apns

import (
	"encoding/json"
	"errors"
	"fmt"
)

type MessagePriority int

// HTTP/2 的APNS 协议支持最大长度为4kb 数据
const MessagePayloadMaxLength = 4096

const (
	MessagePriortyLow       MessagePriority = 5  // 推送消息立即发送
	MessagePriortyHigh      MessagePriority = 10 // 推送消息在节省接收设备电源的时间发送
	MessageAPNSSoundDefault                 = "default"
)

var (
	MessageNotTopicError       = errors.New("topic cannot be empty")
	MessageNotDeviceTokenError = errors.New("deviceToken cannot be empty")
	MessageNotPayloadError     = errors.New("payload cannot be empty")
	MessagePayloadLargeError   = errors.New("the payload exceeds the maximum length and the maximum length is 4096 bytes")
)

// 官方文档
// https://developer.apple.com/documentation/usernotifications/setting_up_a_remote_notification_server/sending_notification_requests_to_apns
type Message struct {
	ApnsID      string          `json:"apns_id,omitempty"`
	CollapseID  string          `json:"collapse_id,omitempty"`
	DeviceToken string          `json:"device_token"`
	Topic       string          `json:"topic"` // App Bundle ID
	Expiration  int64           `json:"expiration,omitempty"`
	Priority    MessagePriority `json:"priority,omitempty"`
	Payload     *Payload        `json:"payload"`
}

func (m *Message) validationIsAvailable() error {
	if m.Topic == "" {
		return MessageNotTopicError
	}
	if m.DeviceToken == "" {
		return MessageNotDeviceTokenError
	}
	if m.Payload == nil {
		return MessageNotPayloadError
	}
	return nil
}

func (m *Message) requestData() ([]byte, error) {
	if err := m.validationIsAvailable(); err != nil {
		return nil, err
	}
	data, err := json.Marshal(m.Payload)
	if err != nil {
		return nil, err
	}

	if len(data) > MessagePayloadMaxLength {
		return nil, MessagePayloadLargeError
	}
	return data, nil
}

func (m *Message) requestHeader() map[string]string {
	header := make(map[string]string)
	if m.ApnsID != "" {
		header["apns-id"] = m.ApnsID
	}
	if m.Topic != "" {
		header["apns-topic"] = m.Topic
	}
	if m.CollapseID != "" {
		header["apns-collapse-id"] = m.CollapseID
	}
	if m.Priority == MessagePriortyLow || m.Priority == MessagePriortyHigh {
		header["apns-priority"] = fmt.Sprintf("%v", m.Priority)
	} else {
		header["apns-priority"] = fmt.Sprintf("%v", MessagePriortyHigh)
	}
	if m.Expiration > 0 {
		header["apns-expiration"] = fmt.Sprintf("%v", m.Expiration)
	}
	return header
}
