/**
* @program: console
*
* @description:
*
* @author: lemo
*
* @create: 2020-10-02 02:19
**/

package console

func panicIfError(args ...interface{}) {
	if len(args) == 0 {
		return
	}

	var err = args[len(args)-1]

	if err == nil {
		return
	}

	panic(err)
}
