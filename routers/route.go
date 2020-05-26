package routers

import (
	"go_cgi_project/controllers"

	"go_cgi_project/common"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	v1 := router.Group("/api")
	//通过appid+key可以唯一定位
	v1.GET("/kvGet", controllers.KVGet)
	v1.POST("/kvPost", controllers.KVPost)
	v1.DELETE("/kvDelete", controllers.KVDelete)

	v1.GET("/kvGetbyId", controllers.KVGetbyId)
	//v1.POST("/kvUpdatebyId",controllers.KVPostbyId) //kv的更新不提供根据id
	v1.DELETE("/kvDeletebyId", controllers.KVDeletebyId)

	v1.GET("/dataGet", controllers.DataGet)
	v1.GET("/dataGetbyId", controllers.DataGetbyId)
	v1.POST("/dataUpdate", controllers.DataUpdate)
	v1.POST("/dataCreate", controllers.DataCreate)
	v1.DELETE("/dataDelete", controllers.DataDelete)

	v1.GET("/treeGetbyId", controllers.TreeGetById)
	v1.GET("/treeGetbyName", controllers.TreeGetByName)

	v1.POST("/login", controllers.Login)
	v1.POST("/register", controllers.RegisterUser)

	taR := router.Group("/data")
	taR.Use(common.JWTAuth())
	{
		taR.GET("/dataByTime", controllers.GetDataByTime)
	}
	//router.Get("/api/data",controllers.DataGet)
	return router

}
