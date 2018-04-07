package ux

import (
	"fmt"
)

type Entity struct {

	jsfile *JsFile
}

func NewEntity(filename string) *Entity {

	entity := &Entity{}

	entity.jsfile = NewJsFile(filename)

	return entity
}

func (e *Entity) SetFloat(name string, value float32) {
	vm.Set(name, value)
}

func (e *Entity) SetInt(name string, value int) {
	vm.Set(name, value)
}

func (e *Entity) Draw(x, y, w, h int, text string) {

	e.jsfile.check()

	vm.Set("x", float32(x))
	vm.Set("y", float32(y))

	vm.Set("w", float32(w))
	vm.Set("h", float32(h))

	vm.Set("text", text)

	_, er := vm.Run(e.jsfile.js)

	if er != nil {
		fmt.Println(e)
	}
}
