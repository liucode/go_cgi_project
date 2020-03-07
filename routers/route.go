package routers

import (
    "github.com/gin-gonic/gin"
    "go_cgi_project/controllers"
)

func InitRouter() *gin.Engine {

    router := gin.Default()
    //注册：
    router.GET("/api/kv",controllers.KVGet)

    //router.Get("/api/data",controllers.DataGet)
    return router

}
