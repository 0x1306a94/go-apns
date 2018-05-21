package token

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"io/ioutil"
	"sync"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

const (
	TokenExpired = 3000
)

var (
	AuthKeyNotPEM   = errors.New("authKey must be a valid p8 PEM file")
	AuthKeyNotECDSA = errors.New("authKey must be of type ecdsa.PrivateKey")
	AuthKeyNil      = errors.New("authKey was nil")
)

type Token struct {
	sync.Mutex
	AuthKey  *ecdsa.PrivateKey
	KeyID    string
	TeamID   string
	Bearer   string
	IssuedAt int64
}

func AuthKeyFromFile(fileName string) (*ecdsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	return AuthKeyFromData(data)
}

func AuthKeyFromData(data []byte) (*ecdsa.PrivateKey, error) {
	block, _ := pem.Decode(data)
	if block == nil {
		return nil, AuthKeyNotPEM
	}
	key, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	switch key.(type) {
	case *ecdsa.PrivateKey:
		return key.(*ecdsa.PrivateKey), nil
	default:
		return nil, AuthKeyNotECDSA

	}
}

func NewToken(authKeyPath, teamID, keyID string) (*Token, error) {
	if authKeyPath == "" {
		return nil, errors.New("authKeyPath cannot be empty")
	}
	if teamID == "" {
		return nil, errors.New("teamID cannot be empty")
	}
	if keyID == "" {
		return nil, errors.New("keyID cannot be empty")
	}
	authKey, err := AuthKeyFromFile(authKeyPath)
	if err != nil {
		return nil, err
	}
	token := &Token{
		AuthKey: authKey,
		TeamID:  teamID,
		KeyID:   keyID,
	}
	return token, nil
}

func (t *Token) Expired() bool {
	return time.Now().Unix() >= (t.IssuedAt + TokenExpired)
}
func (t *Token) GenerateIfExpired() {
	t.Lock()
	defer t.Unlock()
	if t.Expired() {
		t.Generate()
	}
}

func (t *Token) Generate() (bool, error) {
	if t.AuthKey == nil {
		return false, AuthKeyNil
	}
	issuedAt := time.Now().Unix()
	jwtToken := &jwt.Token{
		Header: map[string]interface{}{
			"alg": "ES256",
			"kid": t.KeyID,
		},
		Claims: jwt.MapClaims{
			"iss": t.TeamID,
			"iat": issuedAt,
		},
		Method: jwt.SigningMethodES256,
	}
	bearer, err := jwtToken.SignedString(t.AuthKey)
	if err != nil {
		return false, err
	}
	t.IssuedAt = issuedAt
	t.Bearer = bearer
	return true, nil
}
