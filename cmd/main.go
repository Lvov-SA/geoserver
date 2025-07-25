package main

import (
	"fmt"
	"geoserver/internal/config"
	"geoserver/internal/db"
	"geoserver/internal/handlers"
	"geoserver/internal/loader"

	"log"
	"net/http"
)

const TileSize = 256

func main() {

	err := config.Init()
	if err != nil {
		fmt.Printf("Ошибка инициализации конфигурации приложения: %v", err)
		return
	}
	err = db.Init()
	if err != nil {
		fmt.Printf("Ошибка инициализации базы данных: %v", err)
		return
	}
	err = loader.GeoTiff()
	if err != nil {
		fmt.Printf("Ошибка загрузки данных слоев: %v", err)
		return
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.IndexHandler)
	mux.HandleFunc("GET /{tile}/{z}/{x}/{y}", handlers.TileHandler)

	log.Println("Server started at :8080")
	log.Println("Access example: http://localhost:8080/tile/0/0/0.png")
	log.Println("Look at map: http://localhost:8080")

	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		fmt.Printf("Ошибка запуска сервера: %v", err)
		return
	}
}
