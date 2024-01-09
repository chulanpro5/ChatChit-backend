package entity

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID                 uint           `gorm:"primarykey" json:"id"`
	CreatedAt          time.Time      `json:"-"`
	UpdatedAt          time.Time      `json:"-"`
	DeletedAt          gorm.DeletedAt `gorm:"index" json:"-"`
	Name               string         `json:"name"`
	Email              string         `json:"email" gorm:"unique"`
	Password           []byte         `json:"-"`
	ProfileImageBase64 string         `json:"avatar"`
	LanguageId         uint           `json:"languageId"`
	Language           Language       `gorm:"foreignKey:LanguageId" json:"language"`
}
