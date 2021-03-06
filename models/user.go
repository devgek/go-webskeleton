package models

import (
	"github.com/devgek/webskeleton/dtos"
	"github.com/devgek/webskeleton/types"
)

//User ...
type User struct {
	Entity
	Name  string         `gorm:"type:varchar(50);not null;unique" form:"gkvName"`
	Pass  string         `gorm:"type:text;not null" form:"gkvPass"`
	Email string         `gorm:"type:varchar(100);not null" form:"gkvEmail"`
	Role  types.RoleType `gorm:"type:integer;not null" form:"gkvRole"`
}

//BuildEntityOption ...
func (u *User) BuildEntityOption() dtos.EntityOption {
	o := dtos.EntityOption{}
	o.ID = u.ID
	o.Value = u.Name + ":" + u.Email

	return o
}
