package main

import (
  // "encoding/json"
  // "errors"
  "fmt"
  "net/http"
  "github.com/gorilla/mux"
  "encoding/json"
  // "controllers"
  // "html/template"
)

type Vegetables map[string]int

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
}