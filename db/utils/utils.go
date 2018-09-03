package utils

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gchaincl/dotsql"
)

func OpenMySQL(uname, pass string) *sql.DB {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/test?parseTime=true",
                        uname, pass))
    if err != nil {
        panic(err.Error())
    }
    return db
}

func Migrate(uname, pass string) {
    db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/",
                        uname, pass))
    err = db.Ping()
    if err != nil {
        panic(err.Error())
    }
    db.Exec("DROP DATABASE test")
    db.Exec("CREATE DATABASE test")
    db.Exec("USE test")
    dot, err := dotsql.LoadFromFile("db/migrations.sql")
    if err != nil {
        panic(err.Error())
    }
    fmt.Println("Applying Migrations...")
    fmt.Println("Migrating table users...")
    dot.Exec(db, "users-table")
    fmt.Println("Migrating table user_auth...")
    dot.Exec(db, "user-auth-table")
    fmt.Println("Migrating table posts...")
    dot.Exec(db, "posts-table")
    fmt.Println("Migrating table likes...")
    dot.Exec(db, "likes-table")
    fmt.Println("Migrating table comments...")
    dot.Exec(db, "comments-table")
    fmt.Println("Successfully applied migrations.")
}
