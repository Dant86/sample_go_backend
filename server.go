package main

import (
    "SamplegoServer/api/routes"
    "net/http"
    "log"
    "SampleGoServer/db/utils"
    "SampleGoServer/db/dao"
    "fmt"
)

func main() {
    user_dao := dao.UserDao{}
    utils.Migrate("root", "dbpassword")
    id := user_dao.CreateUser("vedant", "pass")
    u := user_dao.ReadUser(1)
    http.HandleFunc("/users", routes.UserHandler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
