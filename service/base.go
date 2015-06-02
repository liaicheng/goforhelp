package service

import (
	"github.com/go-martini/martini"
	"gopkg.in/mgo.v2"
)

const (
	C_PROPERTY = iota
	C_USER
	C_TAGS
)


func ServiceHandler() martini.Handler {
	return func(c martini.Context, db *mgo.Database) {
		myDber := dber{db}
		
		adminUser := adminUserService{}
		c.MapTo(&adminUser, (*IAdminUserService)(nil))
		
		property := propertyService{&myDber} 
		c.MapTo(&property,(*IPropertyService)(nil))
	}
}

type IDber interface {
	C(collection int) *mgo.Collection	
}

type dber struct {
	db *mgo.Database
}

func (d *dber) C(collection int) *mgo.Collection {
	switch collection {
	case C_PROPERTY:
		return d.db.C("PropertyDetail")
	case C_USER:
		return d.db.C("user")
	case C_TAGS:
		return d.db.C("Tags")
	default:
		panic("The collection not found!")
	}
}