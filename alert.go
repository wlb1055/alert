package alert

import (
	"errors"
	"github.com/garyburd/redigo/redis"
	"gopkg.in/gomail.v2"
)

var from string

var dialer = &gomail.Dialer{}

func Setup(username, password string) {
	from = username
	dialer = gomail.NewDialer("smtp.qq.com", 587, username, password)
}

func SendMail(conn redis.Conn, module, subject, body string, to ...string) (e error) {
	if "" == from {
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

func hitRule(module, subject string, conn redis.Conn) (ok bool, e error) {
	rule, ok := rules[module]
	if !ok {
		return
	}

	n, e := redis.Int(conn.Do("INCR", subject))
	if e != nil {
		return
	}
	if 1 == n {
		conn.Do("EXPIRE", subject, rule.Minutes*60)
	}

	return rule.Times == n, nil
}

func sendMail(subject, body string, to ...string) error {
	m := gomail.NewMessage()

	m.SetHeader("From", from)
	m.SetHeader("To", to...)
	m.SetHeader("Subject", subject)

	m.SetBody("text/html", body)

	if e := dialer.DialAndSend(m); e != nil {
		return e
	}

	return nil
}
