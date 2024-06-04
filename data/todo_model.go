package data


import "gorm.io/gorm"


type Todo struct {
    gorm.Model

    Title string
    Desc  string
}
