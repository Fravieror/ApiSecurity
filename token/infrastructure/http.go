package infrastructure

import (
	"apiSecurity/token/application"
	userDomain "apiSecurity/user/domain"
	"encoding/json"
	"net/http"
)

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	var user userDomain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	bussines := application.NewBussines(&MySQLRepository{})
	response := bussines.GenerateToken(&user)
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		w.Write([]byte("Error cast struct to json"))
	}
	w.Write(jsonResponse)
	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.

}
