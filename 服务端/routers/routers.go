package routers

import (
	"fmt"
	"smartswitch/aliiot"

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
