package ux

import (
	"time"
	"os"
	"fmt"
	"io/ioutil"
)

type JsFile struct {

	filename    string
	updatedTime time.Time
	js          string
}

var g_jsFilesByFilename = make(map[string] *JsFile)

var g_lastLoopTime int64

func CheckJsFiles_throttled() bool {

	somethingChanged := false

	if time.Now().UnixNano() - g_lastLoopTime < 3000000000 { // 3.0 seconds

		// too fast
		return somethingChanged
	}

	for _, jsfile := range g_jsFilesByFilename {

		somethingChanged = somethingChanged || jsfile.check()
	}

	g_lastLoopTime = time.Now().UnixNano()

	return somethingChanged
}

func NewJsFile(filename string) *JsFile {

	jsfile, ok := g_jsFilesByFilename[filename]

	if !ok {

		jsfile = &JsFile{filename: filename}

		jsfile.check()

		g_jsFilesByFilename[filename] = jsfile
	}

	return jsfile
}

func (jsfile *JsFile) check() bool {

	info, err := os.Stat(jsfile.filename)

	if err != nil {
		fmt.Println(err, " Unable to open JS file ", jsfile.filename)
		panic("Unable to open JS file " + jsfile.filename)
	}

	if info == nil {
		fmt.Println("Unable to open file ", jsfile.filename)
		panic("Unable to open JS file " + jsfile.filename)
	}

	updatedTime :=  info.ModTime()

	if jsfile.updatedTime != updatedTime {

		jsfile.updatedTime = updatedTime

		buf, _ := ioutil.ReadFile(jsfile.filename)

		jsfile.js = string(buf)

		fmt.Println("loaded js file..." + jsfile.filename)

		return true
	}

	return false
}


