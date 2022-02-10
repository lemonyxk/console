/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-11 18:19
**/

package console

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type Formatter interface {
	Format(entry *Entry) string
}

func NewTextFormatter() *textFormatter {
	return &textFormatter{}
}

type textFormatter struct{}

func (f *textFormatter) Format(entry *Entry) string {

	var flags []string

	if entry.Level != "" {
		flags = append(flags, string(entry.Level))
	}

	if !entry.Time.IsZero() {
		flags = append(flags, entry.Time.Format("2006-01-02 15:04:05"))
	}

	if entry.File != "" {
		flags = append(flags, entry.File+":"+strconv.Itoa(entry.Line))
	}

	if entry.ID != "" {
		flags = append(flags, entry.ID)
	}

	var format = "%s " + entry.Format
	if len(flags) == 0 {
		format = entry.Format
	}
	var args = append([]interface{}{strings.Join(flags, " ")}, entry.Args...)

	return fmt.Sprintf(format, args...)
}

func NewJsonFormatter() *jsonFormatter {
	return &jsonFormatter{}
}

type jsonFormatter struct{}

func (f *jsonFormatter) Format(entry *Entry) string {
	var data = make(map[string]interface{})

	if entry.Level != "" {
		data["level"] = entry.Level
	}

	if !entry.Time.IsZero() {
		data["time"] = entry.Time.Format("2006-01-02 15:04:05")
	}

	if entry.File != "" {
		data["file"] = entry.File + ":" + strconv.Itoa(entry.Line)
	}

	if entry.ID != "" {
		data["id"] = entry.ID
	}

	data["msg"] = fmt.Sprintf(entry.Format, entry.Args...)

	var bts, err = json.Marshal(data)
	if err != nil {
		panic(err)
	}

	if len(entry.Format) > 0 && entry.Format[len(entry.Format)-1] == '\n' {
		return string(bts) + "\n"
	}

	return string(bts)
}
