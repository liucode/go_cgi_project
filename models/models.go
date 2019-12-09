package models
import (
	_ "github.com/go-sql-driver/mysql"
	_"fmt"
	db "../database"
	_"database/sql"
	_"sort"
)


func QueryRows(sql string) [4]int{
    var mytype string
    var count int
    var cmap [4]int
    rows, err := db.SqlDB.Query(sql)
    if err != nil {
        panic(err.Error())
    }
    i:= 0
    for rows.Next(){
        err := rows.Scan(&mytype,&count)
    if err != nil {
        panic(err.Error())
    }
      cmap[i]=count
      i++
      if i>=4 {
       break
       }
 }   
 return cmap
}
func QueryCount(sql string) int{
    var outcount int
    rows, err := db.SqlDB.Query(sql)
    if err != nil {
        panic(err.Error())
    }
    for rows.Next() {
    err := rows.Scan(&outcount)
    if err != nil {
        panic(err.Error())
    }
    }
    return outcount

}
func ExistZero() (int,int){
    //查询
    zcsql := "select count(1) from zchang where price=0;"
    zczero := QueryCount(zcsql)
    
    ecsql := "select count(1) from echang where price=0;"
    eczero := QueryCount(ecsql)
    return zczero,eczero
}

func ExistD() (int,int){
    //查询
    zcsql := "select count(1) from zchang where price>0 and price<100;"
    zczero := QueryCount(zcsql)
    
    ecsql := "select count(1) from echang where price>0 and price<100;"
    eczero := QueryCount(ecsql)
    return zczero,eczero
}

func ExistZ() (int,int){
    //查询
    zcsql := "select count(1) from zchang where price<1000 and price>=100;"
    zczero := QueryCount(zcsql)
    
    ecsql := "select count(1) from echang where price<1000 and price>=100;"
    eczero := QueryCount(ecsql)
    return zczero,eczero
}

func ExistG() (int,int){
    //查询
    zcsql := "select count(1) from zchang where price>=1000;"
    zczero := QueryCount(zcsql)
    
    ecsql := "select count(1) from echang where price>=1000;"
    eczero := QueryCount(ecsql)
    return zczero,eczero
}

func GetType() ([4]int,[4]int){
    zcsql := "select type,count(1) from zchang group by type order by type;"
    zcmap := QueryRows(zcsql)
    
    ecsql := "select type,count(1) from echang group by type order by type;"
    ecmap := QueryRows(ecsql)
    return zcmap,ecmap
}
