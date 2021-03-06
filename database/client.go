package database

import (
	"log"
	"starter_kit_rest_api_golang/model"

	"github.com/jinzhu/gorm"
)

//Connector variable used for CRUD operation's
var Connector *gorm.DB

//Connect creates MySQL connection
func Connect(connectionString string) error {
	var err error
	Connector, err = gorm.Open("mysql", connectionString)
	if err != nil {
		return err
	}
	log.Println("Connection was successful!!")
	return nil
}

func CloseDatabase(connection *gorm.DB) {
	sqldb := connection.DB()
	sqldb.Close()
}

//Migrate create/updates database table
func Migrate(table *model.User) {
	Connector.AutoMigrate(&table)
	log.Println("Table migrated")
}
