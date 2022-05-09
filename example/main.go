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
	"net/http"

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

func (h *hook) Fire(entry *console.Entry) {
	// var bts, _ = json.Marshal(entry)
	// log.Println(string(bts))
}

func main() {

	var con = console.NewLogger()

	con.Info("hello")

	console.Pretty.Dump(Test{
		Person: []struct{ Name string }{{Name: "a"}},
		One:    0,
		Three:  map[int]string{},
		Two:    0,
	})

	console.Pretty.Dump(console.NewLogger())

	console.Reset.Println("1", "2", "3")
	console.BgBlue.Println("1", "2", "3")
	console.BgBlue.Printf("%s-%s-%s\n", "1", "2", "3")
	console.BgBlue.Printf("%s-%s-%s\n", "1", "2", "3")
	console.BgBlue.Printf("%s-%s-%s\n", "1", "2", "3")
	fmt.Println("1", "2", "3")
	var a = &hook{}
	console.SetHook(a)
	// console.SetFlags(console.LEVEL)
	console.Info("hello")
	console.AddField("name", "lemo")
	console.SetFormatter(console.NewJsonFormatter())
	console.Warningf("%s???\n", "hello")

	console.Pretty.Dump(&http.Server{})
	type b struct {
		a any
	}
	console.Pretty.Dump(b{a: true}, b{a: false})
}
