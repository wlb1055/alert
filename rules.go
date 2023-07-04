package alert

import (
	"errors"
	"github.com/gomodule/redigo/redis"
)

//告警周期，如每5分钟告警一次

type Rule struct {
	Minutes int
	Times   int
}

var rules = map[string]*Rule{
	//
}

//var rules = map[string]*struct {
//	Minutes int32
//	Times   int32
//}{
//	"wh": {
//		Minutes: 5,
//		Times:   1,
//	},
//}

func Add(module string, rule *Rule) {
	rules[module] = rule
}

func hitRule(module, subject string, conn redis.Conn) (ok bool, e error) {
	rule, ok := rules[module]
	if !ok {
		return false, errors.New("alert cycle is not set")
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
