package route

import (
	"rest-go-demo/controllers"

	"github.com/gorilla/mux"
)

func Route(router *mux.Router) {
	router.HandleFunc("/", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	router.HandleFunc("/get/{id}", controllers.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}
