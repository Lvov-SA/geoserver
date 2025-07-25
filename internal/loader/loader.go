package loader

import (
	"geoserver/internal/db"
	"geoserver/internal/db/models"

	"github.com/Lvov-SA/gdal"
)

var Layers map[string]models.Layer

func GeoTiff() error {
	var layers []models.Layer
	Layers = make(map[string]models.Layer)

	db, err := db.GetConnection()
	if err != nil {
		return err
	}

	db.Find(&layers)
	for _, layer := range layers {

		src := "../resource/map/" + layer.SourcePath
		dataset, err := gdal.Open(src, gdal.ReadOnly)
		if err != nil {
			return err
		}
		dataset.Close()
		Layers[layer.Name] = layer
	}
	return nil
}
