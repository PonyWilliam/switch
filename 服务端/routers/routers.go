package routers

import (
	"fmt"
	"math/rand"
	"smartswitch/aliiot"
	"smartswitch/message"
	"smartswitch/redis"
	"smartswitch/users"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetAllDevices(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": aliiot.AllDevices,
	})
}
func ReflashAllDevices(c *gin.Context) {
	aliiot.Reflash()
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
	})
}

func SetSwitch(c *gin.Context) {
	key := c.PostForm("key")
	name := c.PostForm("name")
	value := c.PostForm("value")
	fmt.Println("set method")
	err := aliiot.Smart_setDeviceProperty(name, key, "IsOpen", value)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "设置成功",
	})
}

func GetSwitch(c *gin.Context) {
	key := c.PostForm("key")
	name := c.PostForm("name")
	ress := aliiot.Getproperty(name, key)
	var res string
	for _, v := range ress {
		fmt.Println(v)
		if v.Indentifier == "IsOpen" {

			if v.Value.Data > 1 || v.Value.Data < 0 {
				res = "-1"
			} else {
				if v.Value.Data == 1 {
					res = "1"
				} else {
					res = "0"
				}
			}
		}
	}
	if res == "-1" {
		c.JSON(200, gin.H{
			"code": 200,
			"data": "未知",
		})
		return
	}
	fmt.Println("res:" + res)
	if res == "1" {
		c.JSON(200, gin.H{
			"code": 200,
			"data": "开",
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"data": "关",
		})
	}
}

func SetocValue(c *gin.Context) {
	key := c.PostForm("key")
	name := c.PostForm("name")
	oc := c.PostForm("oc")
	value := c.PostForm("value")
	var err error
	if oc == "c" {
		err = aliiot.Smart_setDeviceProperty(name, key, "cvalue", value)
	} else {
		err = aliiot.Smart_setDeviceProperty(name, key, "ovalue", value)
	}
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "设置成功",
	})
}

func GetocValue(c *gin.Context) {
	key := c.PostForm("key")
	name := c.PostForm("name")
	oc := c.PostForm("oc")
	var result int64
	if oc == "c" {
		ress := aliiot.Getproperty(name, key)
		for _, v := range ress {
			if v.Indentifier == "cvalue" {
				result = int64(v.Value.Data)
			}
		}
	} else {
		ress := aliiot.Getproperty(name, key)
		for _, v := range ress {
			if v.Indentifier == "ovalue" {
				result = int64(v.Value.Data)
			}
		}
	}
	c.JSON(200, gin.H{
		"code": 200,
		"data": result,
	})
}

func CheckCode(c *gin.Context) {
	//校验是否能登录
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
	})
}

func GetCode(c *gin.Context) {
	//用于生成验证码，后续通过验证码登录或者注册账号
	phone := c.PostForm("phone")
	reqIP := c.ClientIP()
	if reqIP == "::1" {
		reqIP = "127.0.0.1"
	}
	if !redis.CheckIp(reqIP) {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "申请验证码需要间隔1分钟",
		})
		return
	}
	var temp string
	temp = ""
	for i := 0; i < 6; i++ {
		temp += strconv.FormatInt(rand.Int63n(10), 10)
	}
	//生成随机码，调用阿里云接口发送到用户手机
	err := message.SendMsg(phone, temp)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	//发送成功存储到redis
	redis.SetUser(phone, reqIP, temp)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "发送验证码成功",
	})
}

func Registry(c *gin.Context) {
	//用于注册账号
	username := c.PostForm("username")
	password := c.PostForm("password")
	//验证code与phone，组合判断是否能注册
	phone := c.PostForm("phone")
	code := c.PostForm("code")
	//这个code需要调用redis获取
	if len(username) < 4 {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "账号名需大于4位",
		})
		return
	}
	if len(password) < 6 {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "密码需要大于6位",
		})
		return
	}
	temp_code, err := redis.GetUser(phone)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	if code != temp_code {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "验证码错误",
		})
		return

	}
	err = users.Registry(username, password, phone)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	redis.DelUser(phone)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}
func Login(c *gin.Context) {
	user := c.PostForm("username")
	pwd := c.PostForm("password")
	ok, res := users.Login(user, pwd)
	if !ok {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "账号或密码错误",
		})
		return
	}
	if len(res) != 1 {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "未知错误:" + strconv.FormatInt(int64(len(res)), 10),
		})
		return
	}
	//返回账号token
	temp := NewJWT()
	claims := &CustomClaims{
		Id:       res[0].Id,
		Username: res[0].Username,
		Phone:    res[0].Phone,
	}
	res2, err := temp.CreateToken(*claims)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "创建token时出错:" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "登录成功",
		"token": res2,
	})
}
func ReflashToken(c *gin.Context) {
	//刷新登录的token
	token := c.PostForm("token")
	temp := NewJWT()
	res, err := temp.RefreshToken(token)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  "刷新token出错:" + err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code":  200,
		"msg":   "刷新成功",
		"token": res,
	})
}

//给hyh暴露的2个set、get接口
func SetByID(c *gin.Context) {
	key := c.Param("id")
	name := c.Query("name")
	val := c.Query("val")
	err := aliiot.Smart_setDeviceProperty(name, key, "IsOpen", val)
	if err != nil {
		c.JSON(200, gin.H{
			"code": 500,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "设置成功",
	})
}
func GetByID(c *gin.Context) {
	key := c.Param("id")
	name := c.Query("name")
	ress := aliiot.Getproperty(name, key)
	var res string
	for _, v := range ress {
		fmt.Println(v)
		if v.Indentifier == "IsOpen" {

			if v.Value.Data > 1 || v.Value.Data < 0 {
				res = "-1"
			} else {
				if v.Value.Data == 1 {
					res = "1"
				} else {
					res = "0"
				}
			}
		}
	}
	if res == "-1" {
		c.JSON(200, gin.H{
			"code": 200,
			"data": -1,
		})
		return
	}
	fmt.Println("res:" + res)
	if res == "1" {
		c.JSON(200, gin.H{
			"code": 200,
			"data": 1,
		})
	} else {
		c.JSON(200, gin.H{
			"code": 200,
			"data": 0,
		})
	}
}
