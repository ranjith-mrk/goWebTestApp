package models

type User struct {
  gorm.Model
  FirstName string
  LastName string
  email string
}


func (u User) fullName() string {
  return (u.FirstName +  u.LastName)
}



