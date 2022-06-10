//check whether user is authorized or not
package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"starter_kit_rest_api_golang/helpers"
	"starter_kit_rest_api_golang/model"

	"github.com/casbin/casbin"
	"github.com/dgrijalva/jwt-go"
)

var (
	secretkey string = "secretkeyjwt"
)

func Authorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Token"] == nil {
			var err model.Error
			err = helpers.SetError(err, "No Token Found")
			json.NewEncoder(w).Encode(err)
			return
		}

		var mySigningKey = []byte(secretkey)

		token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("There was an error in parsing token.")
			}
			return mySigningKey, nil
		})

		if err != nil {
			var err model.Error
			err = helpers.SetError(err, "Your Token has been expired.")
			json.NewEncoder(w).Encode(err)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

			// if claims["role"] == "admin" {
			// 	next.ServeHTTP(w, r)
			// }

			var role = claims["role"]
			log.Println(role)

			authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
			if err != nil {
				log.Fatal(err)
			}

			res, err := authEnforcer.EnforceSafe(role, r.URL.Path, r.Method)
			if err != nil {
				var reserr model.Error
				reserr = helpers.SetError(reserr, "Error")
				json.NewEncoder(w).Encode(reserr)
				return
			}
			if res {

				next.ServeHTTP(w, r)
			} else {
				var reserr model.Error
				reserr = helpers.SetError(reserr, "UnAuthorized ")
				json.NewEncoder(w).Encode(reserr)
				return
			}
		}

		// var reserr model.Error
		// reserr = helpers.SetError(reserr, "Not Authorized.")
		// json.NewEncoder(w).Encode(reserr)

	})
}
