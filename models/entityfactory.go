package models

import (
	"strings"

	"github.com/devgek/webskeleton/types"
)

//EntityFactory create Entities by name
type EntityFactory struct {
}

//Get return entity struct by name
func (ef EntityFactory) Get(entityName string) interface{} {
	entityType := types.ParseEntityType(strings.ToLower(entityName))
	switch entityType {
	case types.EntityTypeUser:
		return &User{}
	case types.EntityTypeContact:
		return &Contact{}
	case types.EntityTypeContactAddress:
		return &ContactAddress{}
	default:
		panic("Undefind entity " + entityName)
	}
}

//GetSlice return slice of entity struct by name
func (ef EntityFactory) GetSlice(entityName string) interface{} {
	switch strings.ToLower(entityName) {
	case "user":
		return &[]User{}
	case "contact":
		return &[]Contact{}
	case "contactaddress":
		return &[]ContactAddress{}
	default:
		panic("Undefind entity " + entityName)
	}
}
