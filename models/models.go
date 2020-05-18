package models
import (
    "fmt"
    "go_cgi_project/common"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    _"sort"
    "encoding/json"
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


func XueliDb()  *gorm.DB{
    t_db, err := gorm.Open("mysql", "liu:123456@tcp(118.24.92.135:3306)/xueli?charset=utf8")
    if err != nil{
        fmt.Println("failed to connect database "+err.Error())
    }
    t_db.DB().SetMaxIdleConns(1000)
    t_db.DB().SetMaxOpenConns(1000)
    return t_db
}


func DB2NODE(data *common.JiaZu) *common.TreeNode{
    node := new(common.TreeNode)
    node.Name = data.Name
    node.Value = data.Other
    return node
}


func GetAll() (map[int][]int,map[int]*common.JiaZu){
    all := make([]*common.JiaZu,0)
    f_db := XueliDb()
    f_db.Table("jiazu").Find(&all)
    datamap := make(map[int]*common.JiaZu,0)
    for _,temp := range all{
        datamap[temp.Id] = temp
    }


    allmap := make(map[int][]int)
    for _,temp := range all{
        var c_list []int
        for _,temp2:=range all{
            if temp.Id == temp2.FatherId{
                c_list = append(c_list,temp2.Id)
            }
        }
        allmap[temp.Id] = c_list
    } 
    return allmap,datamap
}





func GetTreeByChildren(father *common.JiaZu) []*common.TreeNode {
    f_db := XueliDb()
    childrens:=make([]*common.JiaZu,0)
    childrens_node:=make([]*common.TreeNode,0)
    f_db.Table("jiazu").Where("father_id=?",father.Id).Find(&childrens)
    for _,children:=range childrens{
        children_node := DB2NODE(children)
        out_children := GetTreeByChildren(children)
        children_node.Children = out_children
        childrens_node = append(childrens_node,children_node)
    } 
    fmt.Println(childrens_node)
    return childrens_node
}

func GetTreeByName(name string) string{
    father := new(common.JiaZu)
    f_db := XueliDb()
    f_db.Table("jiazu").Where("name=?",name).First(father)
    father_node :=DB2NODE(father)
    father_node.Children = GetTreeByChildren(father)
    result, _ := json.Marshal(father_node)
    return string(result)
}


func GetTreeById2(id int) string{
    father := new(common.JiaZu)
    f_db := XueliDb()
    f_db.Table("jiazu").Where("id=?",id).First(father)
    father_node :=DB2NODE(father)
    father_node.Children = GetTreeByChildren(father)
    result, _ := json.Marshal(father_node)
    return string(result)
}

func GetTreeList(ids []int,allmap map[int][]int,datamap map[int]*common.JiaZu)[]*common.TreeNode{
    childrens_node:=make([]*common.TreeNode,0)
    for _,id := range ids{
        children_node :=DB2NODE(datamap[id])
        children_node.Children = GetTreeList(allmap[id],allmap,datamap)
        childrens_node = append(childrens_node,children_node)
    }
    return childrens_node
}

func GetTreeById(id int) string{
    allmap,datamap := GetAll()
    father_node :=DB2NODE(datamap[id])
    father_node.Children = GetTreeList(allmap[id],allmap,datamap)
    result, _ := json.Marshal(father_node)
    return string(result)
}

func GetV(appid int,k string) (v string){
    var data common.KV
    db.Table("kv").Where("appid=?",appid).Where("k=?",k).First(&data)
    fmt.Println("get result: ",data)
    v = data.V
    return 
}



func UpdateKV(kvcase common.KVCASE){
    var data common.KV
    data.Appid = kvcase.Appid
    data.K = kvcase.K
    data.V = kvcase.V


    fmt.Println("result: ",data)
    //先看在不在这个表里
    var count int
    db.Table("kv").Where("appid=?",kvcase.Appid).Where("k=?",kvcase.K).Count(&count)

    //不在表里，进行插入
    if count == 0 {
        err := db.Table("kv").Create(&data)
        if err!=nil{
            fmt.Println(err)
        }
    }else{
        //否则更新
        db.Table("kv").Where("appid=?",kvcase.Appid).Where("k=?",kvcase.K).Update(&data)
    }

    fmt.Println("update result: ",data)
    return 
}


//不直接传入结构体，防止以后变
func DeleteKV(appid int,k string){
    //先查出来id
    var data common.KV
    db.Table("kv").Where("appid=?",appid).Where("k=?",k).First(&data)
    //然后删除
    db.Table("kv").Delete(&data)

    fmt.Println("del result: ",data)
    return 
}




func GetVbyId(id int64) (v string){
    var data common.KV
    db.Table("kv").Where("id=?",id).First(&data)
    fmt.Println("get result: ",data)
    v = data.V
    return 
}



//不直接传入结构体，防止以后变
func DeleteKVbyId(id int64){
    //先查出来id
    var data common.KV
    data.ID = id
    //然后删除
    db.Table("kv").Delete(&data)

    fmt.Println("del result: ",data)
    return 
}



//获取一批数据
func GetData(appid int,subappid int) []*common.DATA{
    data :=make([]*common.DATA,0)
    db.Table("data").Where("appid=?",appid).Where("subappid=?",subappid).Find(&data)
    fmt.Println("get result: ",data)
    return data
}


//获取一个数据
func GetDatabyId(id int64) common.DATA{
    var data common.DATA
    db.Table("data").Where("id=?",).First(&data)
    fmt.Println("get result: ",data)
    return data
}

//创建一个新的
func CreateData(datacase common.DATACASE){
    var data common.DATA
    
    //转换数据
    data.Appid = datacase.Appid
    data.SubAppid = datacase.SubAppid
    data.Tempid = datacase.Tempid
    data.Title = datacase.Title
    data.Content = datacase.Content
    data.Pic = datacase.Pic
    data.Video = datacase.Video

 
    db.Table("data").Create(&data)
    fmt.Println("update result: ",data)
    return 
}

//更新数据
func UpdateData(datacase common.DATACASE){
    var data common.DATA
    
    //转换数据
    data.Appid = datacase.Appid
    data.SubAppid = datacase.SubAppid
    data.Tempid = datacase.Tempid
    data.Title = datacase.Title
    data.Content = datacase.Content
    data.Pic = datacase.Pic
    data.Video = datacase.Video

    db.Table("data").Where("id=?",datacase.ID).Update(&data)
    
    fmt.Println("update result: ",data)
    return 
}


//不直接传入结构体，防止以后变
func DeleteData(id int64){
    //先查出来id
    var data common.DATA
    data.ID = id
    //然后删除
    db.Table("data").Delete(&data)

    fmt.Println("del result: ",data)
    return 
}