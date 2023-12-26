// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"time"
)

const TableNameThemeSetting = "theme_setting"

// ThemeSetting mapped from table <theme_setting>
type ThemeSetting struct {
	ID           int32      `gorm:"column:id;type:int(11);primaryKey;autoIncrement:true" json:"id"`
	CreateTime   time.Time  `gorm:"column:create_time;type:datetime(6);not null" json:"create_time"`
	UpdateTime   *time.Time `gorm:"column:update_time;type:datetime(6)" json:"update_time"`
	SettingKey   string     `gorm:"column:setting_key;type:varchar(255);not null;index:theme_setting_setting_key,priority:1" json:"setting_key"`
	ThemeID      string     `gorm:"column:theme_id;type:varchar(255);not null;index:theme_setting_theme_id,priority:1" json:"theme_id"`
	SettingValue string     `gorm:"column:setting_value;type:longtext;not null" json:"setting_value"`
}

// TableName ThemeSetting's table name
func (*ThemeSetting) TableName() string {
	return TableNameThemeSetting
}
