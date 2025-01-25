/**
* @program: console
*
* @create: 2025-01-25 15:51
**/

package test

import (
	"github.com/lemonyxk/console"
	"github.com/rs/zerolog"
	"io"
	"testing"
)

func BenchmarkLog(b *testing.B) {
	var defaultLogger = zerolog.New(io.Discard).With().Caller().Logger()
	var Info = console.NewLogger(&defaultLogger, zerolog.InfoLevel)
	for i := 0; i < b.N; i++ {
		Info.Log("hello")
	}
}
