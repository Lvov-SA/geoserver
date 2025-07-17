package main

import (
	"fmt"
	"geoserver/internal/admin"
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
	mux := admin.Init(db.GetConnection())
	_, err = loader.GeoTiff()
	if err != nil {
		panic(err.Error())
	}
	//defer loader.Dataset.Close()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("/image-info", handlers.ImageInfoHandler)
	mux.HandleFunc("/tile", handlers.TileHandler)

	log.Println("Server started at :8080")
	log.Println("Access example: http://localhost:8080/tile?z=0&x=0&y=0")
	err = http.ListenAndServe(":8080", mux)
	fmt.Printf("%v", err)
}
