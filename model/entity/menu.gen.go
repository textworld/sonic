// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameMenu = "menu"

// Menu mapped from table <menu>
type Menu struct {
	ID         int32      `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	CreateTime time.Time  `gorm:"column:create_time;type:datetime(6);not null" json:"create_time"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime(6)" json:"update_time"`
	Icon       string     `gorm:"column:icon;type:varchar(50);not null" json:"icon"`
	Name       string     `gorm:"column:name;type:varchar(50);not null;index:menu_name,priority:1" json:"name"`
	ParentID   int32      `gorm:"column:parent_id;type:int;not null;index:menu_parent_id,priority:1" json:"parent_id"`
	Priority   int32      `gorm:"column:priority;type:int;not null" json:"priority"`
	Target     string     `gorm:"column:target;type:varchar(20);not null;default:_self" json:"target"`
	Team       string     `gorm:"column:team;type:varchar(255);not null" json:"team"`
	URL        string     `gorm:"column:url;type:varchar(1023);not null" json:"url"`
}

// TableName Menu's table name
func (*Menu) TableName() string {
	return TableNameMenu
}
