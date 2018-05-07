package main

import(
  _ "github.com/go-sql-driver/mysql"
  "github.com/ranjith-mrk/goWebTestApp/models"
  "github.com/jinzhu/gorm"
)

func main() {
  db, _ := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
  defer db.Close()

  db.Create(&models.User{FirstName: "John", LastName: "Wick", Email: "johnwick@johnwick.com"})
}