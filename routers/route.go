package routers

import (
    "github.com/gin-gonic/gin"
    "../controllers"
)

func InitRouter() *gin.Engine {

    router := gin.Default()
    router.LoadHTMLGlob("views/*")
    //注册：
    router.GET("/",controllers.IndexGet)
    router.GET("/api/table",controllers.TableGet)
    router.GET("/detail",controllers.DetailGet)
    return router

}
