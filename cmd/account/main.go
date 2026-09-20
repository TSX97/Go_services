package main

import (
	"fmt"
	"strconv"
	"net/http"
	"encoding/json"
	//"github.com/jackc/pgx/v5"	
)

type User struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

var users []User

func newUser(id int, name string) *User {
	return &User{id, name}
}

// === === === ===REST===API=== === === === \\


//	GET /users
func getUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(users)
}

//	GET /users{id}
func getUser(w http.ResponseWriter, r *http.Request){

	w.Header().Set("Content-Type", "application/json")
	id, err := strconv.Atoi(r.PathValue("id"))
	
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(users); i++{
		if users[i].Id == id {
			json.NewEncoder(w).Encode(users[i])
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
}

//	POST /users
func addUser(w http.ResponseWriter, r *http.Request){	
	w.Header().Set("Content-Type", "application/json")
	
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(users); i++ {
		if users[i].Id == user.Id {
			http.Error(w, "User already exists", http.StatusConflict)
			return
		}
	}
	users = append(users, user)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)

}

//	PATCH /users/{id}
func patchUserName(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}

	
	var patch struct {
		Name *string `json:"name"`
	}
	err = json.NewDecoder(r.Body).Decode(&patch)
	if err != nil {
		http.Error(w, "Invalid json", http.StatusBadRequest)
		return
	}

	if patch.Name == nil {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			users[i].Name = *patch.Name
			json.NewEncoder(w).Encode(&users[i])
			return
		}
	}

	http.Error(w, "User not found", http.StatusNotFound)
}

func deleteUser(w http.ResponseWriter, r *http.Request){
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return
	}
	for i := 0; i < len(users); i++ {
		if users[i].Id == id {
			users = append(users[:i], users[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "User not found", http.StatusNotFound)
}

func main(){
	fmt.Println("start serve on localhost:8080")
	http.HandleFunc("GET /users", getUsers)
	http.HandleFunc("GET /users/{id}", getUser)
	http.HandleFunc("POST /users", addUser)
	http.HandleFunc("PATCH /users/{id}", patchUserName)
	http.HandleFunc("DELETE /users/{id}", deleteUser)

	http.ListenAndServe(":8080", nil)
}
