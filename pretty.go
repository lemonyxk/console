/**
* @program: go
*
* @description:
*
* @author: lemo
*
* @create: 2020-10-01 01:13
**/

package console

import (
	"fmt"
	"reflect"
	"strings"
	"sync"
	"text/tabwriter"
)

var mux sync.Mutex

var visited = make(map[uintptr]bool)
var tw = tabwriter.NewWriter(writer, 0, 0, tabWidth, ' ', 0)

var deep = 0
var isComplex = false

var public = true
var private = true
var details = false

var tabWidth = 4

type pretty int

var Pretty pretty

func (p pretty) Dump(v ...interface{}) {
	mux.Lock()
	for i := 0; i < len(v); i++ {
		dump(reflect.ValueOf(v[i]))
	}
	mux.Unlock()
}

func (p pretty) Public(v ...interface{}) {
	mux.Lock()
	public = true
	private = false
	for i := 0; i < len(v); i++ {
		dump(reflect.ValueOf(v[i]))
	}
	mux.Unlock()
}

func (p pretty) Private(v ...interface{}) {
	mux.Lock()
	public = false
	private = true
	for i := 0; i < len(v); i++ {
		dump(reflect.ValueOf(v[i]))
	}
	mux.Unlock()
}

func (p pretty) Details(v ...interface{}) {
	mux.Lock()
	details = true
	for i := 0; i < len(v); i++ {
		dump(reflect.ValueOf(v[i]))
	}
	mux.Unlock()
}

func dump(v reflect.Value) {
	format(v)
	reset()
}

func reset() {

	public = true
	private = true
	details = false

	deep = 0
	isComplex = false
	visited = make(map[uintptr]bool)
	panicIfError(tw.Flush())
}

func writeString(s string) {
	panicIfError(fmt.Fprintf(tw, "%s", s))
}

func writeValue(s string) {
	if isComplex {
		s += ","
	}
	panicIfError(fmt.Fprintf(tw, "%s\n", s))
}

func writeKey(s string) {
	panicIfError(fmt.Fprintf(tw, "%s", strings.Repeat(" ", deep*tabWidth)+s))
}

func writeStart(s string) {
	panicIfError(fmt.Fprintf(tw, "%s\n", s))
}

func writeEnd(deep int, s string) {
	if deep > 0 {
		s += ","
	}
	panicIfError(fmt.Fprintf(tw, "%s\n", strings.Repeat(" ", deep*tabWidth)+s))
}

func format(rv reflect.Value) {

	// if is error or Stringer
	// config details
	if !details && rv.IsValid() {

		if rv.Kind() == reflect.Interface {
			format(rv.Elem())
			return
		}

		// has Error method
		var m = rv.MethodByName("Error")
		if m.IsValid() && m.String() == `<func() string Value>` {
			writeValue(Bold.Mixed(FgGreen).Sprint(`"` + m.Call(nil)[0].String() + `"`))
			return
		}

		// has String method
		m = rv.MethodByName("String")
		if m.IsValid() && m.String() == `<func() string Value>` {
			writeValue(Bold.Mixed(FgGreen).Sprint(`"` + m.Call(nil)[0].String() + `"`))
			return
		}

	}

	switch rv.Kind() {

	// SIMPLE TYPE
	case reflect.Bool:
		writeValue(Bold.Mixed(FgCyan).Sprint(simple(rv)))
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64,
		reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Uintptr, reflect.Complex64, reflect.Complex128:
		writeValue(Bold.Mixed(FgBlue).Sprint(simple(rv)))
	case reflect.Float32, reflect.Float64:
		writeValue(Bold.Mixed(FgBlue).Sprint(simple(rv)))
	case reflect.String:
		writeValue(Bold.Mixed(FgGreen).Sprint(simple(rv)))
	case reflect.Func:
		printFunc(rv)
	case reflect.UnsafePointer:
		printUnsafePointer(rv)
	case reflect.Chan:
		printChan(rv)
	case reflect.Invalid:
		writeValue(Bold.Mixed(FgHiRed).Sprint("nil"))

	// COMPLEX TYPE
	case reflect.Map:
		printMap(rv)
	case reflect.Struct:
		printStruct(rv)
	case reflect.Array, reflect.Slice:
		printSlice(rv)
	case reflect.Ptr:
		printPtr(rv)
	case reflect.Interface:
		format(rv.Elem())
	default:
		writeValue(simple(rv))
	}
}

func printFunc(v reflect.Value) {
	if !v.IsNil() {
		writeValue(Bold.Mixed(FgYellow).Sprintf("%s {...}", typeString(v)))
	} else {
		writeValue(Bold.Mixed(FgYellow).Sprintf("(%s {...}) (nil)", typeString(v)))
	}
}

func printChan(v reflect.Value) {
	if !v.IsNil() {
		writeValue(Bold.Mixed(FgMagenta).Sprintf("(%s) (%s)", typeString(v), addrString(v)))
	} else {
		writeValue(Bold.Mixed(FgMagenta).Sprintf("(%s) (nil)", typeString(v)))
	}
}

func printUnsafePointer(v reflect.Value) {
	writeValue(Bold.Sprintf("(%s) (%s)", typeString(v), addrString(v)))
}

func printMap(v reflect.Value) {
	isComplex = true

	var d = deep
	deep++

	if v.Len() == 0 {
		if !v.IsNil() {
			writeValue(Bold.Sprintf("%s{}", typeString(v)))
		} else {
			writeValue(Bold.Sprintf("(%s) (nil)", typeString(v)))
		}
		deep = d
		return
	}

	if visited[v.Pointer()] {
		writeValue(Bold.Sprintf("%s{...}", typeString(v)))
		deep = d
		return
	}

	visited[v.Pointer()] = true

	writeStart(Bold.Mixed(FgRed).Sprint(typeString(v)) + "{")

	keys := v.MapKeys()
	for i := 0; i < v.Len(); i++ {
		value := v.MapIndex(keys[i])
		writeKey(Bold.Sprint(simple(keys[i]) + ":" + " "))
		format(value)
	}

	writeEnd(d, Bold.Sprint("}"))

	deep = d
}

func printSlice(v reflect.Value) {

	isComplex = true

	var d = deep
	deep++

	if v.Len() == 0 {
		if !v.IsNil() {
			writeValue(Bold.Sprintf("%s{}", typeString(v)))
		} else {
			writeValue(Bold.Sprintf("(%s) (nil)", typeString(v)))
		}
		deep = d
		return
	}

	//  if is array, will be handled in printPtr
	if v.Kind() == reflect.Slice {
		if visited[v.Pointer()] {
			writeValue(Bold.Sprintf("%s{...}", typeString(v)))
			deep = d
			return
		}
		visited[v.Pointer()] = true
	}

	writeStart(Bold.Mixed(FgRed).Sprint(typeString(v)) + "{")

	for i := 0; i < v.Len(); i++ {
		value := v.Index(i)
		writeKey(" ")
		format(value)
	}

	writeEnd(d, Bold.Sprint("}"))

	deep = d
}

func printStruct(v reflect.Value) {
	isComplex = true

	var d = deep
	deep++

	writeStart(Bold.Mixed(FgRed).Sprint(typeString(v)) + "{")

	for i := 0; i < v.NumField(); i++ {
		field := v.Type().Field(i).Name
		value := v.Field(i)

		// if is private
		// config private & public
		if value.CanInterface() {
			if public {
				writeKey(Bold.Sprint(field + ":" + " "))
				format(value)
			}
		} else {
			if private {
				writeKey(Bold.Sprint(field + ":" + " "))
				format(value)
			}
		}
	}

	writeEnd(d, Bold.Sprint("}"))

	deep = d
}

func printPtr(v reflect.Value) {

	if visited[v.Pointer()] {
		writeValue(Bold.Sprintf("&%s{...}", elemTypeString(v)))
		return
	}

	if v.Pointer() != 0 {
		visited[v.Pointer()] = true
	}

	if v.Elem().IsValid() {
		writeString(Bold.Sprint("&"))
		format(v.Elem())
	} else {
		writeValue(Bold.Sprintf("(%s) (nil)", typeString(v)))
	}
}

func simple(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Bool:
		return fmt.Sprintf("%#v", v.Bool())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return fmt.Sprintf("%#v", v.Int())
	case reflect.Uint, reflect.Uintptr:
		return fmt.Sprintf("%#v", v.Uint())
	case reflect.Uint8:
		return fmt.Sprintf("0x%02x", v.Uint())
	case reflect.Uint16:
		return fmt.Sprintf("0x%04x", v.Uint())
	case reflect.Uint32:
		return fmt.Sprintf("0x%08x", v.Uint())
	case reflect.Uint64:
		return fmt.Sprintf("0x%016x", v.Uint())
	case reflect.Float32, reflect.Float64:
		return fmt.Sprintf("%f", v.Float())
	case reflect.Complex64, reflect.Complex128:
		return fmt.Sprintf("%#v", v.Complex())
	case reflect.String:
		return fmt.Sprintf("%#v", v.String())
	case reflect.Invalid:
		return "nil"
	default:
		return fmt.Sprintf("%#v", v.Interface())
	}
}

func typeString(v reflect.Value) string {
	return v.Type().String()
}

func elemTypeString(v reflect.Value) string {
	return v.Elem().Type().String()
}

func addrString(v reflect.Value) string {
	return fmt.Sprintf("%#v", v.Pointer())
}
