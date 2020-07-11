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

type Level int

const (
	ERR Level = iota + 1
	INF
	DEB
	WAR
)

func (level Level) String() string {
	switch level {
	case ERR:
		return "ERR"
	case INF:
		return "INF"
	case DEB:
		return "DEB"
	case WAR:
		return "WAR"
	}
	panic("invalid level")
}

type Entry struct {
	Level  Level
	File   string
	Line   int
	Time   time.Time
	Args   []interface{}
	Format string
}
