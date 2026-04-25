package entity

import (
	entitybase "github.com/ArthurTirta/monogo/internal/entity/base"
)

type Admin struct {
	entitybase.Base
	Name     string `gorm:"column:name;type:varchar(255)"`
	Email    string `gorm:"column:email;type:varchar(255);not null;unique"`
	Password string `gorm:"column:password;type:varchar(255);not null"`
}

func (a *Admin) TableName() string {
	return "monogo.admins"
}
