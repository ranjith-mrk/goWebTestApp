package main

import (
  _ "github.com/go-sql-driver/mysql"
  _ "github.com/jinzhu/gorm/dialects/mysql"
  "github.com/ranjith-mrk/goWebTestApp/models"
  "github.com/jinzhu/gorm"
)

func main() {
  db, _ := gorm.Open("mysql", "root:@/test?charset=utf8&parseTime=True&loc=Local")
  defer db.Close()

  db.AutoMigrate(&models.User{})
}