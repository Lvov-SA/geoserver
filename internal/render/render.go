package render

import (
	"fmt"
	"image"
	"math"
	"os"
	"os/exec"
	"strconv"

	"github.com/lukeroth/gdal"
)

func GdalTile(dataset gdal.Dataset, tileSize, x, y, z, xSize, ySize int) image.Image {

	coef := 1 << z
	maxSize := max(dataset.RasterXSize(), dataset.RasterYSize())
	readSize := int(float64(maxSize) / float64(coef))
	fmt.Printf("Размер тайла %v, Размер чтения %v", tileSize, readSize)
	fmt.Println()
	// Проверка границ
	if x*readSize >= maxSize || y*readSize >= maxSize {
		panic("Проверка границ")
	}
	// 2. Настройки для gdal.Translate
	options := []string{
		"-srcwin",
		fmt.Sprintf("%d", x*readSize),
		fmt.Sprintf("%d", y*readSize),
		fmt.Sprintf("%d", readSize),
		fmt.Sprintf("%d", readSize),
		"-outsize",
		fmt.Sprintf("%d", tileSize),
		fmt.Sprintf("%d", tileSize),
	}
	fmt.Println(options)
	// 3. Создаем временный файл
	zstr := strconv.Itoa(z)
	ystr := strconv.Itoa(y)
	xstr := strconv.Itoa(x)
	outputPath := "../resource/" + zstr + ystr + xstr + ".png"

	// 4. Выполняем преобразование
	outputDS, err := gdal.Translate(outputPath, dataset, options)
	if err != nil {
		return nil
	}

	defer outputDS.Close()
	file, err := os.Open(outputPath)
	if err != nil {
		panic(err.Error())
	}
	imageRGBA, _, _ := image.Decode(file)
	return imageRGBA
}

func CustomTile(dataset gdal.Dataset, tileSize, x, y, z, xSize, ySize int) image.Image {

	coef := 1 << z
	maxSize := min(dataset.RasterXSize(), dataset.RasterYSize())
	readSize := int(float64(maxSize) / float64(coef))
	fmt.Printf("Размер тайла %v, Размер чтения %v", tileSize, readSize)
	fmt.Println()
	// Проверка границ
	if x*readSize >= xSize || y*readSize >= ySize {
		panic("Проверка границ")
	}
	options := []string{
		"-srcwin",
		fmt.Sprintf("%d", x*readSize),
		fmt.Sprintf("%d", y*readSize),
		fmt.Sprintf("%d", readSize),
		fmt.Sprintf("%d", readSize),
		"-outsize",
		fmt.Sprintf("%d", tileSize),
		fmt.Sprintf("%d", tileSize),
	}
	dataset, _ = gdal.Translate("", dataset, options)
	outputPath := "../resource/tilesss.png"
	gdal.Translate(outputPath, dataset, nil)
	file, err := os.Open(outputPath)
	if err != nil {
		panic(err.Error())
	}
	imageRGBA, _, _ := image.Decode(file)
	return imageRGBA
}

func CliRender(dataset gdal.Dataset, tileSize, x, y, z, xSize, ySize int) image.Image {
	coef := math.Pow(2, float64(z))
	maxSize := min(dataset.RasterXSize(), dataset.RasterYSize())
	readSize := int(float64(maxSize) / float64(coef))
	fmt.Printf("Размер тайла %v, Размер чтения %v", tileSize, readSize)
	fmt.Println()
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
	file, _ := os.Open(tmpPath)
	imageRGBA, _, _ := image.Decode(file)
	return imageRGBA
}
