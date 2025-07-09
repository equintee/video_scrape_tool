package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Content struct {
	gorm.Model
	Id         uuid.UUID `gorm:"primarykey"`
	Name       string    `gorm:"type:varchar(255)"`
	ContentUrl string    `gorm:"type:varchar(255)"`
	Tags       []string  `gorm:"type:varchar(255)"`
	Song       Song
}

type Song struct {
	Name   string `gorm:"type:varchar(255)"`
	Artist string `gorm:"type:varchar(255)"`
}
