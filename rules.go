package alert

type Rule struct {
	Minutes int
	Times   int
}

var rules = map[string]*Rule{
	//
}

func Add(module string, rule *Rule) {
	rules[module] = rule
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
