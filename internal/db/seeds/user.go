package seeds

import (
	"fmt"
	"geoserver/internal/config"
	"geoserver/internal/db/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func User(db *gorm.DB) error {
	var user models.User
	user.Login = config.Configs.USER_LOGIN
	hash, err := bcrypt.GenerateFromPassword([]byte(config.Configs.USER_PASSWORD), 0)
	if err != nil {
		return fmt.Errorf("Ошибка генерации пароля пользователю %w", err)
	}
	user.Passowrd = string(hash)
	db.Create(&user)
	return nil
}
