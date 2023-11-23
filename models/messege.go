package models

import "time"

type Messege struct {
	Id      int64     `gorm:"primaryKey" json:"id"`
	From    string    `gorm:"type:varchar(300)" json:"from"`
	Content string    `gorm:"type:varchar(300)" json:"content"`
	Time    time.Time `gorm:"default:CURRENT_TIMESTAMP()" json:"time"`
}
