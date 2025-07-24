package loader

import (
	"fmt"
	"log"

	"github.com/Lvov-SA/gdal"
)

var Dataset gdal.Dataset

func GeoTiff() (gdal.Dataset, error) {
	var err error
	Dataset, err = gdal.Open("../resource/map/geo_map.tif", gdal.ReadOnly)
	if err != nil {
		log.Fatal("Failed to open GeoTIFF:", err)
	}

	fmt.Println()
	return Dataset, err
}
