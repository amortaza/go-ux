package ux

var Icon IconTable

type IconTable struct {
	Trash string
	Check string
	ChevronRight string
	Search string
	CircledCross string
}

func iconToStr(cp int) string {
	return string([]rune{rune(cp)})
}
