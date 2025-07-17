package main

import (
	"fmt"
	"geoserver/internal/db"
	"geoserver/internal/handlers"
	"geoserver/internal/loader"

	"log"
	"net/http"

	"github.com/joho/godotenv"
)

const TileSize = 256

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.Init()
	_, err = loader.GeoTiff()
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf(
		"GeoTIFF info: %dx%d pixels, bands count: %v",
		loader.Dataset.RasterXSize(),
		loader.Dataset.RasterXSize(),
		loader.Dataset.RasterCount(),
	)
	fmt.Println()
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/image-info", handlers.ImageInfoHandler)
	http.HandleFunc("/tile", handlers.TileHandler)

	log.Println("Server started at :8080")
	log.Println("Access example: http://localhost:8080/tile?z=0&x=0&y=0")
	http.ListenAndServe(":8080", nil)
}
