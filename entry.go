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
	Level  Level
	File   string
	Line   int
	Time   time.Time
	Format string
	Args   []interface{}
}
