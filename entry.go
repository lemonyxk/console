/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2020-07-11 18:21
**/

package console

import "time"

type Level string

const (
	ERR Level = "ERR"
	INF Level = "INF"
	DEB Level = "DEB"
	WAR Level = "WAR"
)

type Field struct {
	Key   string
	Value any
}

type Entry struct {
	Level  Level     `json:"level,omitempty"`
	File   string    `json:"file,omitempty"`
	Line   int       `json:"line,omitempty"`
	Time   time.Time `json:"time,omitempty"`
	Format string    `json:"format"`
	Fields []Field   `json:"fields,omitempty"`
	Args   []any     `json:"args"`
}
