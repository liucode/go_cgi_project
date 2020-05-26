package controllers

import (
	"fmt"
	"go_cgi_project/common"
	"go_cgi_project/models"
	"log"
	"net/http"
	"time"

	myjwt "go_cgi_project/common"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func TreeGetByName(c *gin.Context) {
	var jzcase common.JIAZUCASE
	if c.ShouldBind(&jzcase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	if jzcase.Name == "" {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	outv := models.GetTreeByName(jzcase.Name)
	c.String(http.StatusOK, string(outv))
}

func TreeGetById(c *gin.Context) {
	var jzcase common.JIAZUCASE
	if c.ShouldBind(&jzcase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	fmt.Println(jzcase)
	if jzcase.ID == 0 {
		jzcase.ID = 1
	}
	outv := models.GetTreeById(jzcase.ID)
	c.String(http.StatusOK, string(outv))
}

//KV TABLE
func KVGet(c *gin.Context) {
	var kvcase common.KVCASE
	if c.ShouldBind(&kvcase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	if kvcase.Appid == 0 || kvcase.K == "" {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	fmt.Println(kvcase)
	outv := models.GetV(kvcase.Appid, kvcase.K)
	c.String(http.StatusOK, string(outv))
}

func KVPost(c *gin.Context) {
	var kvcase common.KVCASE
	if c.ShouldBind(&kvcase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	if kvcase.Appid == 0 || kvcase.K == "" || kvcase.V == "" {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	models.UpdateKV(kvcase)
	c.String(http.StatusOK, string(0))
}

func KVDelete(c *gin.Context) {
	var kvcase common.KVCASE
	if c.ShouldBind(&kvcase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	if kvcase.Appid == 0 || kvcase.K == "" {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	models.DeleteKV(kvcase.Appid, kvcase.K)
	c.String(http.StatusOK, string(0))
}

//KV TABLE BY ID
func KVGetbyId(c *gin.Context) {
	var kvcase common.KVCASE
	if c.ShouldBind(&kvcase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	if kvcase.ID == 0 {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	fmt.Println(kvcase)
	outv := models.GetVbyId(kvcase.ID)
	c.String(http.StatusOK, string(outv))
}

func KVDeletebyId(c *gin.Context) {
	var kvcase common.KVCASE
	if c.ShouldBind(&kvcase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	if kvcase.ID == 0 {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	models.DeleteKVbyId(kvcase.ID)
	c.String(http.StatusOK, string(0))
}

//Data -->多值的我直接传个结构体去前端
func DataGet(c *gin.Context) {
	var datacase common.DATACASE
	if c.ShouldBind(&datacase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	if datacase.Appid == 0 || datacase.SubAppid == 0 {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	outlist := models.GetData(datacase.Appid, datacase.SubAppid)
	c.JSON(http.StatusOK, outlist)
}

//只查一个
func DataGetbyId(c *gin.Context) {
	var datacase common.DATACASE
	if c.ShouldBind(&datacase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	if datacase.ID == 0 {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	outlist := models.GetDatabyId(datacase.ID)
	c.JSON(http.StatusOK, outlist)
}

//新建数据
func DataCreate(c *gin.Context) {
	var datacase common.DATACASE
	if c.ShouldBind(&datacase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	//todo 需要加限制
	models.CreateData(datacase)
	c.String(http.StatusOK, string(0))
}

//更新数据
func DataUpdate(c *gin.Context) {
	var datacase common.DATACASE
	if c.ShouldBind(&datacase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	if datacase.ID == 0 {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	//TODO 需要加限制
	models.UpdateData(datacase)
	c.String(http.StatusOK, string(0))
}

func DataDelete(c *gin.Context) {
	var datacase common.DATACASE
	if c.ShouldBind(&datacase) != nil {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	if datacase.ID == 0 {
		c.String(http.StatusBadRequest, string("参数错误"))
		return
	}
	models.DeleteData(datacase.ID)
	c.String(http.StatusOK, string(0))
}

// 注册信息
type RegistInfo struct {
	// 手机号
	Phone string `json:"mobile"`
	// 密码
	Pwd string `json:"pwd"`
}

// Register 注册用户
func RegisterUser(c *gin.Context) {
	var registerInfo RegistInfo
	if c.BindJSON(&registerInfo) == nil {
		err := models.Register(registerInfo.Phone, registerInfo.Pwd)
		if err == nil {
			c.JSON(http.StatusOK, gin.H{
				"status": 0,
				"msg":    "注册成功！",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "注册失败" + err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "解析数据失败！",
		})
	}
}

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	models.User
}

// Login 登录
func Login(c *gin.Context) {
	var loginReq models.LoginReq
	if c.BindJSON(&loginReq) == nil {
		isPass, user, err := models.LoginCheck(loginReq)
		if isPass {
			generateToken(c, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "验证失败," + err.Error(),
			})
		}
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    "解析数据失败",
		})
	}
}

// 生成令牌
func generateToken(c *gin.Context, user models.User) {
	j := &myjwt.JWT{
		[]byte("corey"),
	}
	claims := myjwt.CustomClaims{
		user.UserId,
		user.Username,
		user.Usermobile,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "corey",                         //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	data := LoginResult{
		User:  user,
		Token: token,
	}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   data,
	})
	return
}

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}
