package apns

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"errors"

	"github.com/king129/go-apns/token"
	context "golang.org/x/net/context"
	"golang.org/x/net/http2"
)

const (
	_HostDevelopment = "https://api.development.push.apple.com"
	_HostProduction  = "https://api.psuh.apple.com"
	_DefaultHost     = _HostDevelopment
)

var (
	_TLSDialTimeOut    = 20 * time.Second
	_HTTPClientTimeOut = 60 * time.Second
	_TCPKeepAlive      = 60 * time.Second
)

var dialTLS = func(network, addr string, cfg *tls.Config) (net.Conn, error) {
	dialer := &net.Dialer{
		Timeout:   _TLSDialTimeOut,
		KeepAlive: _TCPKeepAlive,
	}
	return tls.DialWithDialer(dialer, network, addr, cfg)
}

type clientConnectionCloser interface {
	CloseIdleConnections()
}

type Client struct {
	host        string
	certificate tls.Certificate
	httpClient  *http.Client
	token       *token.Token
}

func NewClientWithCer(certificate tls.Certificate) *Client {
	tlsCfg := &tls.Config{
		Certificates: []tls.Certificate{certificate},
	}
	if len(certificate.Certificate) > 0 {
		tlsCfg.BuildNameToCertificate()
	}
	transport := &http2.Transport{
		TLSClientConfig: tlsCfg,
		DialTLS:         dialTLS,
	}

	return &Client{
		httpClient: &http.Client{
			Transport: transport,
			Timeout:   _HTTPClientTimeOut,
		},
		certificate: certificate,
		host:        _DefaultHost,
	}
}

func NewClientWithToken(token *token.Token) *Client {
	transport := &http2.Transport{
		DialTLS: dialTLS,
	}
	return &Client{
		httpClient: &http.Client{
			Transport: transport,
			Timeout:   _HTTPClientTimeOut,
		},
		host:  _DefaultHost,
		token: token,
	}
}

func (c *Client) Development() *Client {
	c.host = _HostDevelopment
	return c
}

func (c *Client) Production() *Client {
	c.host = _HostProduction
	return c
}

func (c *Client) Host() string {
	return c.host
}

func (c *Client) Close() {
	c.CloseIdleConnections()
}

func (c *Client) CloseIdleConnections() {
	c.httpClient.Transport.(clientConnectionCloser).CloseIdleConnections()
}
func (c *Client) Push(m *Message) (*Response, error) {
	return c.PushWithContext(nil, m)
}

func (c *Client) PushWithContext(ctx context.Context, m *Message) (*Response, error) {

	if m == nil {
		return nil, errors.New("message is not empty")
	}
	if err := m.validationIsAvailable(); err != nil {
		return nil, err
	}
	var (
		payload []byte
		err     error
	)

	if payload, err = m.requestData(); err != nil {
		return nil, err
	}

	url := fmt.Sprintf("%v/3/device/%v", c.host, m.DeviceToken)
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(payload))
	if err != nil {
		return nil, err
	}

	if c.token != nil {
		c.setTokenHeader(req)
	}

	setHeaders(req, m)
	httpResp, err := c.requestWithContext(ctx, req)
	if err != nil {
		return nil, err
	}

	defer httpResp.Body.Close()

	data, err := ioutil.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	response := new(Response)
	response.StatusCode = httpResp.StatusCode
	response.ApnsID = httpResp.Header.Get("Apns-id")
	if len(data) == 0 && response.StatusCode == http.StatusOK {
		return response, nil
	} else if len(data) > 0 {
		if err = json.Unmarshal(data, response); err != nil {
			return nil, err
		}
		return response, nil
	}
	return nil, errors.New("unknown error")

}

func (c *Client) requestWithContext(ctx context.Context, req *http.Request) (*http.Response, error) {
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	return c.httpClient.Do(req)
}

func (c *Client) setTokenHeader(r *http.Request) {
	c.token.GenerateIfExpired()
	r.Header.Set("authorization", fmt.Sprintf("bearer %v", c.token.Bearer))
}

func setHeaders(r *http.Request, m *Message) {
	r.Header.Set("Content-Type", "application/json; charset=utf-8")
	header := m.requestHeader()
	for k, v := range header {
		r.Header.Set(k, v)
	}
}
