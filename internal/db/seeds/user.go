package seeds

import (
	"geoserver/internal/db/models"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func User(db *gorm.DB) {
	var user models.User
	user.Login = os.Getenv("USER_LOGIN")
	hash, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("USER_PASSWORD")), 0)
	if err != nil {
		panic(err.Error())
	}
	user.Passowrd = string(hash)
	db.Create(&user)
}
