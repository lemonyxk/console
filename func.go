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
	"os"
	"runtime"
	"strings"
	"unsafe"
)

func caller(deep int) (string, int) {
	_, file, line, ok := runtime.Caller(deep + 1)
	if !ok {
		return "", 0
	}

	var rootPath, err = os.Getwd()
	if err != nil {
		return file, line
	}

	if runtime.GOOS == "windows" {
		rootPath = strings.Replace(rootPath, "\\", "/", -1)
	}

	if rootPath == "/" {
		return file, line
	}
	if strings.HasPrefix(file, rootPath) {
		file = file[len(rootPath)+1:]
	}

	return file, line
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	return (*eFace)(unsafe.Pointer(&i)).data == nil
}

type eFace struct {
	_type unsafe.Pointer
	data  unsafe.Pointer
}

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
