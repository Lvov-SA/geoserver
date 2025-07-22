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
	defer loader.Dataset.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/image-info", handlers.ImageInfoHandler)
	mux.HandleFunc("GET /tile/{z}/{x}/{y}", handlers.TileHandler)

	log.Println("Server started at :8080")
	log.Println("Access example: http://localhost:8080/tile/0/0/0.png")
	http.ListenAndServe(":8080", mux)
}
