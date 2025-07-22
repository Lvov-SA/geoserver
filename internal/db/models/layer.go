package models

import "gorm.io/gorm"

// Layer - картографический слой WMTS
type Layer struct {
	gorm.Model

	// Идентификаторы
	Identifier string `gorm:"size:100;uniqueIndex;not null"` // Уникальный ID (напр. 'osm_roads')
	Title      string `gorm:"size:255;not null"`             // Человекочитаемое название
	Abstract   string `gorm:"type:text"`                     // Описание слоя
	SourcePath string `gorm:"type:string"`                   // Файл слоя : имя + расширение
	// Система координат
	CRS string `gorm:"size:20;default:'EPSG:3857'"` // Система координат (напр. 'EPSG:3857')

	// Границы слоя
	MinX float64 `gorm:"not null"` // Левая граница (в единицах CRS)
	MinY float64 `gorm:"not null"` // Нижняя граница
	MaxX float64 `gorm:"not null"` // Правая граница
	MaxY float64 `gorm:"not null"` // Верхняя граница

	// Уровни масштабирования
	MinZoom int `gorm:"default:0"`  // Минимальный zoom
	MaxZoom int `gorm:"default:22"` // Максимальный zoom

	// Форматы тайлов
	Formats string `gorm:"size:100;default:'image/png'"` // Доступные форматы ('image/png,image/jpeg')

	// Состояние
	IsActive bool `gorm:"default:true"` // Активен ли слой

	// Связи
	Styles          []Style       `gorm:"foreignKey:LayerID"` // Стили слоя
	TileMatrixSetID uint          // Ссылка на матрицу тайлов
	TileMatrixSet   TileMatrixSet `gorm:"foreignKey:TileMatrixSetID;constraint:OnDelete:CASCADE"`
}
