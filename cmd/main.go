package main

import (
	"geoserver/pkg/loader"
	"geoserver/pkg/render"

	"image/png"
	"log"
	"net/http"
	"strconv"
)

const tileSize = 256

func main() {
	dataset, err := loader.GeoTiff()
	if err != nil {
		panic(err.Error())
	}
	xSize := dataset.RasterXSize()
	ySize := dataset.RasterYSize()
	log.Printf("GeoTIFF info: %dx%d pixels", xSize, ySize)

	http.HandleFunc("/tile", func(w http.ResponseWriter, r *http.Request) {
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

		img := render.Tile(&dataset, tileSize, x, y, z, xSize, ySize)

		w.Header().Set("Content-Type", "image/png")
		if err := png.Encode(w, img); err != nil {
			log.Printf("PNG encode error: %v", err)
		}
	})

	log.Println("Server started at :8080")
	log.Println("Access example: http://localhost:8080/tile?z=0&x=0&y=0")
	http.ListenAndServe(":8080", nil)
}
