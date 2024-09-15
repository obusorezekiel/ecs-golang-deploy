package model

// Student represents a student with personal and academic details.
type Student struct {
 ID     int    `gorm:"type:int;primary_key"`
 Name   string `gorm:"type:varchar(255)"`
 Description string `gorm:"type:varchar(255)"`
}