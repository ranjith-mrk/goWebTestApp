package handlers

import (
  c "github.com/ranjith-mrk/goWebTestApp/controllers"
  "net/http"
  "github.com/gorilla/mux"
  "strconv"
)

func UsersIndex(w http.ResponseWriter, r *http.Request) {
  var jData = c.UsersIndex()
  //Passing user params
  w.Header().Set("Content-Type", "application/json")
  w.Write(jData)
}


func UsersObject(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  var user_id, _= strconv.Atoi(vars["id"])
  var jData = c.UserDetails(user_id)
  //Passing user params
  w.Header().Set("Content-Type", "application/json")
  w.Write(jData)
}