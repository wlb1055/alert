package alert

import (
	"fmt"
	"runtime"
)

var HtmlTmpl = `
<!DOCTYPE html>
<html>
<head>
</head>
<body>

<pre>%s</pre>
<pre>%s</pre>

</body>
</html>
`

func TraceId() (traceId string, ok bool) {
	_, f, l, ok := runtime.Caller(2)
	if !ok {
		return
	}
	return fmt.Sprintf("%s:%d", f, l), true
}
