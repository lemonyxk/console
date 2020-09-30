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
	"fmt"

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

func (c Color) Println(v ...interface{}) {
	var format = ""
	for i := 0; i < len(v); i++ {
		format += "%v "
	}
	format = format[:len(format)-1]
	fmt.Printf("%s\n", text.Color(c).Sprintf(format, v...))
}

func (c Color) Printf(format string, v ...interface{}) {
	fmt.Printf("%s", text.Color(c).Sprintf(format, v...))
}

func (c Color) Info(v ...interface{}) {
	var str = handlerLogger.GetLevelStringln(INF, "%s", v...)

	if handlerLogger.DisableColor {
		fmt.Printf("%s\n", str)
		return
	}

	c.Printf("%s\n", str)
}

func (c Color) Debug(v ...interface{}) {
	var str = handlerLogger.GetLevelStringln(DEB, "%s", v...)

	if handlerLogger.DisableColor {
		fmt.Printf("%s\n", str)
		return
	}

	c.Printf("%s\n", str)
}

func (c Color) Warning(v ...interface{}) {
	var str = handlerLogger.GetLevelStringln(WAR, "%s", v...)

	if handlerLogger.DisableColor {
		fmt.Printf("%s\n", str)
		return
	}

	c.Printf("%s\n", str)
}

func (c Color) Error(v ...interface{}) {
	var str = handlerLogger.GetLevelStringln(ERR, "%s", v...)

	if handlerLogger.DisableColor {
		fmt.Printf("%s\n", str)
		return
	}

	c.Printf("%s\n", str)
}

func (c Color) Infof(format string, v ...interface{}) {
	var str = handlerLogger.GetLevelStringf(INF, format, v...)

	if handlerLogger.DisableColor {
		fmt.Printf("%s", str)
		return
	}

	c.Printf("%s", str)
}

func (c Color) Warningf(format string, v ...interface{}) {
	var str = handlerLogger.GetLevelStringf(WAR, format, v...)

	if handlerLogger.DisableColor {
		fmt.Printf("%s", str)
		return
	}

	c.Printf("%s", str)
}

func (c Color) Debugf(format string, v ...interface{}) {
	var str = handlerLogger.GetLevelStringf(DEB, format, v...)

	if handlerLogger.DisableColor {
		fmt.Printf("%s", str)
		return
	}

	c.Printf("%s", str)
}

func (c Color) Errorf(format string, v ...interface{}) {
	var str = handlerLogger.GetLevelStringf(ERR, format, v...)

	if handlerLogger.DisableColor {
		fmt.Printf("%s", str)
		return
	}

	c.Printf("%s", str)
}
