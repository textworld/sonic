// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"

	"github.com/go-sonic/sonic/consts"
)

const TableNameUser = "user"

// User mapped from table <user>
type User struct {
	ID          int32          `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	CreateTime  time.Time      `gorm:"column:create_time;type:datetime(6);not null" json:"create_time"`
	UpdateTime  *time.Time     `gorm:"column:update_time;type:datetime(6)" json:"update_time"`
	Avatar      string         `gorm:"column:avatar;type:varchar(1023);not null" json:"avatar"`
	Description string         `gorm:"column:description;type:varchar(1023);not null" json:"description"`
	Email       string         `gorm:"column:email;type:varchar(127);not null" json:"email"`
	ExpireTime  *time.Time     `gorm:"column:expire_time;type:datetime(6)" json:"expire_time"`
	MfaKey      string         `gorm:"column:mfa_key;type:varchar(64);not null" json:"mfa_key"`
	MfaType     consts.MFAType `gorm:"column:mfa_type;type:bigint;not null" json:"mfa_type"`
	Nickname    string         `gorm:"column:nickname;type:varchar(255);not null" json:"nickname"`
	Password    string         `gorm:"column:password;type:varchar(255);not null" json:"password"`
	Username    string         `gorm:"column:username;type:varchar(50);not null" json:"username"`
}

// TableName User's table name
func (*User) TableName() string {
	return TableNameUser
}
