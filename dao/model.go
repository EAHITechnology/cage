package dao

import (
	"time"
)

type DemoModel struct {
	Id         int64     `gorm:"primaryKey"`
	Info       string    `gorm:"column:info"`
	CreateTime time.Time `gorm:"column:create_time"`
	UpdateTime time.Time `gorm:"column:update_time"`
}
