package main

//iot层放在./aliiot/aliiot.go中
//业务层相关放在./routers/routers.go,包名routers
import (
	aliiot "smartswitch/aliiot"
	"smartswitch/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	aliiot.Ali_init()
	r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowOriginFunc:  func(origin string) bool { return true },
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "token"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "站点正在正常运行",
		})
	})
	r.GET("/devices", routers.GetAllDevices) //获取设备列表
	r.POST("/reflash", routers.ReflashAllDevices)
	group1 := r.Group("/get")
	{
		group1.POST("/", routers.GetSwitch)
		group1.POST("/switch", routers.GetSwitch)
		group1.POST("/ocvalue", routers.GetocValue)
		//暴露一个通过id进行get的接口
		group1.GET("/:id", routers.GetByID)
	}
	group2 := r.Group("/set")
	{
		group2.POST("/", routers.SetSwitch)
		group2.POST("/switch", routers.SetSwitch)
		group2.POST("/ocvalue", routers.SetocValue) //0 -> close value 1 -> open value
		//暴露一个通过id进行set的接口
		group2.GET("/:id", routers.SetByID)
	}
	group3 := r.Group("/user")
	{
		group3.OPTIONS("/", routers.CheckCode)
		group3.GET("/", routers.JWTAuth(), routers.CheckCode)
		group3.POST("/login", routers.Login)
		group3.POST("/registry", routers.Registry)
		group3.POST("/apply", routers.GetCode)
	}
	r.Run(":9090")
}
