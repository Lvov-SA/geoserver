package loader

import (
	"log"

	"github.com/lukeroth/gdal"
)

var (
	geoDataset *gdal.Dataset // GeoTIFF-файл
)

func LoadByGdall() {
	dataset, err := gdal.Open("../resourse/map/NE1_LR_LC.tif", gdal.ReadOnly)
	if err != nil {
		log.Fatal("Ошибка загрузки GeoTIFF:", err)
	}
	geoDataset = &dataset
	_ = geoDataset.GeoTransform()
}
