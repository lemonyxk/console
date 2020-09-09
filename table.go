/**
* @program: lemo
*
* @description:
*
* @author: lemo
*
* @create: 2019-11-27 20:39
**/

package console

import (
	list2 "github.com/jedib0t/go-pretty/list"
	tab "github.com/jedib0t/go-pretty/table"
)

type SortMode int

const (
	// Asc sorts the column in Ascending order alphabetically.
	Asc SortMode = iota
	// AscNumeric sorts the column in Ascending order numerically.
	AscNumeric
	// Dsc sorts the column in Descending order alphabetically.
	Dsc
	// DscNumeric sorts the column in Descending order numerically.
	DscNumeric
)

type table struct {
	write tab.Writer
}

type list struct {
	write list2.Writer
}

func NewList() *list {
	return &list{list2.NewWriter()}
}

func (l *list) AppendItem(v interface{}) *list {
	l.write.AppendItem(v)
	return l
}

func (l *list) AppendItems(v ...interface{}) *list {
	l.write.AppendItems(v)
	return l
}

func (l *list) Render() string {
	return l.write.Render()
}

func (l *list) Indent() *list {
	l.write.Indent()
	return l
}

func (l *list) UnIndent() *list {
	l.write.UnIndent()
	return l
}

func (l *list) Length() int {
	return l.write.Length()
}

func NewTable() *table {
	return &table{tab.NewWriter()}
}

func (t *table) Header(v ...interface{}) *table {
	t.write.AppendHeader(v)
	return t
}

func (t *table) SetPageSize(limit int) *table {
	t.write.SetPageSize(limit)
	return t
}

func (t *table) Style() *tab.Style {
	return t.write.Style()
}

func (t *table) Row(v ...interface{}) *table {
	t.write.AppendRow(v)
	return t
}

func (t *table) Footer(v ...interface{}) *table {
	t.write.AppendFooter(v)
	return t
}

func (t *table) SortByName(name string, mode SortMode) *table {
	t.write.SortBy([]tab.SortBy{{Name: name, Mode: tab.SortMode(mode)}})
	return t
}

func (t *table) SortByNumber(number int, mode SortMode) *table {
	t.write.SortBy([]tab.SortBy{{Number: number, Mode: tab.SortMode(mode)}})
	return t
}

func (t *table) Title(format string, v ...interface{}) *table {
	t.write.SetTitle(format, v...)
	return t
}

func (t *table) Render() string {
	return t.write.Render()
}
