package loader

import (
	"fmt"
	"log"

	"github.com/lukeroth/gdal"
)

var Dataset gdal.Dataset

func GeoTiff() (gdal.Dataset, error) {
	var err error
	Dataset, err = gdal.Open("../resource/map/geo_map.tif", gdal.ReadOnly)
	if err != nil {
		log.Fatal("Failed to open GeoTIFF:", err)
	}

	fmt.Printf(
		"GeoTIFF info: %dx%d pixels, bands count: %v",
		Dataset.RasterXSize(),
		Dataset.RasterYSize(),
		Dataset.RasterCount(),
	)
	fmt.Println()
	return Dataset, err
}
