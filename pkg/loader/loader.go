package loader

import (
	"log"

	"github.com/lukeroth/gdal"
)

func GeoTiff() (gdal.Dataset, error) {
	dataset, err := gdal.Open("../resourse/map/NE1_LR_LC.tif", gdal.ReadOnly)
	if err != nil {
		log.Fatal("Failed to open GeoTIFF:", err)
	}
	defer dataset.Close()
	return dataset, err
}
