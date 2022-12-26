package models

import "gorm.io/gorm"

type Family struct {
	// Id        int64          `json:"id" gorm:"primaryKey"`
	gorm.Model
	UserId   int64  `json:"user_id"`
	ParentId uint32 `json:"parent_id"`
	Name     string `json:"name"`
	Gender   string `json:"gender"`
	Order    int32  `json:"order"`
	// Families []Family `json:"families" gorm:"foreignkey:ParentId;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
