package main

import "src/github.com/gorilla/mux"
import "log"
import "net/http"
import "encoding/json"

type Person struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"firstname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,omitempty"`
}

var people []Person

//GetPersonEndPoint Get single person endpoint
func GetPersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for _, item := range people {

		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Person{})
}

//GetPeopleEndPoint Get all people
func GetPeopleEndPoint(w http.ResponseWriter, req *http.Request) {
	json.NewEncoder(w).Encode(people)
}

//CreatePersonEndPoint Create a new person endpoint
func CreatePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var person Person
	_ = json.NewDecoder(req.Body).Decode(&person)
	person.ID = params["id"]
	people = append(people, person)
	json.NewEncoder(w).Encode(people)
}

//DeletePersonEndPoint delete a person endpoint
func DeletePersonEndPoint(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	for index, person := range people {
		if person.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)
}

func main() {
	router := mux.NewRouter()
	people = append(people, Person{ID: "1", FirstName: "Sajjan", LastName: "Jyothi", Address: &Address{City: "London", State: "Hounslow"}})
	people = append(people, Person{ID: "2", FirstName: "Sreemithra", LastName: "Rajeevan"})
	router.HandleFunc("/people", GetPeopleEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersonEndPoint).Methods("GET")
	router.HandleFunc("/people/{id}", CreatePersonEndPoint).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersonEndPoint).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":12345", router))

}
