package seeders

import (
	"user-service/constants"
	"user-service/domain/models"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func RunUserSeeder(db *gorm.DB) {
	// Generate hashed password securely
	password, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		logrus.Fatalf("Failed to hash password: %v", err)
		return
	}

	user := models.User{
		UUID:        uuid.New(),
		Name:        "Administrator",
		Username:    "admin",
		Password:    string(password),
		PhoneNumber: "08893789676",
		Email:       "admin@mail.com",
		RoleID:      constants.Admin,
	}

	// Use OnConflict to prevent duplicate entries efficiently
	err = db.Clauses(clause.OnConflict{DoNothing: true}).Create(&user).Error
	if err != nil {
		logrus.Errorf("Failed to seed user %s: %v", user.Username, err)
	} else {
		logrus.Infof("User %s successfully seeded or already exists", user.Username)
	}
}
