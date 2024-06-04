package data


import (
    "fmt"
    "os"
    "gorm.io/gorm"
    "gorm.io/driver/sqlite"

    _ "github.com/tursodatabase/libsql-client-go/libsql"
)


func Init() ( *gorm.DB , error ) {
    url      := fmt.Sprintf( "%s?authToken=%s" , os.Getenv( "TURSO_DATABASE_URLL" ) , os.Getenv( "TURSO_AUTH_TOKEN" ) )
    db , err := gorm.Open( sqlite.Open( url ) , &gorm.Config{} )

    if err != nil {
        return nil , err
    }

    return db , nil
}
