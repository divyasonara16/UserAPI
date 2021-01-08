# Create Simple User API in Golang
 In this file create simple server which can handle HTTP requests.and three distinc functions.

 * homepage : function that will handle requests to our root URl
 * handleRequest : function that will match the URL path hit with a defined function
 * main : function which starts API

 ```go
  package main

import (
    "fmt"
    "log"
    "net/http"
)

func HomePage(w http.ResponseWriter, r *http.Request){
  fmt.Println("Hello this is HomePage on GO !")
}


func handleRequests(){

  http.HandleFunc("/",HomePage)
  log.Fatal(http.ListenAndServe(":8080",nil))
}

func main() {
    handleRequests()
}
```
run this code by **go run BasicApi.go** and navigate to **http://localhost:8080/** in local browser so output should be **"Hello this is HomePage on Go !"** print out on screen.

# User Structure

Now creating a REST API that alloes to **Create**,**Read**, **Update**, and **Delete** the User.

First,define 'User' structure. in this structure Features like ID, Name, Addsress, MobileNumber, City, Email.

```go
  type User struct{
      ID string `json:"Id"`
      Name string `json:"name"`
      Addsress string `json:Addsress`
      MobileNumber int `json:MobileNumber`
      City string `json:city`
      Email string `json:email`
}

var users []User
```

To display all users,we need to import the **"encoding/json"** package

update **main** function to insert data for User

```go
  func main(){

  users = append(users, User{
          ID : "1",
          Name: "Divya",
          Addsress: "xyz Road",
          MobileNumber: 999999999,
          City: "rajkot",
          Email: "divya@improwised.com",
  })

  users = append(users, User{
          ID : "2",
          Name: "Sonara",
          Addsress: "ABC Road",
          MobileNumber: 888888888,
          City: "ahemdabad",
          Email: "sonara@improwised.com",
  })
  handleRequests()
}
```
# Retrieving All Users

Now when click with **HTTP GET** request, it will return all User for that define getAllUser function and set header ,which will return all user in JSON format:

```go
  func getAllUser(w http.ResponseWriter, r *http.Request){

  w.Header().Set("Content-Type","application/json")
  json.NewEncoder(w).Encode(users)
}
```
The call to json.NewEncoder(w).Encode(users) does encoding User array into a JSON string and then writing as part of response.

Before this will work, we’ll also need to add a new route to our handleRequests function that will map any calls to http://localhost:8080/user to newly defined function.

```go
  func handleRequests(){

  http.HandleFunc("/",HomePage)
  http.HandleFunc("/user",getAllUser)
  log.Fatal(http.ListenAndServe(":8080",nil))
}
```
open up http://localhost:8080/user in browser and you should see a JSON representation of your list of users like so:

```JSON
  [
    {
        "Id": "1",
        "name": "Divya",
        "Addsress": "xyz Road",
        "MobileNumber": 999999999,
        "City": "rajkot",
        "Email": "divya@improwised.com"
    },
    {
        "Id": "2",
        "name": "Sonara",
        "Addsress": "ABC Road",
        "MobileNumber": 888888888,
        "City": "ahemdabad",
        "Email": "sonara@improwised.com"
    }
]
```
Now to update API we must use **'github.com/gorilla/mux'** router.
this routers will enable to more easily perform tasks such as parsing any path or query parameters that may reside within an incoming HTTP request.

update import statement and modify **handleRrquests** function so that it creates a new router.

```go
  package main

import (
  "fmt"
  "log"
  "net/http"
  "encoding/json"
  "github.com/gorilla/mux"
)
unc handleRequests(){

  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/",HomePage)
  myRouter.HandleFunc("/user",getAllUser)

  log.Fatal(http.ListenAndServe(":8080",myRouter))
}
unc main(){

  users = append(users, User{
          ID : "1",
          Name: "Divya",
          Addsress: "xyz Road",
          MobileNumber: 999999999,
          City: "rajkot",
          Email: "divya@improwised.com",
  })

  users = append(users, User{
          ID : "2",
          Name: "Sonara",
          Addsress: "ABC Road",
          MobileNumber: 888888888,
          City: "ahemdabad",
          Email: "sonara@improwised.com",
  })
  handleRequests()
}

```
# Path Variables

Now using gorilla/mux we can add variables to our paths and select particular user.
To create a new route with **handlerequests()** function as:

```go
  myRouter.HandleFunc("/user/{Id}",GetUser).Methods("GET")
```
Now define Function that accept {id} and return particular User

```go
  unc GetUser(w http.ResponseWriter, r *http.Request){
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

```
Now type **'http://localhost:8080/user/1'** and get user which ID is 1
```go
  {
    "Id": "1",
    "name": "Divya",
    "Addsress": "xyz Road",
    "MobileNumber": 999999999,
    "City": "rajkot",
    "Email": "divya@improwised.com"
}
```
# Create and Update User

create new function **'CreateUser'** and also define route into **'handleRequests'** function

``` go
  func CreateUser(w http.ResponseWriter, r *http.Request){
  w.Header().Set("Content-Type","application/json")
  var new User
  _ = json.NewDecoder(r.Body).Decode(&new)
  new.ID = strconv.Itoa(rand.Intn(1000000))
  users = append(users,new)
  json.NewEncoder(w).Encode(new)

}
```
Here we use two funcitons rand of math/rand package and strcov package for string Convertion, here we genrating random Id for user.

we’ll be adding .Methods("POST") to the end of our route to specify that we only want to call this function when the incoming request is a HTTP POST request

```go
  func handleRequests(){

  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/",HomePage)
  myRouter.HandleFunc("/user",getAllUser).Methods("GET")
  myRouter.HandleFunc("/user/{Id}",GetUser).Methods("GET")
  myRouter.HandleFunc("/user/new",CreateUser).Methods("POST")
}
```
Now to update

```go
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
```
define handlerequest function

```go
  func handleRequests(){

  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/",HomePage)
  myRouter.HandleFunc("/user",getAllUser).Methods("GET")
  myRouter.HandleFunc("/user/{Id}",GetUser).Methods("GET")
  myRouter.HandleFunc("/user/new",CreateUser).Methods("POST")
  myRouter.HandleFunc("/user/{Id}",UpdateUser).Methods("PUT")

  log.Fatal(http.ListenAndServe(":8080",myRouter))
}
```

Pass Id which you want to update:
**http://localhost/user/2**

```go
  {
    "Id": "2",
    "name": "siya",
    "Addsress": "demo Road",
    "MobileNumber": 999999999,
    "City": "rajkot",
    "Email": "divya@improwised.com"
}
```
#Delete User

function receives HTTP DELETE requests and deletes user if they match the given Id path parameter.

```go
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
```
define route in handleRequests function.

```go
  func handleRequests(){

  myRouter := mux.NewRouter().StrictSlash(true)
  myRouter.HandleFunc("/",HomePage)
  myRouter.HandleFunc("/user",getAllUser).Methods("GET")
  myRouter.HandleFunc("/user/{Id}",GetUser).Methods("GET")
  myRouter.HandleFunc("/user/new",CreateUser).Methods("POST")
  myRouter.HandleFunc("/user/{Id}",UpdateUser).Methods("PUT")
  myRouter.HandleFunc("/user/{Id}",DeleteUser).Methods("DELETE")

  log.Fatal(http.ListenAndServe(":8080",myRouter))
}
```
HTTP DELETE request to http://localhost:8080/user/2. This will delete the second user within User array and Now when you go http://localhost:8080/user with a HTTP GET request, you should see second User will be deleted.
