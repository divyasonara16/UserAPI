package main

import (
  "database/sql"
  "encoding/json"
  "fmt"
  "net/http"
  "strconv"
  "strings"

  _ "github.com/lib/pq"
)
//http://localhost:8080/company/?Id=3&Name=nara&Age=20&Address=rajkot&Salary=10000
type Company struct {
  Id       int    `json:"Id"`
  Name     string `json:"Name"`
  Age      int    `json:"Age"`
  Address  string `json:"Address"`
  Salary   int    `json:"Salary"`
  JoinDate string `json:"JoinDate"`
}

func dbConn() (db *sql.DB) {

  const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    password = "1234"
    dbname   = "postgres"
  )

  psqlInfo := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%d sslmode=disable",
    user, dbname, password, host, port)

  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err.Error())
  }

  fmt.Println("Database connection is succesful")
  return db
}

func Index(w http.ResponseWriter, r *http.Request) {

  w.Header().Set("Content-Type", "application/json")

  switch r.Method {
    // v1/account POST
    // v1/account GET

  case "GET":
    fmt.Println("======================")
    fmt.Println("Inside GET")
    fmt.Println("======================")
    get(w, r)
    w.WriteHeader(http.StatusOK)

  case "POST":
    insert(w, r)
    w.WriteHeader(http.StatusCreated)

  /*case "PUT":

    fmt.Println("======================")
    fmt.Println("Inside PUT")
    fmt.Println("======================")
    update(w, r)
    w.WriteHeader(http.StatusAccepted)*/

  case "DELETE":

    fmt.Println("======================")
    fmt.Println("Inside DELETE")
    fmt.Println("======================")
    delete(w, r)
    w.WriteHeader(http.StatusOK)

  default:
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"message": "not found"}`))

  }
}

func String2int(s string) int {

  i, err := strconv.Atoi(s)
  if err != nil {
    panic(err)
    return 0
  }
  return i
}

func insert(w http.ResponseWriter, r *http.Request) {
  db := dbConn()

  var c Company

  c.Id = String2int(r.FormValue("Id"))
  c.Name = r.FormValue("Name")
  c.Age = String2int(r.FormValue("Age"))
  c.Address = r.FormValue("Address")
  c.Salary = String2int(r.FormValue("Salary"))

  sqlStatement := `INSERT INTO company (id, name, age, address, salary) VALUES ($1, $2, $3, $4, $5)`

  _, err := db.Exec(sqlStatement, c.Id, c.Name, c.Age, c.Address, c.Salary)
  if err != nil {
    panic(err)
  }

  w.WriteHeader(http.StatusCreated)
  w.Write([]byte(`{"message": "success"}`))

  defer db.Close()
}

func struct2json(c Company) string {
  result, err := json.Marshal(c)
  if err != nil {
    panic(err)
    return "{}"
  }
  return string(result)
}

func handleGet(db *sql.DB, id int) *sql.Rows {
  sqlStatement := ``

  if id == -1 {
    sqlStatement = `SELECT * FROM COMPANY;`
    rows, _ := db.Query(sqlStatement)
    return rows
  } else if id > -1 {
    sqlStatement = `SELECT * FROM COMPANY WHERE id=$1;;`
    rows, _ := db.Query(sqlStatement, id)
    return rows
  } else {
    var rows *sql.Rows
    return rows
  }
}

func get(w http.ResponseWriter, r *http.Request) {

  db := dbConn()
  // var c Company
  id := 0
  if strings.Split(r.URL.Path, "/")[2] == "" {
    id = -1
  } else {
    // c.Id = String2int(strings.Split(r.URL.Path, "/")[2])
    id, _ = strconv.Atoi(strings.Split(r.URL.Path, "/")[2])
  }

  if id < -1 {
    w.Write([]byte(`{"message": "error"}`))
    defer db.Close()
    return
  }
  rows := handleGet(db, id)

  var cArr []Company
  var c Company
  for rows.Next() {
    err := rows.Scan(&c.Id, &c.Name, &c.Age, &c.Address, &c.Salary)
    if err != nil {
      // handle this error
      panic(err)
    }
   // c.Address = strings.Trim(c.Address, " ")
    cArr = append(cArr, c)
  }

  fmt.Println(cArr)


  if len(cArr) == 1 {
    result, _ := json.Marshal(cArr[0])
    w.Write([]byte(result))
  } else {
    result, _ := json.Marshal(cArr)
    w.Write([]byte(result))
  }

  defer db.Close()
}

/*func update(w http.ResponseWriter, r *http.Request) {

  fmt.Fprint(w,"Hello this is Update on GO !")
  db := dbConn()

  var c Company
  c.Id = strconv.Atoi(strings.Split(r.URL.Path, "/")[2])

  sqlStatement := `SELECT * FROM COMPANY WHERE id=$1;`
  row := db.QueryRow(sqlStatement, c.Id)

  switch err := row.Scan(&c.Id, &c.Name, &c.Age, &c.Address, &c.Salary); err {

  case sql.ErrNoRows:
    fmt.Println("No rows were returned!")
    w.Write([]byte(`{"message": "error"}`))

    var p Company

    // Try to decode the request body into the struct. If there is an error,
    // respond to the client with the error message and a 400 status code.
    err := json.NewDecoder(r.Body).Decode(&p)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)

        return
      }

      fmt.Println(p.Name)
      fmt.Println(p.Age)
      fmt.Println(p.Address)
      fmt.Println(p.Salary)
       /*c.Name = p.Name
         c.Age = p.Age
         c.Address = p.Address
         c.Salary = p.Salary*/

  /*  sqlStatement := `UPDATE COMPANY SET name = $2, age = $3, address = $4, salary = $5  WHERE id = $1;`

    _, err = db.Exec(sqlStatement, c.Id, c.Name, c.Age, c.Address, c.Salary)

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "success"}`))

  default:
    panic(err)
  }

  defer db.Close()
//
*/
func delete(w http.ResponseWriter, r *http.Request) {

  db := dbConn()

  var c Company
  c.Id = String2int(strings.Split(r.URL.Path, "/")[2])

  sqlStatement := `DELETE FROM COMPANY WHERE id=$1;`

  res, err := db.Exec(sqlStatement, c.Id)
  if err != nil {
    panic(err)
  }

      count, err := res.RowsAffected()
    if err != nil {
      panic(err)
    }
    fmt.Println(count)


  w.Write([]byte(`{"message": "success"}`))

  defer db.Close()
}

func main() {
  fmt.Println("Server started on: http://localhost:8080")
  http.HandleFunc("/company/", Index)
  http.ListenAndServe(":8080", nil)
}
