package main

import (
    "../myproject/routers"
    db "../myproject/database"
)

func main() {
    defer db.SqlDB.Close()
    router := routers.InitRouter()
    //静态资源
    router.Static("/static", "./static")
    router.Run(":8084")
}
