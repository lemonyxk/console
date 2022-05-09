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

type Entry struct {
	Level  Level          `json:"level,omitempty"`
	File   string         `json:"file,omitempty"`
	Line   int            `json:"line,omitempty"`
	Time   time.Time      `json:"time,omitempty"`
	Format string         `json:"format"`
	Fields map[string]any `json:"fields,omitempty"`
	Args   []any          `json:"args"`
}
