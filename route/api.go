package route

import (
	"starter_kit_rest_api_golang/controllers"
	"starter_kit_rest_api_golang/middleware"

	"github.com/gorilla/mux"
)

func Route(router *mux.Router) {

	userrouter := router.PathPrefix("/user").Subrouter()
	userrouter.Use(middleware.Authorization)
	userrouter.HandleFunc("/", controllers.GetAllPerson).Methods("GET")
	userrouter.HandleFunc("/create", controllers.CreatePerson).Methods("POST")
	userrouter.HandleFunc("/find/{id}", controllers.GetPersonByID).Methods("GET")
	userrouter.HandleFunc("update/{id}", controllers.UpdatePersonByID).Methods("PUT")
	userrouter.HandleFunc("delete/{id}", controllers.DeletPersonByID).Methods("DELETE")
}

func Auth(router *mux.Router) {
	router.HandleFunc("/signin", controllers.SignIn).Methods("POST")
	router.HandleFunc("/signup", controllers.SignUp).Methods("POST")
	router.HandleFunc("/refresh_token", controllers.GetAllPerson).Methods("GET")
	router.HandleFunc("/logout", controllers.GetAllPerson).Methods("GET")
}
