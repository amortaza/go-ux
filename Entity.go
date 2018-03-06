package ux

import (
	"time"
	"os"
	"io/ioutil"
	"fmt"
)

type Entity struct {
	Filename string
	Last     time.Time
	Js       string
}

func (e *Entity) check() {

	info, err := os.Stat(e.Filename)

	if err != nil {
		fmt.Println(err, " Unable to open file ", e.Filename)
		panic("Unable to open file " + e.Filename)
	}

	if info == nil {
		fmt.Println("Unable to open file ", e.Filename)
		panic("Unable to open file " + e.Filename)
	}

	updated :=  info.ModTime()

	if e.Last != updated {

		e.Last = updated

		buf, _ := ioutil.ReadFile(e.Filename)

		e.Js = string(buf)

		fmt.Println("loading..." + e.Filename)
	}
}

func (e *Entity) SetFloat(name string, value float32) {
	vm.Set(name, value)
}

func (e *Entity) SetInt(name string, value int) {
	vm.Set(name, value)
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
