package seeders

import (
	"user-service/domain/models"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func RunRoleSeeder(db *gorm.DB) {
	roles := []models.Role{
		{Code: "ADMIN", Name: "Administrator"},
		{Code: "CUSTOMER", Name: "Customer"},
	}

	for _, role := range roles {
		err := db.Clauses(clause.OnConflict{DoNothing: true}).Create(&role).Error

		if err != nil {
			logrus.Errorf("Failed to seed role %s: %v", role.Code, err)
		} else {
			logrus.Infof("Role %s successfully seeded or already exists", role.Code)
		}
	}
}
