package models

import "gorm.io/gorm"

type Style struct {
	gorm.Model

	// Идентификаторы
	Identifier string `gorm:"size:50;not null"` // Идентификатор стиля (напр. 'dark')
	Title      string `gorm:"size:255"`         // Название стиля

	// Настройки
	IsDefault bool   `gorm:"default:false"` // Стиль по умолчанию?
	LegendURL string `gorm:"type:text"`     // URL легенды (опционально)

	// Связи
	LayerID uint  `gorm:"index;not null"` // Ссылка на родительский слой
	Layer   Layer `gorm:"foreignKey:LayerID;constraint:OnDelete:CASCADE"`
}
