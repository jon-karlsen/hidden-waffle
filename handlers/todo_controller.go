package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)


type TodoRequest struct {
    Desc string `json:"desc,omitempty"`
}


type TodoResponse struct {
    Id string `json:"id,omitempty"`
}


type TodoResource struct {
    db *gorm.DB
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


func ( ts TodoResource ) Index ( w http.ResponseWriter , r *http.Request ) {
    w.Write ( []byte ( "sdkljaskldjalksjdaksjdaskdj" ) )
}


func ( ts TodoResource ) CreateTodo ( w http.ResponseWriter , r *http.Request  ) {
    w.Write ( []byte ( "sdkljaskldjalksjdaksjdaskdj" ) )
}


func ( ts TodoResource ) DeleteTodo ( w http.ResponseWriter , r *http.Request ) {
    w.Write ( []byte ( "sdkljaskldjalksjdaksjdaskdj" ) )
}
