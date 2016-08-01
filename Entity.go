package ux

import (
	"time"
	"os"
	"io/ioutil"
	"fmt"
)

type Entity struct {
	File string
	Last time.Time
	Js   string
}

func (e *Entity) check() {
	//fmt.Println("Checkingn file ", e.File)

	info, _ := os.Stat(e.File)
	updated :=  info.ModTime()

	if e.Last != updated {

		e.Last = updated

		buf, _ := ioutil.ReadFile(e.File)

		e.Js = string(buf)

		fmt.Println("loading..." + e.File)
	}
}

func (e *Entity) Draw(x, y, w, h int, text string) {

	e.check()

	vm.Set("x", float32(x))
	vm.Set("y", float32(y))

	vm.Set("w", float32(w))
	vm.Set("h", float32(h))

	vm.Set("text", text)

	_, er := vm.Run(e.Js)

	if er != nil {
		fmt.Println(e)
	}
}
