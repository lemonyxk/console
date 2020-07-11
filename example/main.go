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
	"log"

	"github.com/lemoyxk/console"
)

type hook struct{}

func (h *hook) Fire(entry *console.Entry) {
	var bts, _ = json.Marshal(entry)
	log.Println(string(bts))
}

func main() {

	var a = &hook{}
	console.SetHook(a)
	console.Info("hello")
	console.SetColor(false)
	console.SetFormatter(console.NewJsonFormatter())
	console.Info("hello")
}
