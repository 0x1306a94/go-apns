package apns

import (
	"testing"
	"os"
	"github.com/king129/go-apns/cer"
	"fmt"
)

func TestFeedback(t *testing.T) {
	p12Path := os.Getenv("P12PATH")
	p12Passwd := os.Getenv("P12PASSWD")
	
	cer, err := cert.FromP12File(p12Path, p12Passwd)
	if err != nil {
		t.Error(err)
	}
	
	feedback := NewFeedBackWithCer(cer)
	for v := range feedback.Receive() {
		fmt.Println(v)
	}
}