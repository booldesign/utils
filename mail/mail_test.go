package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	options := &Options{
		MailHost: "smtp.163.com",
		MailPort: 25,
		MailUser: "***@163.com",
		MailPass: "****", //密码或授权码
		MailTo:   []string{"***@icloud.com", "***@vip.qq.com"},
		Subject:  "subject",
		Body:     "body",
	}
	err := Send(options)
	if err != nil {
		t.Error("Mail Send error", err)
		return
	}
	t.Log("success")
}
