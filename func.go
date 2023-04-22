/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-11 19:42
**/

package console

import (
	"bytes"
	"fmt"
)

func joinInterface(v []any, sep string) string {
	var buf bytes.Buffer
	for i := 0; i < len(v); i++ {
		switch v[i].(type) {
		case string:
			buf.WriteString(v[i].(string))
		default:
			buf.WriteString(fmt.Sprint(v[i]))
		}
		if i != len(v)-1 {
			buf.WriteString(sep)
		}
	}
	return buf.String()
}
