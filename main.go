package main

import (
	"encoding/json"
	"net/http"
	// "strconv"

	"github.com/gorilla/mux"

)

var profiles []Profile = []Profile{}

type User struct{
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
}

type Profile struct{
	Department  string `json:"department"`
	Designation string `json:"designation"`
	Employee 	User `json:"employee"`
}

func getAllProfiles(q http.ResponseWriter, r *http.Request){
	q.Header().Set("Content-Type", "application/json")
	json.NewEncoder(q).Encode(profiles)
}
func addItem(q http.ResponseWriter, r *http.Request) {
	var newProfile Profile
	json.NewDecoder(r.Body).Decode(&newProfile)

	q.Header().Set("Content-Type", "application/json")

	profiles = append(profiles, newProfile)

	json.NewEncoder(q).Encode(profiles)
}

func main() {
	router:= mux.NewRouter()

	router.HandleFunc("/profiles", addItem).Methods("POST")
	router.HandleFunc("/profiles", getAllProfiles).Methods("GET")
	
	http.ListenAndServe(":5000", router)
}
