package models
import (
    "fmt"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
	_"sort"
)


var db  *gorm.DB

func init() {
    t_db, err := gorm.Open("mysql", "root:661020@tcp(127.0.0.1:3306)/base?charset=utf8&parseTime=True&loc=Local")
    if err != nil{
        panic("failed to connect database")
    }
    t_db.DB().SetMaxIdleConns(10)
    t_db.DB().SetMaxOpenConns(100)
    db = t_db //全局变量
}


type KV struct{
    ID int `gorm:"type:int(20);column:id;primary_key"`
    Appid int `gorm:"type:int(11);column:appid`
    K string `gorm:"type:varchar(255);column:k`
    V string `gorm:"type:text;column:v`
}

type DATA struct {
    ID int64 `gorm:"type:bigint(20);column:id;primary_key"`
    Appid int `gorm:"type:int(11);column:appid`
    SubAppid int `gorm:"type:int(11);column:subappid`
    Tempid int `gorm:"type:int(11);column:tempid`
    Sortid int `gorm:"type:int(11);column:sortid`
    Title string `gorm:"type:text;column:title`
    Content string `gorm:"type:text;column:content`
    Pic string `gorm:"type:text;column:pic`
    Video string `gorm:"type:text;column:video`
}


func GetV(appid int,k string) (v string){
    var data KV
    fmt.Println(db)
    db.Table("kv").Where("appid=?",appid).Where("k=?",k).First(&data)
    fmt.Println("result: ",data)
    v = data.V
    return 
}
