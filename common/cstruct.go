package common
type KVCASE struct {
    ID int64 `form:"id"`
    Appid       int    `form:"appid"`
    K    string    `form:"k"`
    V    string   `form:"v"`
}

type DATACASE struct{
    ID int64 `form:"id"`
    Appid int `form:appid`
    SubAppid int `form:subappid`
    Tempid int `form:tempid`
    Sortid int `form:sortid`
    Title string `form:title`
    Content string `form:content`
    Pic string `form:pic`
    Video string `form:video`
}

type KV struct{
    ID int64 `gorm:"type:int(20);column:id;primary_key"`
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