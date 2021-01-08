# Create Simple APi in Golang
 In this file create simple server which can handle HTTP requests.and three distinc functions.

 * homepage : function that will handle requests to our root URl
 * handleRequest : function that will hmatch the URL path hit with a defined function
 * main : function which starts API

 ```go
  package main

import (
    "fmt"
    "log"
    "net/http"
)

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

func main() {
    handleRequests()
}
```

