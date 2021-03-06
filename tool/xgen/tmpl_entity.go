// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

var tmplEntity = `// generated by xgen -- DO NOT EDIT
package <%.PackageName%>
<%$ := .%>
import (
	"bytes"<%if .IsTimeField%>
	"time"<%end%>

	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/schema"<%if .IsValidationField%>
	"gopkg.in/goyy/goyy.v0/data/validate"<%end%>
	"gopkg.in/goyy/goyy.v0/util/jsons"
	"gopkg.in/goyy/goyy.v0/util/strings"
)<%range $e := .Entities%>

var (
	<%tname $e.Name $e.AllColumnMaxLen%> = schema.TABLE("<%lower $e.Table%>", "<%$e.GetComment%>")<%if eq $e.Extend "pk"%><%if not (existcol $e "id")%>
	<%uncamel $e.Name|upper%>_<%cname "id" $e.AllColumnMaxLen%> = <%uncamel $e.Name|upper%>.PRIMARY("id", "<%message "col.comment.id"%>")<%end%><%else if eq $e.Extend "sys"%><%range $c := $.SysColumns%><%if not (existcol $e $c)%>
	<%uncamel $e.Name|upper%>_<%cname $c $e.AllColumnMaxLen%> = <%uncamel $e.Name|upper%>.<%coltype $c%>("<%$c%>", "<%message (printf "col.comment.%s" $c)%>")<%end%><%end%><%else if eq $e.Extend "tree"%><%range $c := $.TreeColumns%><%if not (existcol $e $c)%>
	<%uncamel $e.Name|upper%>_<%cname $c $e.AllColumnMaxLen%> = <%uncamel $e.Name|upper%>.<%coltype $c%>("<%$c%>", "<%message (printf "col.comment.%s" $c)%>")<%end%><%end%><%end%><%range $f := $e.Fields%>
	<%uncamel $e.Name|upper%>_<%cname $f.Column $e.AllColumnMaxLen%> = <%uncamel $e.Name|upper%>.<%coltypef $f%>("<%lower $f.Column%>", "<%$f.GetComment%>")<%end%>
)

func New<%$e.Name%>() *<%$e.Name%> {
	e := &<%$e.Name%>{}
	e.init()
	return e
}
<%range $f := $e.Fields%>
func (me *<%$e.Name%>) <%upper1 $f.Name%>() <%$f.Type%> {
	return me.<%$f.Name%>.Value()
}

func (me *<%$e.Name%>) Set<%upper1 $f.Name%>(v <%$f.Type%>) {
	me.<%$f.Name%>.SetValue(v)
}
<%end%>
func (me *<%$e.Name%>) init() {
	me.table = <%uncamel $e.Name|upper%>
	me.initSetDict()
	me.initSetColumn()
	me.initSetDefault()
	me.initSetField()
	me.initSetExcel()
	me.initSetJson()
	me.initSetXml()
}

func (me *<%$e.Name%>) initSetDict() {<%range $f := $e.Fields%><%if notblank $f.Dict%>
	<%uncamel $e.Name|upper%>_<%upper $f.Column%>.SetDict("<%$f.Dict%>")<%end%><%end%>
}

func (me *<%$e.Name%>) initSetColumn() {<%if eq $e.Extend "pk"%><%if not (existcol $e "id")%>
	if t, ok := me.Pk.Type("id"); ok {
		t.SetColumn(<%uncamel $e.Name|upper%>_<%upper "id"%>)
	}<%end%><%else if eq $e.Extend "sys"%><%range $c := $.SysColumns%><%if not (existcol $e $c)%>
	if t, ok := me.Sys.Type("<%$c%>"); ok {
		t.SetColumn(<%uncamel $e.Name|upper%>_<%upper $c%>)
	}<%end%><%end%><%else if eq $e.Extend "tree"%><%range $c := $.TreeColumns%><%if not (existcol $e $c)%>
	if t, ok := me.Tree.Type("<%$c%>"); ok {
		t.SetColumn(<%uncamel $e.Name|upper%>_<%upper $c%>)
	}<%end%><%end%><%end%><%range $f := $e.Fields%>
	me.<%$f.Name%>.SetColumn(<%uncamel $e.Name|upper%>_<%upper $f.Column%>)<%end%>
}

func (me *<%$e.Name%>) initSetDefault() {<%if eq $e.Extend "sys"%>
	if t, ok := me.Sys.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Sys.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}<%else if eq $e.Extend "tree"%>
	if t, ok := me.Tree.Type("created"); ok {
		t.SetDefault("-62135596800")
	}
	if t, ok := me.Tree.Type("modified"); ok {
		t.SetDefault("-62135596800")
	}<%end%><%range $f := $e.Fields%><%if notblank $f.Default%>
	me.<%$f.Name%>.SetDefault("<%$f.Default%>")<%end%><%end%>
}

func (me *<%$e.Name%>) initSetField() {<%if eq $e.Extend "pk"%>
	if t, ok := me.Pk.Type("id"); ok {
		t.SetField(entity.DefaultField())
	}<%else if eq $e.Extend "sys"%>
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}<%else if eq $e.Extend "tree"%>
	for _, c := range entity.TreeColumns {
		if t, ok := me.Tree.Type(c); ok {
			t.SetField(entity.DefaultField())
		}
	}<%end%><%range $f := $e.Fields%>
	me.<%$f.Name%>.SetField(entity.DefaultField())<%end%>
}

func (me *<%$e.Name%>) initSetExcel() {<%range $f := $e.Fields%><%if $f.IsExcel%>
	me.<%$f.Name%>.Field().SetExcel(entity.NewExcelBy("<%$f.Excel.Value%>", "<%$f.Excel.Title%>", "<%$f.Excel.Format%>", <%$f.Excel.Genre%>, <%$f.Excel.Align%>, <%$f.Excel.Sort%>, <%$f.Excel.Width%>))<%end%><%end%>
}

func (me *<%$e.Name%>) initSetJson() {<%if eq $e.Extend "pk"%>
	if t, ok := me.Pk.Type("id"); ok {
		t.Field().SetJson(entity.NewJsonBy("id"))
	}<%else if eq $e.Extend "sys"%>
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.Field().SetJson(entity.NewJsonBy(c))
		}
	}<%else if eq $e.Extend "tree"%>
	for _, c := range entity.TreeColumns {
		if t, ok := me.Tree.Type(c); ok {
			t.Field().SetJson(entity.NewJsonBy(c))
		}
	}<%end%><%range $f := $e.Fields%><%if $f.IsJSON%>
	me.<%$f.Name%>.Field().SetJson(entity.NewJsonBy("<%$f.JSON.Tag%>"))<%end%><%end%>
}

func (me *<%$e.Name%>) initSetXml() {<%if eq $e.Extend "pk"%>
	if t, ok := me.Pk.Type("id"); ok {
		t.Field().SetXml(entity.NewXmlBy("id"))
	}<%else if eq $e.Extend "sys"%>
	for _, c := range entity.SysColumns {
		if t, ok := me.Sys.Type(c); ok {
			t.Field().SetXml(entity.NewXmlBy(c))
		}
	}<%else if eq $e.Extend "tree"%>
	for _, c := range entity.TreeColumns {
		if t, ok := me.Tree.Type(c); ok {
			t.Field().SetXml(entity.NewXmlBy(c))
		}
	}<%end%><%range $f := $e.Fields%><%if $f.IsXML%>
	me.<%$f.Name%>.Field().SetXml(entity.NewXmlBy("<%$f.XML.Tag%>"))<%end%><%end%>
}

func (me <%$e.Name%>) New() entity.Interface {
	return New<%$e.Name%>()
}

func (me *<%$e.Name%>) Get(column string) interface{} {
	switch column {<%range $f := $e.Fields%>
	case <%uncamel $e.Name|upper%>_<%uncamel $f.Column|upper%>.Name():
		return me.<%$f.Name%>.Value()<%end%>
	}<%if eq $e.Extend "pk"%>
	return me.Pk.Get(column)<%else if eq $e.Extend "sys"%>
	return me.Sys.Get(column)<%else if eq $e.Extend "tree"%>
	return me.Tree.Get(column)<%else%>
	return nil<%end%>
}

func (me *<%$e.Name%>) GetPtr(column string) interface{} {
	switch column {<%range $f := $e.Fields%>
	case <%uncamel $e.Name|upper%>_<%uncamel $f.Column|upper%>.Name():
		return me.<%$f.Name%>.ValuePtr()<%end%>
	}<%if eq $e.Extend "pk"%>
	return me.Pk.GetPtr(column)<%else if eq $e.Extend "sys"%>
	return me.Sys.GetPtr(column)<%else if eq $e.Extend "tree"%>
	return me.Tree.GetPtr(column)<%else%>
	return nil<%end%>
}

func (me *<%$e.Name%>) GetString(field string) string {
	switch strings.ToLowerFirst(field) {<%range $f := $e.Fields%>
	case "<%$f.Name%>":
		return me.<%$f.Name%>.String()<%end%>
	}<%if eq $e.Extend "pk"%>
	return me.Pk.GetString(field)<%else if eq $e.Extend "sys"%>
	return me.Sys.GetString(field)<%else if eq $e.Extend "tree"%>
	return me.Tree.GetString(field)<%else%>
	return ""<%end%>
}

func (me *<%$e.Name%>) SetString(field, value string) error {
	switch strings.ToLowerFirst(field) {<%range $f := $e.Fields%>
	case "<%$f.Name%>":
		return me.<%$f.Name%>.SetString(value)<%end%>
	}<%if eq $e.Extend "pk"%>
	return me.Pk.SetString(field, value)<%else if eq $e.Extend "sys"%>
	return me.Sys.SetString(field, value)<%else if eq $e.Extend "tree"%>
	return me.Tree.SetString(field, value)<%else%>
	return nil<%end%>
}

func (me *<%$e.Name%>) Table() schema.Table {
	return me.table
}

func (me *<%$e.Name%>) Type(column string) (entity.Type, bool) {
	switch column {<%range $f := $e.Fields%>
	case <%uncamel $e.Name|upper%>_<%uncamel $f.Column|upper%>.Name():
		return &me.<%$f.Name%>, true<%end%>
	}<%if eq $e.Extend "pk"%>
	return me.Pk.Type(column)<%else if eq $e.Extend "sys"%>
	return me.Sys.Type(column)<%else if eq $e.Extend "tree"%>
	return me.Tree.Type(column)<%else%>
	return nil, false<%end%>
}

func (me *<%$e.Name%>) Column(field string) (schema.Column, bool) {
	switch strings.ToLowerFirst(field) {<%range $f := $e.Fields%>
	case "<%$f.Name%>":
		return <%uncamel $e.Name|upper%>_<%uncamel $f.Column|upper%>, true<%end%>
	}<%if eq $e.Extend "pk"%>
	return me.Pk.Column(field)<%else if eq $e.Extend "sys"%>
	return me.Sys.Column(field)<%else if eq $e.Extend "tree"%>
	return me.Tree.Column(field)<%else%>
	return nil, false<%end%>
}

func (me *<%$e.Name%>) Columns() []schema.Column {
	return []schema.Column{<%if eq $e.Extend "pk"%><%if not (existcol $e "id")%>
		<%uncamel $e.Name|upper%>_ID,<%end%><%else if eq $e.Extend "sys"%><%range $c := $.SysColumns%><%if not (existcol $e $c)%>
		<%uncamel $e.Name|upper%>_<%upper $c%>,<%end%><%end%><%else if eq $e.Extend "tree"%><%range $c := $.TreeColumns%><%if not (existcol $e $c)%>
		<%uncamel $e.Name|upper%>_<%upper $c%>,<%end%><%end%><%end%><%range $f := $e.Fields%>
		<%uncamel $e.Name|upper%>_<%uncamel $f.Column|upper%>,<%end%>
	}
}

func (me *<%$e.Name%>) Names() []string {
	return []string{<%if eq $e.Extend "pk"%><%if not (existcol $e "id")%>
		"id",<%end%><%else if eq $e.Extend "sys"%><%range $f := $.SysFields%><%if not (existfld $e $f)%>
		"<%$f%>",<%end%><%end%><%else if eq $e.Extend "tree"%><%range $f := $.TreeFields%><%if not (existfld $e $f)%>
		"<%$f%>",<%end%><%end%><%end%><%range $f := $e.Fields%>
		"<%$f.Name%>",<%end%>
	}
}

func (me *<%$e.Name%>) Value() *<%$e.Name%> {
	return me
}

func (me *<%$e.Name%>) Validate() error {<%if $.IsValidationField%>
	validation := &validate.Validation{}<%end%><%range $f := $e.Fields%><%range $v := $f.Validations%><%if eq $v.Name "required"%>
	if err := validation.Required("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "min"%>
	if err := validation.Min("<%$f.Name%>", me.<%$f.Name%>.String(), <%$v.Value%>); err != nil {
		return err
	}<%end%><%if eq $v.Name "max"%>
	if err := validation.Max("<%$f.Name%>", me.<%$f.Name%>.String(), <%$v.Value%>); err != nil {
		return err
	}<%end%><%if eq $v.Name "range"%>
	if err := validation.Range("<%$f.Name%>", me.<%$f.Name%>.String(), <%$v.Value%>); err != nil {
		return err
	}<%end%><%if eq $v.Name "minf"%>
	if err := validation.Minf("<%$f.Name%>", me.<%$f.Name%>.String(), <%$v.Value%>); err != nil {
		return err
	}<%end%><%if eq $v.Name "maxf"%>
	if err := validation.Maxf("<%$f.Name%>", me.<%$f.Name%>.String(), <%$v.Value%>); err != nil {
		return err
	}<%end%><%if eq $v.Name "rangef"%>
	if err := validation.Rangef("<%$f.Name%>", me.<%$f.Name%>.String(), <%$v.Value%>); err != nil {
		return err
	}<%end%><%if eq $v.Name "minlen"%>
	if err := validation.Minlen("<%$f.Name%>", me.<%$f.Name%>.String(), <%$v.Value%>); err != nil {
		return err
	}<%end%><%if eq $v.Name "maxlen"%>
	if err := validation.Maxlen("<%$f.Name%>", me.<%$f.Name%>.String(), <%$v.Value%>); err != nil {
		return err
	}<%end%><%if eq $v.Name "rangelen"%>
	if err := validation.Rangelen("<%$f.Name%>", me.<%$f.Name%>.String(), <%$v.Value%>); err != nil {
		return err
	}<%end%><%if eq $v.Name "email"%>
	if err := validation.Email("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "url"%>
	if err := validation.URL("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "ip"%>
	if err := validation.IP("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "mobile"%>
	if err := validation.Mobile("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "tel"%>
	if err := validation.Tel("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "phone"%>
	if err := validation.Phone("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "zipcode"%>
	if err := validation.Zipcode("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "float"%>
	if err := validation.Float("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "integer"%>
	if err := validation.Integer("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "alpha"%>
	if err := validation.Alpha("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "alrod"%>
	if err := validation.Alrod("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "alnum"%>
	if err := validation.Alnum("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "alnumrod"%>
	if err := validation.Alnumrod("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "alnumhan"%>
	if err := validation.Alnumhan("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "alnumhanrod"%>
	if err := validation.Alnumhanrod("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "alhan"%>
	if err := validation.Alhan("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "alhanrod"%>
	if err := validation.Alhanrod("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "han"%>
	if err := validation.Han("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "hanrod"%>
	if err := validation.Hanrod("<%$f.Name%>", me.<%$f.Name%>.String()); err != nil {
		return err
	}<%end%><%if eq $v.Name "match"%>
	if err := validation.Match("<%$f.Name%>", me.<%$f.Name%>.String(), ` + "`" + `<%$v.Value%>` + "`" + `); err != nil {
		return err
	}<%end%><%if eq $v.Name "nomatch"%>
	if err := validation.Nomatch("<%$f.Name%>", me.<%$f.Name%>.String(), ` + "`" + `<%$v.Value%>` + "`" + `); err != nil {
		return err
	}<%end%><%end%><%end%>
	return nil
}

func (me *<%$e.Name%>) JSON() string {
	var b bytes.Buffer
	b.WriteString("{")<%if eq $e.Extend "pk"%><%if not (existcol $e "id")%>
	b.WriteString(<%jsonm "id"%>)<%end%><%else if eq $e.Extend "sys"%><%range $f := $.SysFields%><%if not (existfld $e $f)%>
	b.WriteString(<%jsonm $f%>)<%end%><%end%><%else if eq $e.Extend "tree"%><%range $f := $.TreeFields%><%if not (existcol $e $f)%>
	b.WriteString(<%jsonm $f%>)<%end%><%end%><%end%><%range $f := $e.Fields%><%if not $f.JSON.Ignored%><%if $f.JSON.Omitempty%>
	if strings.IsNotBlank(me.<%$f.Name%>.String()) {
		b.WriteString(<%jsonmf $f%>)
	}<%else%>
	b.WriteString(<%jsonmf $f%>)<%end%><%end%><%end%>
	b.WriteString("}")
	return b.String()
}

func (me *<%$e.Name%>) ExcelColumns() []string {<%if $e.IsExcelColumns%>
	return []string{<%range $c := $e.GetExcelColumns%>
		"<%$c%>",<%end%>
	}<%else%>
	return nil<%end%>
}<%end%>
`
