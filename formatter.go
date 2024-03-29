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

func NewTextFormatter() *TextFormatter {
	return &TextFormatter{}
}

type TextFormatter struct{}

func (f *TextFormatter) Format(entry *Entry) string {

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

	for i := 0; i < len(entry.Fields); i++ {
		var current = entry.Fields[i]
		if current.Key == "" && current.Value == nil {
			continue
		}

		if current.Key != "" && current.Value == nil {
			flags = append(flags, current.Key)
			continue
		}

		if current.Key == "" && current.Value != nil {
			flags = append(flags, fmt.Sprintf("%v", current.Value))
			continue
		}

		flags = append(flags, current.Key+":"+fmt.Sprintf("%v", current.Value))
	}

	var format = "%s " + entry.Format
	if len(flags) == 0 {
		format = "%s" + entry.Format
	}
	var args = append([]any{strings.Join(flags, " ")}, entry.Args...)

	return fmt.Sprintf(format, args...)
}

func NewJsonFormatter() *JsonFormatter {
	return &JsonFormatter{}
}

type JsonFormatter struct{}

func (f *JsonFormatter) Format(entry *Entry) string {
	var data = make(map[string]any)

	if entry.Level != "" {
		data["level"] = entry.Level
	}

	if !entry.Time.IsZero() {
		data["time"] = entry.Time.Format("2006-01-02 15:04:05")
	}

	if entry.File != "" {
		data["file"] = entry.File + ":" + strconv.Itoa(entry.Line)
	}

	for i := 0; i < len(entry.Fields); i++ {
		var current = entry.Fields[i]
		if current.Key == "" || current.Value == nil {
			continue
		}
		data[current.Key] = current.Value
	}

	data["message"] = fmt.Sprintf(entry.Format, entry.Args...)

	var bts, err = json.Marshal(data)
	if err != nil {
		panic(err)
	}

	if len(entry.Format) > 0 && entry.Format[len(entry.Format)-1] == '\n' {
		return string(bts) + "\n"
	}

	return string(bts)
}
