package ux

var Icon IconTable

type IconTable struct {
	Carrot_Up, Carrot_Down string
	Check string
	ChevronRight string
	CircledCross string
	Search string
	Trash string
}

func iconToStr(cp int) string {
	return string([]rune{rune(cp)})
}
