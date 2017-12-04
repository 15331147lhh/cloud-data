package entities

import (
	"time"
)


type UserInfo struct {
	UID        int `xorm:"pk" xorm:"autoincr"` 
	UserName   string
	DepartName string
	CreatedAt  time.Time `xorm:"created"` 
}


func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.UserName) == 0 {
		panic("UserName shold not be null!")
	}
	if u.CreateAt == nil {
		t := time.Now()
		u.CreateAt = &t
	}
	return &u
}
