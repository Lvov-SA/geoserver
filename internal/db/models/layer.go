package models

import "gorm.io/gorm"

type Layer struct {
	gorm.Model

	// Идентификаторы
	Name  string `gorm:"size:100;uniqueIndex;not null"` // Уникальный ключ (например: "osm_roads")
	Title string `gorm:"size:255;not null"`             // Отображаемое название ("Дороги OSM")

	// Источник данных
	SourceType string `gorm:"size:50;not null"` // "geotiff", "mbtiles", "postgis"
	SourcePath string `gorm:"not null"`         // Путь к файлу или DSN

	// Настройки видимости
	MinZoom  int  `gorm:"default:0"`    // Минимальный zoom
	MaxZoom  int  `gorm:"default:22"`   // Максимальный zoom
	IsActive bool `gorm:"default:true"` // Доступен для запросов

	// Форматы
	DefaultFormat    string `gorm:"size:20;default:'image/png'"` // Основной формат тайлов
	AvailableFormats string `gorm:"type:json"`                   // ["image/png", "image/webp"]
}
