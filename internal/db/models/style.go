package models

import "gorm.io/gorm"

type Style struct {
	gorm.Model

	// Связь с Layer (многие-к-одному)
	LayerID uint  `gorm:"index;not null"`
	Layer   Layer `gorm:"foreignKey:LayerID"`

	// Идентификаторы
	Name  string `gorm:"size:50;not null"` // "default", "dark", "night"
	Title string `gorm:"size:255"`         // "Темная тема"

	// Правила стилизации
	Rules     string `gorm:"type:json;not null"` // JSON с параметрами
	LegendURL string // URL иконки легенды

	// Метка для WMTS
	IsDefault bool `gorm:"default:false"` // Стиль по умолчанию
}
