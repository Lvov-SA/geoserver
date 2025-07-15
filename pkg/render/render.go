package render

import (
	"image"

	"github.com/lukeroth/gdal"
)

func Tile(dataset *gdal.Dataset, tileSize, x, y, z, xSize, ySize int) *image.RGBA {

	var img *image.RGBA
	scale := 1 << z
	readSize := tileSize * scale

	// Проверка границ
	if x*readSize >= xSize || y*readSize >= ySize {
		panic("Проверка границ")
	}
	img = image.NewRGBA(image.Rect(0, 0, readSize, readSize))
	for b := 0; b < 3; b++ { // Первые 3 канала (R,G,B)
		data := make([]uint8, readSize*readSize)
		err := dataset.RasterBand(b+1).IO(gdal.Read, x*readSize, y*readSize,
			readSize, readSize, data, readSize, readSize, 0, 0)
		if err != nil {
			panic(err.Error())
		}

		// Копируем в соответствующий канал
		for i := 0; i < len(data); i++ {
			img.Pix[i*4+b] = data[i]
		}
	}
	// Устанавливаем полную непрозрачность
	for i := 3; i < len(img.Pix); i += 4 {
		img.Pix[i] = 255
	}
	return img
}
