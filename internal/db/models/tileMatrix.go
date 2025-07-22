package models

import "gorm.io/gorm"

// TileMatrix - матрица тайлов для конкретного уровня масштабирования
type TileMatrix struct {
	gorm.Model

	// Уровень масштабирования
	ZoomLevel int `gorm:"not null"` // Уровень масштаба (0-22)

	// Параметры тайлов
	ScaleDenominator float64 `gorm:"not null"`    // Масштабный коэффициент
	TopLeftX         float64 `gorm:"not null"`    // Координата X верхнего левого угла
	TopLeftY         float64 `gorm:"not null"`    // Координата Y верхнего левого угла
	TileWidth        int     `gorm:"default:256"` // Ширина тайла (обычно 256)
	TileHeight       int     `gorm:"default:256"` // Высота тайла (обычно 256)
	MatrixWidth      int     `gorm:"not null"`    // Количество тайлов по X
	MatrixHeight     int     `gorm:"not null"`    // Количество тайлов по Y

	// Связи
	TileMatrixSetID uint          `gorm:"index;not null"` // Ссылка на родительский набор
	TileMatrixSet   TileMatrixSet `gorm:"foreignKey:TileMatrixSetID;constraint:OnDelete:CASCADE"`
}
