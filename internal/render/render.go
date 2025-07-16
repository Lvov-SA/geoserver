package render

import (
	"fmt"
	"image"

	"github.com/lukeroth/gdal"
)

func Tile(dataset gdal.Dataset, tileSize, x, y, z, xSize, ySize int) *image.RGBA {

	var img *image.RGBA
	coef := 1 << z
	maxSize := min(dataset.RasterXSize(), dataset.RasterYSize())
	readSize := int(float64(maxSize) / float64(coef))
	fmt.Printf("Размер тайла %v, Размер чтения %v", tileSize, readSize)
	fmt.Println()
	// Проверка границ
	if x*readSize >= xSize || y*readSize >= ySize {
		panic("Проверка границ")
	}

	img = image.NewRGBA(image.Rect(0, 0, tileSize, tileSize))
	for b := 0; b < 3 && b < dataset.RasterCount(); b++ { // Первые 3 канала (R,G,B)
		data := make([]uint8, tileSize*tileSize) // Всегда 256x256
		err := dataset.RasterBand(b+1).IO(
			gdal.Read,
			x*readSize, y*readSize, // Смещение в исходном растре
			readSize, readSize, // Размер считываемой области
			data,               // Буфер для данных
			tileSize, tileSize, // Размер выходного тайла
			0, 0,
		)
		if err != nil {
			panic(err.Error())
		}

		for i := 0; i < tileSize*tileSize; i++ {
			img.Pix[i*4+b] = data[i]
		}
		fmt.Println("Размер чтения:", readSize)
		fmt.Println("Первые 10 значений data:", data[:10])
		fmt.Println("Первые 10 значений img.Pix:", img.Pix[:40])
	}
	//Устанавливаем полную непрозрачность в 4 канале
	for i := 3; i < len(img.Pix); i += 4 {
		img.Pix[i] = 255
	}
	return img
}
