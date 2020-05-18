package controllers

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "fmt"
    "go_cgi_project/common"
    "go_cgi_project/models"
)


func TreeGetByName(c *gin.Context){
    var jzcase common.JIAZUCASE
    if c.ShouldBind(&jzcase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    if jzcase.Name == "" {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    outv := models.GetTreeByName(jzcase.Name)
    c.String(http.StatusOK,string(outv))
}


func TreeGetById(c *gin.Context){
    var jzcase common.JIAZUCASE
    if c.ShouldBind(&jzcase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    fmt.Println(jzcase)
    if jzcase.ID == 0{
        jzcase.ID = 1
    }
    outv := models.GetTreeById(jzcase.ID)
    c.String(http.StatusOK,string(outv))
}


//KV TABLE
func KVGet(c *gin.Context){
    var kvcase common.KVCASE
    if c.ShouldBind(&kvcase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    if kvcase.Appid==0 || kvcase.K=="" {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    fmt.Println(kvcase)
    outv := models.GetV(kvcase.Appid,kvcase.K)
    c.String(http.StatusOK,string(outv))
}


func KVPost(c *gin.Context){
    var kvcase common.KVCASE
    if c.ShouldBind(&kvcase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    if kvcase.Appid==0 || kvcase.K==""|| kvcase.V==""{
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    models.UpdateKV(kvcase)
    c.String(http.StatusOK,string(0))
}


func KVDelete(c *gin.Context){
    var kvcase common.KVCASE
    if c.ShouldBind(&kvcase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    if kvcase.Appid==0 || kvcase.K==""{
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    models.DeleteKV(kvcase.Appid,kvcase.K)
    c.String(http.StatusOK,string(0))
}

//KV TABLE BY ID
func KVGetbyId(c *gin.Context){
    var kvcase common.KVCASE
    if c.ShouldBind(&kvcase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    if kvcase.ID==0{
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    fmt.Println(kvcase)
    outv := models.GetVbyId(kvcase.ID)
    c.String(http.StatusOK,string(outv))
}




func KVDeletebyId(c *gin.Context){
    var kvcase common.KVCASE
    if c.ShouldBind(&kvcase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    if kvcase.ID==0{
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    models.DeleteKVbyId(kvcase.ID)
    c.String(http.StatusOK,string(0))
}




//Data -->多值的我直接传个结构体去前端
func DataGet(c *gin.Context){
    var datacase common.DATACASE
    if c.ShouldBind(&datacase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    if datacase.Appid==0 || datacase.SubAppid==0{
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    outlist := models.GetData(datacase.Appid,datacase.SubAppid)
    c.JSON(http.StatusOK,outlist)
}


//只查一个
func DataGetbyId(c *gin.Context){
    var datacase common.DATACASE
    if c.ShouldBind(&datacase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    if datacase.ID==0{
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    outlist := models.GetDatabyId(datacase.ID)
    c.JSON(http.StatusOK,outlist)
}


//新建数据
func DataCreate(c *gin.Context){
    var datacase common.DATACASE
    if c.ShouldBind(&datacase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    //todo 需要加限制
    models.CreateData(datacase)
    c.String(http.StatusOK,string(0))
}




//更新数据
func DataUpdate(c *gin.Context){
    var datacase common.DATACASE
    if c.ShouldBind(&datacase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    if datacase.ID==0{
        c.String(http.StatusBadRequest,string("参数错误"))
        return  
    }
    //TODO 需要加限制
    models.UpdateData(datacase)
    c.String(http.StatusOK,string(0))
}

func DataDelete(c *gin.Context){
    var datacase common.DATACASE
    if c.ShouldBind(&datacase) != nil {
        c.String(http.StatusBadRequest,string("参数错误"))
        return
    }
    if datacase.ID==0{
        c.String(http.StatusBadRequest,string("参数错误"))
        return  
    }
    models.DeleteData(datacase.ID)
    c.String(http.StatusOK,string(0))
}