package main

import (
	"log"
	"net/http"
	"starter_kit_rest_api_golang/database"
	"starter_kit_rest_api_golang/model"
	"starter_kit_rest_api_golang/route"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {
	initDB()
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)
	route.Route(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3306",
			User:       "root",
			Password:   "",
			DB:         "go",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&model.Person{})

}
