package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"starter_kit_rest_api_golang/database"
	"starter_kit_rest_api_golang/helpers"
	"starter_kit_rest_api_golang/model"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	// connection := GetDatabase()
	// defer CloseDatabase(connection)

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		var err model.Error
		err = helpers.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var dbuser model.User
	database.Connector.Where("email = ?", user.Email).First(&dbuser)

	//check email is alredy registered or not
	if dbuser.Email != "" {
		var err model.Error
		err = helpers.SetError(err, "Email already in use")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	user.Password, err = helpers.GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("Error in password hashing.")
	}

	//insert user details in database
	log.Println(&user)
	database.Connector.Create(&user)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func SignIn(w http.ResponseWriter, r *http.Request) {

	// defer database.CloseDatabase

	var authDetails model.Authentication

	err := json.NewDecoder(r.Body).Decode(&authDetails)
	if err != nil {
		var err model.Error
		err = helpers.SetError(err, "Error in reading payload.")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var authUser model.User
	database.Connector.Where("email = 	?", authDetails.Email).First(&authUser)

	if authUser.Email == "" {
		var err model.Error
		err = helpers.SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	check := helpers.CheckPasswordHash(authDetails.Password, authUser.Password)

	if !check {
		var err model.Error
		err = helpers.SetError(err, "Username or Password is incorrect")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	validToken, err := helpers.GenerateJWT(authUser.Email, authUser.Role)
	if err != nil {
		var err model.Error
		err = helpers.SetError(err, "Failed to generate token")
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(err)
		return
	}

	var token model.Token
	token.Email = authUser.Email
	token.Role = authUser.Role
	token.TokenString = validToken
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
