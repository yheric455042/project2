package project2

import (
	"github.com/gin-gonic/gin"
)


func run(port string) {
	r := gin.Default()
	r.GET("/getlottery/:e_no", func(c *gin.Context) {
		var e_no = c.Param("e_no")
		var code = 200
		var msg = nil
		if judgeEno(e_no) {
			msg := gin.H{"message": e_no + "is good"})
		} else {
			code = 400
			msg := gin.H{"error": e_no + "is invalid"}
		}
		c.JSON(code,msg)

	})
	r.Run(":"+port) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
