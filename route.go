package project2

import (
	"github.com/gin-gonic/gin"
	
)


func run(port string) {
	r := gin.Default()
	r.GET("/getlottery/:e_no", func(c *gin.Context) {
		var e_no = c.Param("e_no")
		var code = 200
		if judgeEno(e_no) 
		c.JSON(code, )
	})
	r.Run(":"+port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
