package handlers

import (
	"encoding/json"
	"geoserver/internal/loader"
	"geoserver/internal/render"
	"image/png"
	"log"
	"net/http"
	"strconv"
)

const TileSize = 256

func ImageInfoHandler(w http.ResponseWriter, r *http.Request) {
	info := struct {
		Width  int `json:"width"`
		Height int `json:"height"`
	}{
		Width:  loader.Datasets["geo_map"].RasterXSize(),
		Height: loader.Datasets["geo_map"].RasterYSize(),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(info)
}

func TileHandler(w http.ResponseWriter, r *http.Request) {
	z, err := strconv.Atoi(r.URL.Query().Get("z"))
	if err != nil || z < 0 {
		http.Error(w, "Invalid z parameter", http.StatusBadRequest)
		return
	}

	x, err := strconv.Atoi(r.URL.Query().Get("x"))
	if err != nil || x < 0 {
		http.Error(w, "Invalid x parameter", http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil || y < 0 {
		http.Error(w, "Invalid y parameter", http.StatusBadRequest)
		return
	}
	img := render.Tile(loader.Datasets["geo_map"], TileSize, x, y, z, loader.Datasets["geo_map"].RasterXSize(), loader.Datasets["geo_map"].RasterYSize())
	w.Header().Set("Content-Type", "image/png")
	if err := png.Encode(w, img); err != nil {
		log.Printf("PNG encode error: %v", err)
	}
}
