package service

import (
	"github.com/jordan-wright/email"
	"net/smtp"
)

const (
	key = "key"
)

func SendEmail(name,time,reason string) error {
	e := email.NewEmail()
	e.From = name+"<my@qq.com>"
	e.To = []string{"to@qq.com"}
	e.Cc = []string{"my@qq.com"}
	e.Subject = "请假条"
	e.Text = []byte("请假人："+name+"\n请假人值班楼栋：知行苑2舍\n请假时间："+time+"\n请假原因："+reason)

	err := e.Send("smtp.qq.com:25", smtp.PlainAuth("", "my@qq.com", key, "smtp.qq.com"))
	return err
}