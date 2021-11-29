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
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
	// 	AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	//通过nginx反向代理绕过CORS
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
	}
	group2 := r.Group("/set")
	{
		group2.POST("/", routers.SetSwitch)
		group2.POST("/switch", routers.SetSwitch)
		group2.POST("/ocvalue", routers.SetocValue) //0 -> close value 1 -> open value
	}
	r.Run(":9090")
}
