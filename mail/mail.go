package mail

import (
	"gopkg.in/gomail.v2"
)

type Options struct {
	MailHost string
	MailPort int
	MailUser string   // 发件人
	MailPass string   // 发件人密码/授权码
	MailTo   []string // 收件人
	Subject  string   // 邮件主题
	Body     string   // 邮件内容
}

func Send(options *Options) error {

	message := gomail.NewMessage()

	//设置发件人
	message.SetHeader("From", options.MailUser)

	//设置发送给多个用户
	message.SetHeader("To", options.MailTo...)

	//设置邮件主题
	message.SetHeader("Subject", options.Subject)

	//设置邮件正文
	message.SetBody("text/html", options.Body)

	dialer := gomail.NewDialer(options.MailHost, options.MailPort, options.MailUser, options.MailPass)

	return dialer.DialAndSend(message)
}
