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
	"reflect"
)

type em struct{}

var packageName = reflect.TypeOf(em{}).PkgPath()

func joinInterface(v []interface{}, sep string) string {
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
