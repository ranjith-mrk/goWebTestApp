package main

import (
  // "encoding/json"
  // "errors"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  // "encoding/json"
  "github.com/jinzhu/gorm"
   _ "github.com/go-sql-driver/mysql"
  // "github.com/ranjith-mrk/goWebTestApp/models"
  h "github.com/ranjith-mrk/goWebTestApp/handlers"
  d "github.com/ranjith-mrk/goWebTestApp/database"
  // "controllers"
  // "html/template"
)

type Vegetables map[string]int

type Product struct {
  gorm.Model
  Code string
  Price uint
}


var product_data Product;

func main() {
  connectDataBase()
  registerRoutes()
}

func registerRoutes() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  r.HandleFunc("/users", h.UsersIndex)
  r.HandleFunc("/users/{id:[0-9]+}", h.UsersObject)
  // r.HandleFunc("/products", ProductsHandler)
  // r.HandleFunc("/articles", ArticlesHandler)
  http.Handle("/", r)
  http.ListenAndServe(":8080", nil)
}

func connectDataBase() {
  db, _ := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
  db.LogMode(true)
  d.DB = db
  // defer d.DB.Close()
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Home handler is called")

}