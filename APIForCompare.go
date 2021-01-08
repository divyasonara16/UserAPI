package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "math/rand"
  "strconv"
  "github.com/gorilla/mux"
)
type User struct{
      ID string `json:"Id"`
      Name string `json:"name"`
      Addsress string `json:Addsress`
      MobileNumber int `json:MobileNumber`
      City string `json:city`
      Email string `json:email`
}

// init user var
var users []User


//get All Users
func getAllUser(w http.ResponseWriter, r *http.Request){

  w.Header().Set("Content-Type","application/json")
  json.NewEncoder(w).Encode(users)
}

//get single User
func GetUser(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  params := mux.Vars(r) //to get paramters

  //loop through user and find by id

  for _,item := range users {
    if item.ID == params["Id"]{
      json.NewEncoder(w).Encode(item)
      return
    }
  }
  json.NewEncoder(w).Encode(&User{})
}

//create User
func InsertUser(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  var new User
  _ = json.NewDecoder(r.Body).Decode(&new)
  new.ID = strconv.Itoa(rand.Intn(1000000))
  users = append(users,new)
  json.NewEncoder(w).Encode(new)

}

//Update User
func UpdateUser(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  params := mux.Vars(r)
  for index, item := range users{
    if item.ID == params["Id"]{
      users = append(users[:index], users[index+1:]...)
      var new User
      _ = json.NewDecoder(r.Body).Decode(&new)
      new.ID = params["Id"]
      users = append(users,new)
      json.NewEncoder(w).Encode(new)
      return

    }

  }
  json.NewEncoder(w).Encode(users)
}

//delete User
func DeleteUser(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  params := mux.Vars(r)
  for index, item := range users{
    if item.ID == params["Id"]{
      users = append(users[:index], users[index+1:]...)
      break
    }

  }
  json.NewEncoder(w).Encode(users)
}

//first Pagek
func HomePage(w http.ResponseWriter, r *http.Request){
  fmt.Fprint(w,"Hello this is HomePage on GO !")
}

func handleRequests(){

  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/",HomePage)
  myRouter.HandleFunc("/user",getAllUser).Methods("GET")
  myRouter.HandleFunc("/user/{Id}",GetUser).Methods("GET")
  myRouter.HandleFunc("/user/new",InsertUser).Methods("POST")
  myRouter.HandleFunc("/user/{Id}",UpdateUser).Methods("PUT")
  myRouter.HandleFunc("/user/{Id}",DeleteUser).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8080",myRouter))
}
func main(){

  users = append(users, User{
          ID : "1",
          Name: "Divya",
          Addsress: "xyz Road",
          MobileNumber: 999999999,
          City: "rajkot",
          Email: "divya@improwised.com",
  })

    handleRequests()
}
