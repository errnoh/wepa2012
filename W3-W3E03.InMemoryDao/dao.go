package main

import (
	"code.google.com/p/go-uuid/uuid"
)

type Object interface {
	GetID() string
	SetID(string)
}

type DAO interface {
	create(Object)
	read(string) Object
	update(Object)
	remove(string)
	list() []Object
}

//

type MapDAO map[string]Object

func NewMapDAO() (interface{}, error) {
	return make(MapDAO), nil
}

func (dao MapDAO) create(o Object) {
	uuid := uuid.NewRandom().String()
	o.SetID(uuid)
	dao[uuid] = o
}

func (dao MapDAO) read(id string) Object {
	return dao[id]
}

func (dao MapDAO) update(o Object) {
	uuid := o.GetID()
	if _, ok := dao[uuid]; !ok {
		return
	}

	dao[uuid] = o
}

func (dao MapDAO) remove(id string) {
	delete(dao, id)
}

func (dao MapDAO) list() []Object {
	list := make([]Object, 0, len(dao))
	for _, o := range dao {
		list = append(list, o)
	}
	return list
}
