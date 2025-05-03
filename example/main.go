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

import (
	"fmt"
	"sync"
	"time"

	"github.com/lemonyxk/console"
)

type Test struct {
	Three  map[int]string
	Person []struct {
		Name string
	}

	One int
	Two int
}

type hook struct{}

type Name any

func main() {

	// var con = console.NewLogger()
	//
	// con.Info("hello")
	//
	// console.Pretty.Dump(Test{
	// 	Person: []struct{ Name string }{{Name: "a"}},
	// 	One:    0,
	// 	Three:  map[int]string{},
	// 	Two:    0,
	// })
	//
	// console.Pretty.Dump(console.NewLogger())
	//
	// console.Reset.Println("1", "2", "3")
	// console.BgBlue.Println("1", "2", "3")
	// console.BgBlue.Printf("%s-%s-%s\n", "1", "2", "3")
	// console.BgBlue.Printf("%s-%s-%s\n", "1", "2", "3")
	// console.BgBlue.Printf("%s-%s-%s\n", "1", "2", "3")
	// fmt.Println("1", "2", "3")
	// var a = &hook{}
	// console.SetHook(a)
	// // console.SetFlags(console.LEVEL)
	// console.Info("hello")
	// console.AddField("name", "lemo")
	// console.SetFormatter(console.NewJsonFormatter())
	// console.Warningf("%s???\n", "hello")
	//
	// console.Pretty.Dump(&http.Server{})
	// type b struct {
	// 	a any
	// }
	// console.Pretty.Dump(b{a: true}, b{a: false})
	console.Pretty.Dump(struct {
		Time  time.Time
		slice []Test
	}{Time: time.Now(), slice: []Test{{}, {}, {}, {}, {}, {}}})

	fmt.Printf("%+v\n", struct {
		Time  time.Time
		slice []Test
	}{Time: time.Now(), slice: []Test{{}, {}, {}}})

	console.Pretty.Dump(struct {
		ch chan int
	}{make(chan int)})

	console.Color.Println(console.Underline, "hello world", "aaa")

	console.Error.Log(`hello world`)

	console.Error.Log(`hello world`)
	console.Error.Log(`hello world`)
	console.Error.Log(`hello world`)

	console.Info.Logf("%s", "hello world")
	console.LevelInfoValue("info")
	console.Info.Logf("%s", "hello world1")

	console.Info.Logf("%s", "hello world3")

	type B struct {
		Map sync.Map
	}

	var b = B{}

	console.Pretty.Dump(b)
}
