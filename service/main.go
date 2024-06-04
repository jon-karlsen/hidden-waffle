package service


type Todo struct {
    Id   string `json:"id,omitempty"`
    Desc string `json:"desc,omitempty"`
}


func CreateTodo( desc string ) string {
}
