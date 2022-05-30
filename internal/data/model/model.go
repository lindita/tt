package model

import (
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

// BaseModel
// Created https://gorm.io/docs/models.html#Creating-Updating-Time-Unix-Milli-Nano-Seconds-Tracking
// Updated https://gorm.io/docs/models.html#Creating-Updating-Time-Unix-Milli-Nano-Seconds-Tracking
type BaseModel struct {
	Id        uint64                `gorm:"primaryKey;column:id;type:bigint(20) unsigned;not null"`
	IsDeleted soft_delete.DeletedAt `gorm:"column:isDeleted;type:tinyint(1) unsigned;not null;default:0;softDelete:flag"`
	Created   int64                 `gorm:"autoCreateTime"`
	Updated   int64                 `gorm:"autoUpdateTime"`
}

// BeforeSave Hooks Before Save
func (m *BaseModel) BeforeSave(tx *gorm.DB) error {
	// 暂时先去掉
	// m.Updated = time.Now().Unix()
	return nil
}

// BeforeCreate Hooks Before Create
func (m *BaseModel) BeforeCreate(tx *gorm.DB) error {
	// 暂时先去掉
	// m.Created = time.Now().Unix()
	return nil
}
