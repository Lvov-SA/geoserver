package models

import "gorm.io/gorm"

// TileMatrixSet - набор матриц тайлов (например WebMercatorQuad)
type TileMatrixSet struct {
	gorm.Model

	// Идентификаторы
	Identifier  string `gorm:"size:50;uniqueIndex;not null"` // Идентификатор (напр. 'WebMercatorQuad')
	Description string `gorm:"type:text"`                    // Описание набора

	// Система координат
	CRS string `gorm:"size:20;not null"` // Система координат (напр. 'EPSG:3857')

	// Связи
	TileMatrices []TileMatrix `gorm:"foreignKey:TileMatrixSetID;constraint:OnDelete:CASCADE"`
}
