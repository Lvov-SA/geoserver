package db

import (
	"geoserver/internal/db/models"
	"geoserver/internal/db/seeds"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	dbPath := getDbPath()

	_, err := os.Stat(dbPath)
	dbExists := !os.IsNotExist(err)

	// 2. Открываем/создаем БД
	db := GetConnection()

	runMigrate(db)
	// 3. Если БД не существовала - применяем миграции
	if !dbExists {
		runMigrate(db)
		runSeed(db)
	}
}

func getDbPath() string {
	return "../resource/" + os.Getenv("DB_DATABASE")
}

func GetConnection() *gorm.DB {

	dbName := getDbPath()
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func runMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Layer{})
	db.AutoMigrate(&models.Style{})
	db.AutoMigrate(&models.TileMatrix{})
	db.AutoMigrate(&models.TileMatrixSet{})
}

func runSeed(db *gorm.DB) {
	seeds.User(db)
}
