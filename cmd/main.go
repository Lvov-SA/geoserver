package main

import (
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
	//defer loader.Dataset.Close()
	http.HandleFunc("/", handlers.IndexHandler)
	http.HandleFunc("/image-info", handlers.ImageInfoHandler)
	http.HandleFunc("/tile", handlers.TileHandler)

	http.HandleFunc("/wmts/1.0.0/WMTSCapabilities.xml", handlers.GetCapabilitiesHandler)
	http.HandleFunc("/wmts/1.0.0/tile/", handlers.GetTileHandler)
	//http.HandleFunc("/wmts/1.0.0/featureinfo/", GetFeatureInfoHandler)

	log.Println("Server started at :8080")
	log.Println("Access example: http://localhost:8080/tile?z=0&x=0&y=0")
	http.ListenAndServe(":8080", nil)
}
