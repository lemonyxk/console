/**
* @program: console
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-11 21:09
**/

package main

import "github.com/Lemo-yxk/console"

func main() {
	console.Info("hello")
	console.SetFormatter(console.NewJsonFormatter())
	console.Info("hello")
}
