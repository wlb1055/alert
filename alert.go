package alert

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"gopkg.in/gomail.v2"
)

var dialer = &gomail.Dialer{}

func Init(opts ...Option) {
	dialer = gomail.NewDialer("smtp.qq.com", 587, "", "")
	for _, opt := range opts {
		opt()
	}
}

func From(from string) Option {
	return func() {
		dialer.Username = from
	}
}

func Password(password string) Option {
	return func() {
		dialer.Password = password
	}
}

func SendMail(conn redis.Conn, module, subject, body string, to ...string) (e error) {
	if "" == dialer.Username {
		return errors.New("do setup first")
	}

	if nil == conn {
		return errors.New("conn=nil")
	}

	ok, e := hitRule(module, subject, conn)
	if e != nil || !ok {
		return
	}

	return sendMail(subject, body, to...)
}

func sendMail(subject, body string, to ...string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", dialer.Username)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", body)

	if e := dialer.DialAndSend(m); e != nil {
		return e
	}

	return nil
}
