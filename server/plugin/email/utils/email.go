package utils

import (
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/sunyihoo/gin-vue3-admin/server/plugin/email/global"
	"net/smtp"
	"strings"
)

//@author: [maplepie]()
//@function: Email
//@description: Email 发送方法
//@param: subject string, body string
//@return: error

func Email(To, subject string, body string) error {
	to := strings.Split(To, ",")
	return send(to, subject, body)
}

//@author: [s]()
//@function: ErrorToEmail
//@description:  给email中间件错误发送邮件到指定邮箱
//@param: subject string, body string
//@return: error

func ErrotToEmail(subject string, body string) error {
	to := strings.Split(global.Config.To, ",")
	if to[len(to)-1] == "" { // 判断切片的最后一个元素是否为空，为空则移除
		to = to[:len(to)-1]
	}
	return send(to, subject, body)
}

//@author: [m]()
//@function: EmailTest
//@description:  Email测试方法
//@param: subject string, body string
//@return: error

func EmailTest(subject string, body string) error {
	to := []string{global.Config.To}
	return send(to, subject, body)
}

func send(to []string, subject string, body string) error {
	from := global.Config.From
	nickName := global.Config.Nickname
	secret := global.Config.Secret
	host := global.Config.Host
	port := global.Config.Port
	isSSL := global.Config.IsSSL

	auth := smtp.PlainAuth("", from, secret, host)
	e := email.NewEmail()
	if nickName != "" {
		e.From = fmt.Sprintf("%s <%s>", nickName, from)
	} else {
		e.From = from
	}
	e.To = to
	e.Subject = subject
	e.HTML = []byte(body)
	var err error
	hostAddr := fmt.Sprintf("%s:%d", host, port)
	if isSSL {
		err = e.SendWithTLS(hostAddr, auth, &tls.Config{ServerName: host})
	} else {
		err = e.Send(hostAddr, auth)
	}
	return err

}
