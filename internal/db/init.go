package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func init() {
	dbPath := getDbPath()

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

func getDbPath() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return "../resource" + os.Getenv("DB_DATABASE")
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

}

func runSeed(db *gorm.DB) {

}
