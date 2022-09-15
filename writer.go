/**
* @program: console
*
* @description:
*
* @author: lemo
*
* @create: 2020-10-02 02:16
**/

package console

import (
	"io"
	"os"
)

var writer io.Writer = os.Stdout

var errWriter io.Writer = os.Stderr

func SetWriter(w io.Writer) {
	writer = w
}

func SetErrorWriter(w io.Writer) {
	errWriter = w
}

func write(str string) {
	panicIfError(writer.Write([]byte(str)))
}

func errWrite(str string) {
	panicIfError(errWriter.Write([]byte(str)))
}
