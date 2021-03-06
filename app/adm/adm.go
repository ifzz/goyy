package main

import _ "github.com/go-sql-driver/mysql"
import _ "gopkg.in/goyy/goyy.v0/web/conf"
import _ "gopkg.in/goyy/goyy.v0/app/sys/api/dict"

import (
	_ "gopkg.in/goyy/goyy.v0/app/adm/internal"
	_ "gopkg.in/goyy/goyy.v0/app/comm/controller/assets"
	_ "gopkg.in/goyy/goyy.v0/app/comm/controller/captcha"
	_ "gopkg.in/goyy/goyy.v0/app/comm/controller/home"
	_ "gopkg.in/goyy/goyy.v0/app/comm/controller/login"
	_ "gopkg.in/goyy/goyy.v0/app/example"
	_ "gopkg.in/goyy/goyy.v0/app/sys"
	"gopkg.in/goyy/goyy.v0/comm/log"
	"gopkg.in/goyy/goyy.v0/web/xhttp"
)

func main() {
	xhttp.Conf.Addr = ":9097"
	xhttp.Use(xhttp.Recovery())
	err := xhttp.Run()
	if err != nil {
		log.Error(err.Error())
	}
}
