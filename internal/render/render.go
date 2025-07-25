package render

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
	"os/exec"

	"github.com/Lvov-SA/gdal"
)

func CliRender(dataset gdal.Dataset, tileSize, x, y, z, xSize, ySize int) image.Image {
	coef := math.Pow(2, float64(z))
	maxSize := min(dataset.RasterXSize(), dataset.RasterYSize())
	readSize := int(float64(maxSize) / float64(coef))
	if x*readSize >= xSize || y*readSize >= ySize {
		panic("Проверка границ")
	}
	tmpFile, err := os.CreateTemp("../resource", "tile_*.png")
	if err != nil {
		panic(err.Error())
	}
	tmpPath := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpPath)              // Удаляем в конце
	defer os.Remove(tmpPath + ".aux.xml") // Удаляем в конце
	cmd := exec.Command("gdal_translate", "-srcwin",
		fmt.Sprintf("%d", x*readSize),
		fmt.Sprintf("%d", y*readSize),
		fmt.Sprintf("%d", readSize),
		fmt.Sprintf("%d", readSize),
		"-outsize",
		fmt.Sprintf("%d", tileSize),
		fmt.Sprintf("%d", tileSize),
		"../resource/map/geo_map.tif",
		tmpPath)
	cmd.Run()

	file, err := os.Open(tmpPath)
	if err != nil {
		fmt.Printf("%v", err)
	}
	encoder := png.Encoder{
		CompressionLevel: png.BestCompression, // Максимальное сжатие
	}
	imageRGBA, _, err := image.Decode(file)
	encoder.Encode(file, imageRGBA)
	imageRGBA, _, err = image.Decode(file)
	if err != nil {
		fmt.Printf("%v", err)
	}
	return imageRGBA
}
