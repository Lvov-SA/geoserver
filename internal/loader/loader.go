package loader

import (
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

	return Dataset, err
}
