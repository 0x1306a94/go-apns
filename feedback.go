package apns

import (
	"bytes"
	"crypto/tls"
	"encoding/binary"
	"encoding/hex"
	"net"
	"time"
	"fmt"
	"io"
)

const (
	_FeedbackHostDevelopment = "feedback.sandbox.push.apple.com:2196"
	_FeedbackHostProduction  = "feedback.push.apple.com:2196"
	_DefaultFeedbackHost     = _FeedbackHostDevelopment
)

type Feedback struct {
	host   string
	tlsCfg *tls.Config
	conn   net.Conn
}

type FeedbackTuple struct {
	Timestamp   time.Time
	TokenLength uint16
	DeviceToken string
}

func feedbackTupleFromBytes(b []byte) FeedbackTuple {
	r := bytes.NewReader(b)

	var ts uint32
	binary.Read(r, binary.BigEndian, &ts)

	var tokLen uint16
	binary.Read(r, binary.BigEndian, &tokLen)

	tok := make([]byte, tokLen)
	binary.Read(r, binary.BigEndian, &tok)

	return FeedbackTuple{
		Timestamp:   time.Unix(int64(ts), 0),
		TokenLength: tokLen,
		DeviceToken: hex.EncodeToString(tok),
	}
}

func NewFeedBackWithCer(certificate tls.Certificate) *Feedback {
	return &Feedback{
		tlsCfg: &tls.Config{
			Certificates: []tls.Certificate{certificate},
		},
		host: _DefaultFeedbackHost,
	}
}

func (f *Feedback) Development() *Feedback {
	f.host = _FeedbackHostDevelopment
	return f
}

func (f *Feedback) Production() *Feedback {
	f.host = _FeedbackHostProduction
	return f
}

func (f *Feedback) Host() string {
	return f.host
}

func (f *Feedback) connection() error {
	if f.conn != nil {
		f.conn.Close()
	}
	conn, err := tls.Dial("tcp", f.host, f.tlsCfg)
	if err != nil {
		return err
	}
	f.conn = conn
	return nil
}

func (f *Feedback) disConnection() error {
	if f.conn != nil {
		err := f.conn.Close()
		f.conn = nil
		return err
	}
	return nil
}

func (f *Feedback) Receive() <-chan FeedbackTuple {
	fc := make(chan FeedbackTuple)
	go f.receive(fc)
	return fc
}

func (f *Feedback) receive(fc chan FeedbackTuple) {
	err := f.connection()
	if err != nil {
		close(fc)
		fmt.Println("Feedback error: ", err)
		return
	}
	defer f.disConnection()
	for {
		b := make([]byte, 38)
		f.conn.SetReadDeadline(time.Now().Add(100 * time.Millisecond))
		_, err := f.conn.Read(b)
		if nerr, ok := err.(net.Error); ok && nerr.Timeout() {
			fmt.Println("Feedback Timeout...")
			continue
		}
		if err == io.EOF {
			fmt.Println("Feedback start sleep...")
			time.Sleep(time.Minute * 10)
			fmt.Println("Feedback end sleep...")
			continue
		} else {
			close(fc)
			fmt.Println("Feedback error: ", err)
			return
		}
		fc <- feedbackTupleFromBytes(b)
	}
}
