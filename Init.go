package ux

import (
	vgo "github.com/shibukawa/nanovgo"
	"github.com/robertkrimen/otto"
	"fmt"
)

var Ctx *vgo.Context
var vm *otto.Otto

func Init() {

	vm = otto.New()

	var err error

	Ctx, err = vgo.NewContext(vgo.AntiAlias | vgo.StencilStrokes /*| nanovgo.Debug*/)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	vm.Set("ctx", Ctx)

	Ctx.CreateFont("icons", "github.com/shibukawa/nanovgo/sample/entypo.ttf")
	Ctx.CreateFont("sans", "github.com/shibukawa/nanovgo/sample/Roboto-Regular.ttf")

	fmt.Println("(+) Created nanovgo context")

	Icon.Carrot_Down = iconToStr(0xe75c)
	Icon.Carrot_Up = iconToStr(0xe75f)
	Icon.Check = iconToStr(0x2713)
	Icon.ChevronRight = iconToStr(0xE75E)
	Icon.CircledCross = iconToStr(0x2716)
	Icon.Search = iconToStr(0x1F50D)
	Icon.Trash = iconToStr(0xE729)

	vm.Set("vgoLinearGradient", LinearGradient)
	vm.Set("vgoBoxGradient", BoxGradient)
	vm.Set("vgoRadialGradient", RadialGradient)
	vm.Set("vgoRGBA", RGBA)

	vm.Set("vgoAlignLeft", vgo.AlignLeft)
	vm.Set("vgoAlignMiddle", vgo.AlignMiddle)
	vm.Set("vgoAlignCenter", vgo.AlignCenter)
	vm.Set("vgoAlignBaseline", vgo.AlignBaseline)
	vm.Set("vgoAlignBottom", vgo.AlignBottom)
	vm.Set("vgoAlignRight", vgo.AlignRight)
	vm.Set("vgoAlignTop", vgo.AlignTop)

	vm.Set("vgoPathWindingHole", PathWindingHole)

	vm.Set("sysGetTextWidth", GetTextWidth)
	vm.Set("sysGetTextHeight", GetTextHeight)

	vm.Set("GuiIconCheck", Icon.Check)
	vm.Set("GuiIconCarrotUp", Icon.Carrot_Up)
	vm.Set("GuiIconCarrotDown", Icon.Carrot_Down)
	vm.Set("GuiIconSearch", Icon.Search)
	vm.Set("GuiIconCircledCross", Icon.CircledCross)
	vm.Set("GuiIconChevronRight", Icon.ChevronRight)
}

func Uninit() {
	Ctx.Delete()
}

func PathWindingHole() {
	Ctx.PathWinding(vgo.Hole)
}

func RGBA(r,g,b,a int) vgo.Color {
	return vgo.RGBA(uint8(r),uint8(g),uint8(b),uint8(a))
}

func LinearGradient(sx,sy,ex,ey int, iColor,oColor vgo.Color) vgo.Paint {
	return vgo.LinearGradient(float32(sx),float32(sy),float32(ex),float32(ey),iColor,oColor)
}

func BoxGradient(x,y,w,h,r,f int, i,o vgo.Color) vgo.Paint {
	return vgo.BoxGradient(float32(x), float32(y), float32(w), float32(h), float32(r), float32(f), i, o)
}

func RadialGradient(cx,cy,inr,outr int, i,o vgo.Color) vgo.Paint {
	return vgo.RadialGradient(float32(cx), float32(cy), float32(inr), float32(outr), i, o)
}

func GetTextWidth(text string) float32 {

	w, _ := Ctx.TextBounds(0, 0, text)

	return w
}

func GetTextHeight(text string) float32 {

	_, a := Ctx.TextBounds(0, 0, text)

	fmt.Println(len(a))
	fmt.Println(a[0], a[1], a[2], a[3])

	return a[3] - a[1]
}




