package handler


import (
	"hidden-waffle/service"
	"net/http"

	"github.com/go-chi/chi/v5"
)


type TodoRequest struct {
    Desc string `json:"desc,omitempty"`
}


type TodoResponse struct {
    Id string `json:"id,omitempty"`
}


type TodoResource struct {
    ts service.TodoService
}


func ( ts TodoResource ) Routes() chi.Router {
    r := chi.NewRouter()

    r.Get  ( "/" , ts.Index )
    r.Post ( "/" , ts.CreateTodo )

    r.Route ( "/{id}" , func ( r chi.Router ) {
        r.Delete ( "/" , ts.DeleteTodo )
    } )

    return r
}


func ( tr TodoResource ) Index ( w http.ResponseWriter , r *http.Request ) {
    res , err := tr.ts.Index ()
    w.Write ( []byte ( "sdkljaskldjalksjdaksjdaskdj" ) )
}


func ( tr TodoResource ) CreateTodo ( w http.ResponseWriter , r *http.Request  ) {
    res , err := tr.ts.Index ()
    w.Write ( []byte ( "sdkljaskldjalksjdaksjdaskdj" ) )
}


func ( tr TodoResource ) DeleteTodo ( w http.ResponseWriter , r *http.Request ) {
    res , err := tr.ts.Index ()
    w.Write ( []byte ( "sdkljaskldjalksjdaksjdaskdj" ) )
}
