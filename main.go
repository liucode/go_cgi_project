package main

import (
    "go_cgi_project/routers"
    _"go_cgi_project/models"
)

func main() {
    router := routers.InitRouter()
    //静态资源
    router.Static("/static", "./static")
    router.Run(":8084")
}
