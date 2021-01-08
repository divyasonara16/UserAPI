#Create Simple APi in Golang
 In this file create simple server which can handle HTTP requests.and three distinc functions.

 *homepage : function that will handle requests to our root URl
 *handleRequest : function that will hmatch the URL path hit with a defined function
 *main : function which starts API

 ```go
  package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func main() {
    handleRequests()
}
```

