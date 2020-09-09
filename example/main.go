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
	"encoding/json"
	"fmt"
	"log"

	"github.com/lemoyxk/console"
)

type hook struct{}

func (h *hook) Fire(entry *console.Entry) {
	var bts, _ = json.Marshal(entry)
	log.Println(string(bts))
}

func main() {

	console.Reset.Println("1", "2", "3")
	console.BgBlue.Println("1", "2", "3")
	console.BgBlue.Printf("%s-%s-%s\n", "1", "2", "3")
	console.BgBlue.Printf("%s-%s-%s\n", "1", "2", "3")
	console.BgBlue.Printf("%s-%s-%s\n", "1", "2", "3")
	fmt.Println("1", "2", "3")
	var a = &hook{}
	console.SetHook(a)
	console.SetFlags(console.LEVEL)
	console.Info("hello")
	console.SetFormatter(console.NewJsonFormatter())
	console.Warning("hello")

	console.FgHiGreen.Info("hello", "world")
}
