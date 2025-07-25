package render

import (
	"errors"
	"fmt"
	"geoserver/internal/db/models"
	"image"
	"math"
	"os"
	"os/exec"
)

func CliRender(layer models.Layer, z, x, y int) (image.Image, error) {
	coef := math.Pow(2, float64(z))
	maxSize := min(layer.Width, layer.Height)
	readSize := int(float64(maxSize) / float64(coef))
	if x*readSize >= layer.Width || y*readSize >= layer.Height {
		return nil, errors.New("Выход за границы")
	}
	tmpFile, err := os.CreateTemp("../resource", "tile_*.png")
	if err != nil {
		return nil, err
	}
	tmpPath := tmpFile.Name()
	tmpFile.Close()
	defer os.Remove(tmpPath)
	defer os.Remove(tmpPath + ".aux.xml")
	cmd := exec.Command("gdal_translate", "-srcwin",
		fmt.Sprintf("%d", x*readSize),
		fmt.Sprintf("%d", y*readSize),
		fmt.Sprintf("%d", readSize),
		fmt.Sprintf("%d", readSize),
		"-outsize",
		fmt.Sprintf("%d", layer.TileSize),
		fmt.Sprintf("%d", layer.TileSize),
		"../resource/map/"+layer.SourcePath,
		tmpPath)
	cmd.Run()

	file, err := os.Open(tmpPath)
	if err != nil {
		return nil, err
	}
	imageRGBA, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return imageRGBA, nil
}
