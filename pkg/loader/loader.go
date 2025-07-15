package loader

import (
	"log"

	"github.com/lukeroth/gdal"
)

func GeoTiff() (gdal.Dataset, error) {
	dataset, err := gdal.Open("../resource/map/geo_map.tif", gdal.ReadOnly)
	if err != nil {
		log.Fatal("Failed to open GeoTIFF:", err)
	}
	return dataset, err
}
