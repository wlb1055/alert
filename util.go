package alert

import (
	"fmt"
	"runtime"
)

func TraceId() (traceId string, ok bool) {
	_, f, l, ok := runtime.Caller(2)
	if !ok {
		return
	}
	return fmt.Sprintf("%s:%d", f, l), true
}
