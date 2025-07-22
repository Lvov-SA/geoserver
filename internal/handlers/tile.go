package handlers

import (
	"geoserver/internal/db"
	"geoserver/internal/db/models"
	"geoserver/internal/loader"
	"geoserver/internal/render"
	"image/png"
	"net/http"
	"strconv"
)

func GetTileHandler(w http.ResponseWriter, r *http.Request) {
	db := db.GetConnection()
	layerName := r.FormValue("layer")
	styleName := r.FormValue("style")
	format := r.FormValue("format")
	tileMatrix, _ := strconv.Atoi(r.FormValue("tilematrix")) // Zoom
	tileCol, _ := strconv.Atoi(r.FormValue("tilecol"))       // X
	tileRow, _ := strconv.Atoi(r.FormValue("tilerow"))       // Y

	// Проверяем слой
	var layer models.Layer
	if err := db.Where("name = ?", layerName).First(&layer).Error; err != nil {
		http.Error(w, "not existt layer", http.StatusBadRequest)
		return
	}

	// Проверяем стиль (если указан)
	if styleName != "" {
		var style models.Style
		if err := db.Where("layer_id = ? AND name = ?", layer.ID, styleName).First(&style).Error; err != nil {
			http.Error(w, "not existt style", http.StatusBadRequest)
			return
		}
	}
	_ = format

	img := render.Tile(loader.Datasets[layerName], TileSize, tileCol, tileRow, tileMatrix, loader.Datasets[layerName].RasterXSize(), loader.Datasets[layerName].RasterYSize())
	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, img); err != nil {
		http.Error(w, "Faild png encode", http.StatusBadRequest)
	}
}
