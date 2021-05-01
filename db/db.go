package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"luxot.tech/stratum/model"
)

var Db *gorm.DB

func InitPostgresDB(dbUser string, dbPwd string, dbName string, dbHost string, dbPort string) {
	var err error
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, dbUser, dbName, dbPwd)
	Db, err = gorm.Open("postgres", dbUrl)
	if err != nil {
		fmt.Printf("Failed to connect to database %v", err)
		log.Panic("Database connection failed")
	} else {
		fmt.Println("Database connection successful")
		runDatabaseMigration()
	}
}

func runDatabaseMigration() {
	Db.AutoMigrate(&model.AuthorizationRequest{}, &model.Subscription{})
}
