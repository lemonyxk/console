/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-11-27 20:16
**/

package console

import (
	"github.com/jedib0t/go-pretty/text"
)

type Color int

// Base colors -- attributes in reality
const (
	Reset Color = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground colors
const (
	FgBlack Color = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity colors
const (
	FgHiBlack Color = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background colors
const (
	BgBlack Color = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity colors
const (
	BgHiBlack Color = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

type Colors []Color

func (c Colors) Println(v ...any) {
	var str = joinInterface(v, " ")
	var colors = text.Colors{}
	for i := 0; i < len(c); i++ {
		colors = append(colors, text.Color(c[i]))
	}
	write(colors.Sprint(str + "\n"))
}

func (c Colors) Sprintf(format string, v ...any) string {
	var colors = text.Colors{}
	for i := 0; i < len(c); i++ {
		colors = append(colors, text.Color(c[i]))
	}
	return colors.Sprintf(format, v...)
}

func (c Colors) Sprint(v ...any) string {
	var colors = text.Colors{}
	for i := 0; i < len(c); i++ {
		colors = append(colors, text.Color(c[i]))
	}
	return colors.Sprint(v...)
}

func (c Colors) Print(v ...any) {
	var colors = text.Colors{}
	for i := 0; i < len(c); i++ {
		colors = append(colors, text.Color(c[i]))
	}
	write(colors.Sprint(v...))
}

func (c Colors) Printf(format string, v ...any) {
	var colors = text.Colors{}
	for i := 0; i < len(c); i++ {
		colors = append(colors, text.Color(c[i]))
	}
	write(colors.Sprintf(format, v...))
}

func (c Color) Mixed(color ...Color) Colors {
	var colors = Colors{}
	colors = append(colors, c)
	for i := 0; i < len(color); i++ {
		colors = append(colors, color[i])
	}
	return colors
}

func (c Color) Println(v ...any) {
	var str = joinInterface(v, " ")
	write(text.Color(c).Sprint(str + "\n"))
}

func (c Color) Sprintf(format string, v ...any) string {
	return text.Color(c).Sprintf(format, v...)
}

func (c Color) Sprint(v ...any) string {
	return text.Color(c).Sprint(v...)
}

func (c Color) Print(v ...any) {
	write(text.Color(c).Sprint(v...))
}

func (c Color) Printf(format string, v ...any) {
	write(text.Color(c).Sprintf(format, v...))
}
