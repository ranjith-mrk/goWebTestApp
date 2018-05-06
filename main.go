package main

import (
  // "encoding/json"
  // "errors"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  "github.com/jinzhu/gorm"
   _ "github.com/go-sql-driver/mysql"
  _ "github.com/jinzhu/gorm/dialects/mysql"
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
  registerRoutes()

}






func registerRoutes() {
  r := mux.NewRouter()
  r.HandleFunc("/", HomeHandler)
  // r.HandleFunc("/products", ProductsHandler)
  // r.HandleFunc("/articles", ArticlesHandler)
  http.Handle("/", r)
  http.ListenAndServe(":8080", nil)
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Printf("Home handler is called")
  vegetables := make(map[string]int)
  vegetables["Carrats"] = 10
  vegetables["Beets"] = 0
  jData, err := json.Marshal(vegetables)
  if err != nil {
      panic(err)
      return
  }

  fmt.Printf("jdata", jData)
  // t := template.New("some template")
  // t, _ = t.ParseFiles("views/dashboards/index.html")
  // t.Execute(w, nil)
  w.Header().Set("Content-Type", "application/json")
  w.Write(jData)

    db, _ := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
  defer db.Close()

    // Migrate the schema
  db.AutoMigrate(&Product{})

  // Create
  db.Create(&Product{Code: "L1212", Price: 1000})

  db.First(&product_data, 1) // find product with id 1

  fmt.Printf("%+v\n", product_data)
  
}