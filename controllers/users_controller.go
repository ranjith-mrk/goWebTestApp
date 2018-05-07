package controllers

import (
  "fmt"
  m "github.com/ranjith-mrk/goWebTestApp/models"
  d "github.com/ranjith-mrk/goWebTestApp/database"
  "encoding/json"
)

func UsersIndex() []byte {
  var users  = []m.User{}

  // d.DB.First(&user, userId)
  d.DB.Find(&users)


  jData, err := json.Marshal(users)
  if err != nil {
      panic(err)
      return nil
  }
  fmt.Printf("jdata", jData)
  // w.Header().Set("Content-Type", "application/json")
  // w.Write(jData)

  fmt.Printf("%+v\n", users)
  return jData
}

func UserDetails(userId int) []byte {
  var user m.User
  fmt.Printf("User id is", userId)
  d.DB.First(&user, userId)

  jData, err := json.Marshal(user)
  if err != nil {
      panic(err)
      return nil
  }
  fmt.Printf("jdata", jData)
  // w.Header().Set("Content-Type", "application/json")
  // w.Write(jData)

  fmt.Printf("%+v\n", user)
  return jData
}