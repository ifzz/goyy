// generated by xgen -- DO NOT EDIT
package post

import (
	"gopkg.in/goyy/goyy.v0/data/entity"
	"gopkg.in/goyy/goyy.v0/data/service"
)

var MenuMgr = &MenuManager{
	Manager: service.Manager{
		Entity: func() entity.Interface {
			return NewMenuEntity()
		},
		Entities: func() entity.Interfaces {
			return NewMenuEntities(10)
		},
	},
}

type MenuManager struct {
	service.Manager
}
