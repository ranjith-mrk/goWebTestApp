package models


import (
  "github.com/jinzhu/gorm"
)

type User struct {
  gorm.Model
  FirstName string
  LastName string
  email string
}


func (u User) fullName() string {
  return (u.FirstName +  u.LastName)
}



