package service


import "gorm.io/gorm"


type TodoService struct {
    Db *gorm.DB
}


func ( ts TodoService ) Index ( todo *Todo ) ( *Todo , error ) {
    return todo
}


func ( ts TodoService ) CreateTodo ( todo *Todo ) ( *Todo , error ) {
    return todo
}


func ( ts TodoService ) DeleteTodo ( todo *Todo ) ( *Todo , error ) {
    return todo
}
