// Copyright 2014 The goyy Authors.  All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"path/filepath"

	enti "gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/util/files"
	"gopkg.in/goyy/goyy.v0/util/strings"
)

// factory is a file generation factory.
type factory struct {
	Project           string
	PackageName       string
	EntiPath          string
	CliPath           string
	APIPath           string
	TstPath           string
	HasGenProj        bool
	HasGenEntity      bool
	HasGenService     bool
	HasGenController  bool
	HasGenDto         bool
	HasGenAPI         bool
	HasGenSQL         bool
	HasGenLog         bool
	HasGenUtil        bool
	HasGenConst       bool
	HasGenHTML        bool
	HasGenJs          bool
	IsTimeField       bool
	IsValidationField bool
	IsExtend          bool
	NewProjName       string // Name of new project
	NewProjPath       string // Path of new project
	Entities          []*entity
	SysColumns        []string // goyy>data>entity:SysColumns
	SysFields         []string // goyy>data>entity:SysFields
	TreeColumns       []string // goyy>data>entity:TreeColumns
	TreeFields        []string // goyy>data>entity:TreeFields
}

// Init initializes an File from a path.
func (me *factory) Init(path string) error {
	// Set up the inheritance list of column names
	me.SysColumns = enti.SysColumns[:]
	me.SysFields = enti.SysFields[:]
	me.TreeColumns = enti.TreeColumns[:]
	me.TreeFields = enti.TreeFields[:]

	// set the path
	if strings.HasSuffix(path, ".go") {
		me.EntiPath = strings.TrimSuffix(path, ".go")
	} else {
		return fmt.Errorf("File '%s' is not a Go file.", path)
	}

	f, err := parser.ParseFile(
		token.NewFileSet(),
		path,
		nil,
		parser.ParseComments,
	)
	if err != nil {
		return fmt.Errorf("Unable to parse '%s': %s", path, err)
	}

	// get package name
	if f.Name != nil {
		me.PackageName = f.Name.Name
	} else {
		return fmt.Errorf("Missing package name in '%s'", path)
	}

	// build list of entities
	var isEntity bool
	for _, decl := range f.Decls {

		// get the type declaration
		tdecl, ok := decl.(*ast.GenDecl)
		if !ok || tdecl.Doc == nil {
			continue
		}

		// find the @entity decorator
		isEntity = false
		project := ""
		module := ""
		extend := ""
		relationship := ""
		for _, comment := range tdecl.Doc.List {
			if strings.Contains(comment.Text, "@entity") {
				isEntity = true
				// get entity.Project and entity.Extend
				c := strings.Between(comment.Text, "@entity(", ")")
				if strings.IsNotBlank(c) {
					val := convertUTF8(c)
					if strings.IsNotBlank(val) {
						project = tagItemValue(val, "project")
						module = tagItemValue(val, "module")
						relationship = tagItemValue(val, "relationship")
						if relationship != "slave" {
							relationship = "master"
						}
					}
				}
				break
			}
		}
		if !isEntity {
			continue
		}

		e := &entity{Project: project, Module: module, Relationship: relationship}
		if strings.IsBlank(me.Project) {
			me.Project = project
		}

		// get the name of the entity
		for _, spec := range tdecl.Specs {
			if ts, ok := spec.(*ast.TypeSpec); ok {
				if ts.Name == nil {
					continue
				}
				e.Name = ts.Name.Name
				break
			}
		}
		if e.Name == "" {
			return fmt.Errorf("Unable to extract name from a entity struct.")
		}

		sdecl := tdecl.Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
		fields := sdecl.Fields.List
		for _, f := range fields {
			typ := me.printerType(f.Type)
			if typ == "pk" || typ == "sys" || typ == "tree" {
				extend = typ
				e.Extend = extend
				if "pk" == extend {
					e.AllColumnMaxLen = 2
					e.AllFieldMaxLen = 2
					e.AllTypeMaxLen = 6
				} else if "sys" == extend {
					e.AllColumnMaxLen = 9
					e.AllFieldMaxLen = 9
					e.AllTypeMaxLen = 6
				} else if "tree" == extend {
					e.AllColumnMaxLen = 12
					e.AllFieldMaxLen = 11
					e.AllTypeMaxLen = 6
				}
				switch extend {
				case "pk", "sys", "tree":
					col := &field{Name: "id", Type: "string", Column: "id", IsPrimary: true}
					e.PrimaryKeys = append(e.PrimaryKeys, col)
				}
				break
			}
		}
		// parse the xgen tag and build columns
		for _, f := range fields {
			typ := me.printerType(f.Type)

			if typ == "err" || typ == "pk" || typ == "sys" || typ == "tree" {
				continue
			}

			var items string

			if f.Tag == nil || strings.IsBlank(f.Tag.Value) {
				if typ == "table" {
					items = fmt.Sprintf("table=%s", strings.UnCamel(e.Name, "_"))
				} else {
					items = fmt.Sprintf("column=%s", strings.UnCamel(f.Names[0].Name, "_"))
				}
			} else {
				items = tagItemValue(f.Tag.Value, "db")
			}

			if strings.IsNotBlank(items) {
				if typ == "table" {
					// parse attributes
					attributes := strings.Split(items, "&")
					for _, attr := range attributes {
						pair := strings.Split(attr, "=")
						if len(pair) != 2 {
							return fmt.Errorf("Malformed tag: '%s'", attr)
						}
						switch strings.ToLower(pair[0]) {
						case "table":
							e.Table = pair[1]
						case "comment":
							e.Comment = pair[1]
						}
					}
					continue
				}

				col := &field{}
				if err := col.Init(f.Names[0].Name, typ, items); err != nil {
					return fmt.Errorf(
						"Unable to parse tag '%s' from entity '%s' in '%s': %v",
						items,
						e.Name,
						path,
						err,
					)
				}
				// validation init
				if f.Tag != nil && strings.IsNotBlank(f.Tag.Value) {
					items = tagItemValue(f.Tag.Value, "validation")
					if strings.IsNotBlank(items) {
						if v, ok := newValidations(items); ok {
							col.Validations = v
						} else {
							return fmt.Errorf(
								"Unable to parse tag '%s' from entity '%s' in '%s': %v",
								items,
								e.Name,
								path,
								err,
							)
						}
					}
				}
				// excel init
				if f.Tag != nil && strings.IsNotBlank(f.Tag.Value) {
					items = tagItemValue(f.Tag.Value, "excel")
					if strings.IsNotBlank(items) {
						if v, ok := newExcelField(col, items); ok {
							col.Excel = v
							col.IsExcel = true
						} else {
							return fmt.Errorf(
								"Unable to parse tag '%s' from entity '%s' in '%s': %v",
								items,
								e.Name,
								path,
								err,
							)
						}
					}
				}
				// json init
				if f.Tag != nil && strings.IsNotBlank(f.Tag.Value) {
					items = tagItemValue(f.Tag.Value, "json")
					if strings.IsBlank(items) {
						items = col.Name
					}
					if v, ok := newJSONField(col, items); ok {
						col.JSON = v
						col.IsJSON = true
					} else {
						return fmt.Errorf(
							"Unable to parse tag '%s' from entity '%s' in '%s': %v",
							items,
							e.Name,
							path,
							err,
						)
					}
				}
				// xml init
				if f.Tag != nil && strings.IsNotBlank(f.Tag.Value) {
					items = tagItemValue(f.Tag.Value, "xml")
					if strings.IsBlank(items) {
						items = col.Name
					}
					if v, ok := newXMLField(col, items); ok {
						col.XML = v
						col.IsXML = true
					} else {
						return fmt.Errorf(
							"Unable to parse tag '%s' from entity '%s' in '%s': %v",
							items,
							e.Name,
							path,
							err,
						)
					}
				}
				e.Fields = append(e.Fields, col)
				if col.IsPrimary {
					e.PrimaryKeys = append(e.PrimaryKeys, col)
				}
				if len(col.Name) > e.FieldMaxLen {
					e.FieldMaxLen = len(col.Name)
				}
				if len(col.Name) > e.AllFieldMaxLen {
					e.AllFieldMaxLen = len(col.Name)
				}
				if len(col.Column) > e.ColumnMaxLen {
					e.ColumnMaxLen = len(col.Column)
				}
				if len(col.Column) > e.AllColumnMaxLen {
					e.AllColumnMaxLen = len(col.Column)
				}
				if len(col.Type) > e.TypeMaxLen {
					e.TypeMaxLen = len(col.Type)
				}
				if len(col.Type) > e.AllTypeMaxLen {
					e.AllTypeMaxLen = len(col.Type)
				}
			}
		}
		if len(e.PrimaryKeys) > 0 {
			me.Entities = append(me.Entities, e)
		}
	}

	me.isTimeField()
	me.isValidationField()
	me.isExtend()

	return nil
}

func (me factory) Write() error {
	if me.HasGenProj {
		if err := me.writeNewProj(); err != nil {
			return err
		}
	}
	if me.HasGenEntity {
		if err := me.writeEntityXgen(); err != nil {
			return err
		}
		if err := me.writeEntitiesXgen(); err != nil {
			return err
		}
	}
	if me.HasGenService {
		if err := me.writeServiceXgen(); err != nil {
			return err
		}
		if err := me.writeServiceMain(); err != nil {
			return err
		}
		if strings.IsNotBlank(me.TstPath) {
			if err := me.writeServiceTest(); err != nil {
				return err
			}
		}
	}
	if me.HasGenController {
		if strings.IsNotBlank(me.APIPath) {
			if err := me.writeControllerXgen(); err != nil {
				return err
			}
			if err := me.writeControllerMain(); err != nil {
				return err
			}
			if strings.IsNotBlank(me.TstPath) {
				if err := me.writeControllerTest(); err != nil {
					return err
				}
			}
			if err := me.writeControllerReg(); err != nil {
				return err
			}
		}
	}
	if me.HasGenDto {
		if err := me.writeDtoXgen(); err != nil {
			return err
		}
	}
	if me.HasGenAPI {
		if strings.IsNotBlank(me.APIPath) {
			if err := me.writeAPIMain(); err != nil {
				return err
			}
		}
	}
	if me.HasGenSQL {
		if err := me.writeSQLMain(); err != nil {
			return err
		}
	}
	if me.HasGenLog {
		if err := me.writeLogJSONXgen(); err != nil {
			return err
		}
		if strings.IsNotBlank(me.APIPath) {
			if err := me.writeLogAPIXgen(); err != nil {
				return err
			}
		}
	}
	if me.HasGenUtil {
		if err := me.writeUtilMain(); err != nil {
			return err
		}
	}
	if me.HasGenConst {
		if err := me.writeConstMain(); err != nil {
			return err
		}
	}
	if me.HasGenHTML {
		if err := me.writeHTMLMain(); err != nil {
			return err
		}
	}
	if me.HasGenJs {
		if err := me.writeJsMain(); err != nil {
			return err
		}
	}
	return nil
}

func (me factory) writeBy(typ, content string) error {
	var dir, dstfile string
	if me.HasGenEntity {
		clidir := "../../../" + strings.AfterLast(me.CliPath, "/")
		// get the destination file
		dir, file := filepath.Split(me.EntiPath)
		f, name := me.genFileName(typ, file)
		switch typ {
		case xgenDto:
			dir = clidir + "/internal/" + me.Project + "/" + me.PackageName
		case mainAPI, xgenLogAPI:
			dir = "../../api/" + me.PackageName
		case mainHTML:
			dir = clidir + "/templates/" + me.Project + "/" + name
		case mainJs:
			dir = clidir + "/static/js/" + me.Project + "/" + name
		case xgenCtlReg:
			dir = "../../"
		}
		dstfile = filepath.Join(dir, f)
	}
	if me.HasGenProj {
		switch typ {
		case newProj + ".README":
			dir = "."
			dstfile = "README.md"
		case newProj + ".tst":
			dir = me.NewProjPath + "/" + me.NewProjName + "-tst/"
			dstfile = "tst.go"
		case newProj + ".tst.settings":
			dir = me.NewProjPath + "/" + me.NewProjName + "-tst/conf/env/"
			dstfile = "settings.xml"
		case newProj + ".tst.session":
			dir = me.NewProjPath + "/" + me.NewProjName + "-tst/conf/env/"
			dstfile = "session.xml"
		case newProj + ".tst.db":
			dir = me.NewProjPath + "/" + me.NewProjName + "-tst/conf/env/"
			dstfile = "db.xml"
		case newProj + ".schema.bin.sql.sh":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-sql.sh"
		case newProj + ".schema.bin.sql.bat":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-sql.bat"
		case newProj + ".schema.bin.menu.sh":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-menu.sh"
		case newProj + ".schema.bin.menu.bat":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-menu.bat"
		case newProj + ".schema.bin.entity.sh":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-entity.sh"
		case newProj + ".schema.bin.entity.bat":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/bin/"
			dstfile = "exp-entity.bat"
		case newProj + ".schema.conf.db":
			dir = me.NewProjPath + "/" + me.NewProjName + "-schema/conf/env/"
			dstfile = "db.xml"
		}
		dstfile = filepath.Join(dir, dstfile)
	}
	if files.IsExist(dstfile) {
		if strings.HasPrefix(typ, typXgen) {
			files.Remove(dstfile)
		} else {
			return nil
		}
	} else {
		files.MkdirAll(dir, 0744)
	}

	buf := bytes.Buffer{}
	tmpl := newTmpl(content)
	tmpl.Execute(&buf, me)
	return ioutil.WriteFile(dstfile, buf.Bytes(), 0744)
}

func (me factory) writeNewProj() error {
	if err := me.writeBy(newProj+".README", tmplNewProjReadme); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".tst", tmplNewProjTst); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".tst.settings", tmplNewProjTstSettings); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".tst.session", tmplNewProjTstSession); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".tst.db", tmplNewProjTstDB); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.sql.sh", tmplNewProjSchemaBinSqlSh); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.sql.bat", tmplNewProjSchemaBinSqlBat); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.menu.sh", tmplNewProjSchemaBinMenuSh); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.menu.bat", tmplNewProjSchemaBinMenuBat); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.entity.sh", tmplNewProjSchemaBinEntitySh); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.bin.entity.bat", tmplNewProjSchemaBinEntityBat); err != nil {
		return err
	}
	if err := me.writeBy(newProj+".schema.conf.db", tmplNewProjTstDB); err != nil {
		return err
	}
	return nil
}

func (me factory) writeEntityXgen() error {
	return me.writeBy(xgenEntity, tmplEntity)
}

func (me factory) writeEntitiesXgen() error {
	return me.writeBy(xgenEntities, tmplEntities)
}

func (me factory) writeServiceXgen() error {
	return me.writeBy(xgenService, tmplServiceXgen)
}

func (me factory) writeServiceMain() error {
	return me.writeBy(mainService, tmplServiceMain)
}

func (me factory) writeServiceTest() error {
	return me.writeBy(testService, tmplServiceTest)
}

func (me factory) writeControllerXgen() error {
	return me.writeBy(xgenCtl, tmplControllerXgen)
}

func (me factory) writeControllerMain() error {
	return me.writeBy(mainCtl, tmplControllerMain)
}

func (me factory) writeControllerTest() error {
	return me.writeBy(testCtl, tmplControllerTest)
}

func (me factory) writeControllerReg() error {
	return me.writeBy(xgenCtlReg, tmplControllerReg)
}

func (me factory) writeSQLMain() error {
	return me.writeBy(mainSQL, tmplSQLMain)
}

func (me factory) writeLogJSONXgen() error {
	return me.writeBy(xgenLogJSON, tmplLogXgen)
}

func (me factory) writeLogAPIXgen() error {
	return me.writeBy(xgenLogAPI, tmplLogXgen)
}

func (me factory) writeUtilMain() error {
	return me.writeBy(mainUtil, tmplUtilMain)
}

func (me factory) writeConstMain() error {
	return me.writeBy(mainConst, tmplConstMain)
}

func (me factory) writeHTMLMain() error {
	return me.writeBy(mainHTML, tmplHTMLMain)
}

func (me factory) writeJsMain() error {
	return me.writeBy(mainJs, tmplJsMain)
}

func (me factory) writeDtoXgen() error {
	if strings.IsBlank(me.CliPath) {
		return nil
	}
	return me.writeBy(xgenDto, tmplDtoXgen)
}

func (me factory) writeAPIMain() error {
	return me.writeBy(mainAPI, tmplAPIMain)
}

func (me factory) genFileName(typ, name string) (string, string) {
	switch typ {
	case xgenLogJSON, xgenLogAPI:
		return "log_xgen.go", name
	case mainAPI:
		return me.PackageName + ".go", name
	case mainUtil:
		return me.PackageName + "_util.go", name
	case mainConst:
		return me.PackageName + "_const.go", name
	case xgenCtlReg:
		return me.PackageName + "_register_xgen.go", name
	}
	if strings.HasPrefix(name, typMain) {
		name = strings.After(name, typMain)
	}
	if strings.HasSuffix(name, "_entity") {
		name = strings.Before(name, "_entity")
	}
	if name == "domain" || name == "entity" || name == "main.domain" || name == "main.entity" {
		name = ""
	} else {
		if strings.IsNotBlank(name) && typ != mainHTML && typ != mainJs {
			name = name + "_"
		}
	}
	switch typ {
	case mainCtl:
		typ = "controller"
	case mainService:
		typ = "manager"
	case mainSQL:
		typ = "sql"
	case xgenCtl:
		typ, name = me.resetTypAndName("controller_xgen", name)
	case xgenService:
		typ, name = me.resetTypAndName("manager_xgen", name)
	case xgenEntity:
		typ, name = me.resetTypAndName("entity_xgen", name)
	case xgenEntities:
		typ, name = me.resetTypAndName("entities_xgen", name)
	case xgenDto:
		typ, name = me.resetTypAndName("dto_xgen", name)
	}
	switch typ {
	case mainHTML:
		return fmt.Sprintf("%s.html", name), name
	case mainJs:
		return fmt.Sprintf("%s.js", name), name
	case testCtl:
		return fmt.Sprintf("%scontroller_test.go", name), name
	case testService:
		return fmt.Sprintf("%smanager_test.go", name), name
	default:
		return fmt.Sprintf("%s%s.go", name, typ), name
	}
}

func (me factory) resetTypAndName(typ, name string) (string, string) {
	if strings.HasSuffix(name, "_test_") {
		name = strings.Before(name, "test_")
		typ = typ + "_test"
	}
	return typ, name
}

func (me *factory) isTimeField() {
	for _, e := range me.Entities {
		for _, f := range e.Fields {
			if f.Type == "time.Time" {
				me.IsTimeField = true
			}
		}
	}
}

func (me *factory) isValidationField() {
	for _, e := range me.Entities {
		for _, f := range e.Fields {
			if len(f.Validations) > 0 {
				me.IsValidationField = true
			}
		}
	}
}

func (me *factory) isExtend() {
	for _, e := range me.Entities {
		if strings.IsNotBlank(e.Extend) {
			me.IsExtend = true
		}
	}
}

func (me factory) printerType(e ast.Expr) string {
	var b bytes.Buffer
	printer.Fprint(&b, token.NewFileSet(), e)
	switch b.String() {
	case "schema.Table":
		return "table"
	case "entity.String":
		return "string"
	case "entity.Bool":
		return "bool"
	case "entity.Float32":
		return "float32"
	case "entity.Float64":
		return "float64"
	case "entity.Int":
		return "int"
	case "entity.Int8":
		return "int8"
	case "entity.Int16":
		return "int16"
	case "entity.Int32":
		return "int32"
	case "entity.Int64":
		return "int64"
	case "entity.Time":
		return "time.Time"
	case "entity.Uint":
		return "uint"
	case "entity.Uint8":
		return "uint8"
	case "entity.Uint16":
		return "uint16"
	case "entity.Uint32":
		return "uint32"
	case "entity.Uint64":
		return "uint64"
	case "entity.Pk":
		return "pk"
	case "entity.Sys":
		return "sys"
	case "entity.Tree":
		return "tree"
	default:
		return "err"
	}
}
