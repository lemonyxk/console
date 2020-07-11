/**
* @program: console
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-11 21:10
**/

package console

type Hook interface {
	Fire(entry *Entry)
}
