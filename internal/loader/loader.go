package loader

import (
	"fmt"
	"geoserver/internal/db"
	"geoserver/internal/db/models"
	"log"

	"github.com/lukeroth/gdal"
)

// Ключ - Name из базы
var Datasets map[string]gdal.Dataset

func GeoTiff() (gdal.Dataset, error) {
	Datasets = make(map[string]gdal.Dataset)
	db := db.GetConnection()
	var layers []models.Layer
	db.Find(&layers)
	for _, layer := range layers {
		src := "../resource/map/" + layer.SourcePath
		Dataset, err := gdal.Open(src, gdal.ReadOnly)
		if err != nil {
			log.Fatal("Failed to open GeoTIFF:", err)
			fmt.Printf(
				"GeoTIFF info: %dx%d pixels, bands count: %v",
				Dataset.RasterXSize(),
				Dataset.RasterXSize(),
				Dataset.RasterCount(),
			)
		}
		Datasets[layer.Identifier] = Dataset
	}

	fmt.Println()
	return Datasets["sda"], nil
}
