package main

import (
	"fmt"
	"hidden-waffle/data"
	"hidden-waffle/handlers"
	"hidden-waffle/web"
	"log"

	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)


func loadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Fatal( "Error loading .env file" )
    }
}


func initDb() *gorm.DB {
    db , err := data.Init()

    if err != nil {
        log.Fatal( fmt.Sprintf( "Database error: failed to connect.\n%s" , err ) )
    }

    return db
}


func main() {
    loadEnv ()

    db := initDb ()
    r  := chi.NewRouter ()

    r.Use ( middleware.RequestID )
    r.Use ( middleware.RealIP    )
    r.Use ( middleware.Logger    )
    r.Use ( middleware.Recoverer )

    r.Mount ( "/todo" , handlers.TodoResource{ Db: db }.Routes () )

    http.ListenAndServe ( ":8080" , r ))
}
