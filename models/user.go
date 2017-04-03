package models

import (
	"github.com/jinzhu/gorm"
	//"managesys/db"
	//"strconv"
)

type User struct {
	gorm.Model
	Name     string
	Email    string
	Password string
	Role     string
}

func (u *User) DisplayName() string {
	return u.Name
}
