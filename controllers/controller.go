package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "../models"
    _"github.com/chenjiandongx/go-echarts/charts"
    _"log"
    _"math/rand"
    _"os"
    "encoding/json"
    _"fmt"
)
type Message struct {
    Head string
    Text string
}
func TableGet(c *gin.Context) {
    var l []Message
    l  = append(l,Message{"地址","我是一个测试"})
    l = append(l,Message{"电话","18181818181"})
    b,_:=json.Marshal(l)
    
    c.String(http.StatusOK,string(b))
}
func  DetailGet(c *gin.Context) {
    //返回html
    c.HTML(http.StatusOK,"detail.html",gin.H{"title":"注册页"})
}

func  IndexGet(c *gin.Context) {
    //返回html
    var zcz [4] int;
    var ecz [4] int;
    zcz[0],ecz[0] = models.ExistZero()
    zcz[1],ecz[1] = models.ExistD()
    zcz[2],ecz[2] = models.ExistZ()
    zcz[3],ecz[3] = models.ExistG()
    zcm,ecm := models.GetType()
    
    c.HTML(http.StatusOK,"index.html",gin.H{"zcz":zcz,"ecz":ecz,"zcm":zcm,"ecm":ecm})
}

