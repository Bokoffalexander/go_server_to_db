// Data Base connection postgresql
// ALTER USER user_name WITH SUPERUSER; 
package main

import (
  "database/sql"
  "fmt"
  "log"
  "strconv" // string to integer
  _ "github.com/lib/pq"
)

const (
  host     = "localhost"
  port     = 5432
  user     = "san"
  password = "san"
  dbname   = "san"
)

var Name string // Result of SQL query

func main() {
  server() // start web server
  // 1. Book_id from http query
  // 2. start db connect to postgresql //myDBconnect()
  // 3. myDBconnect() with Book_id query
  // 4. This is ready result for sql query //Name
  // 5. Server вписывает sql ответ в браузер
  
  
}


func myDBconnect(){
  psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
  db, err := sql.Open("postgres", psqlInfo)
  if err != nil {
    panic(err)
  }
  defer db.Close()

  err = db.Ping()
  if err != nil {
    panic(err)
  }

  fmt.Println("Successfully connected!")
  
  var s sql.NullString // sql to string
  
  strVar := Book_id // global Book_id is "string" from server.go
  intVar, err := strconv.Atoi(strVar)
  fmt.Println("Looking for book_id =", intVar, err)

  id := intVar // sql query for this = intVar is book_id
  if id == 0 { // no entered url yet
    id = 1 // initial will be id=1
  }

  err1 := db.QueryRow("SELECT title FROM book where book_id=$1;", id).Scan(&s)
  if err1 != nil {
    log.Fatal(err1)
  }

  // Находим имя и печатаем на экран
  name := "Printing result of s"
  if s.Valid {
    name = s.String
  }
  fmt.Println(name)
  Name = name // global

}

