// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameTag = "tag"

// Tag mapped from table <tag>
type Tag struct {
	ID         int32      `gorm:"column:id;type:int;primaryKey;autoIncrement:true" json:"id"`
	CreateTime time.Time  `gorm:"column:create_time;type:datetime(6);not null" json:"create_time"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime(6)" json:"update_time"`
	Name       string     `gorm:"column:name;type:varchar(255);not null;index:tag_name,priority:1" json:"name"`
	Slug       string     `gorm:"column:slug;type:varchar(50);not null;uniqueIndex:uniq_tag_slug,priority:1" json:"slug"`
	Thumbnail  string     `gorm:"column:thumbnail;type:varchar(1023);not null" json:"thumbnail"`
	Color      string     `gorm:"column:color;type:varchar(25);not null" json:"color"`
}

// TableName Tag's table name
func (*Tag) TableName() string {
	return TableNameTag
}
