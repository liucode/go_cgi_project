package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "go_cgi_project/models"
    "fmt"
)


type KVCASE struct {
    Appid       int    `form:"appid"`
    K    string    `form:"k"`
}

func KVGet(c *gin.Context){
    var kvcase KVCASE
    var v string
    if c.ShouldBind(&kvcase) == nil {
        fmt.Println(kvcase)
        v = models.GetV(kvcase.Appid,kvcase.K)
        v = "123123123"
    }
    c.String(http.StatusOK,string(v+kvcase.K))
}
