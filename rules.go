package alert

type Rule struct {
	Minutes int
	Times   int
}

var Rules = map[string]*Rule{
	//
}

//var Rules = map[string]*struct {
//	Minutes int32
//	Times   int32
//}{
//	"wh": {
//		Minutes: 5,
//		Times:   1,
//	},
//}
