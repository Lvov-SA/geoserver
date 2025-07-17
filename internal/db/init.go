package db

import (
	"geoserver/internal/db/models"
	"geoserver/internal/db/seeds"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Init() {
	dbPath := GetDbPath()

	_, err := os.Stat(dbPath)
	dbExists := !os.IsNotExist(err)

	// 2. Открываем/создаем БД
	db := GetConnection()

	// 3. Если БД не существовала - применяем миграции
	if !dbExists {
		runMigrate(db)
		runSeed(db)
	}
}

func GetDbPath() string {
	return "../resource/" + os.Getenv("DB_DATABASE")
}

func GetConnection() *gorm.DB {

	dbName := GetDbPath()
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
}

func runSeed(db *gorm.DB) {
	seeds.User(db)
}
