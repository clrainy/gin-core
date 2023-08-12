package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func M1(c *gin.Context) {
	fmt.Println("m1 in...")
	// 计时
	start := time.Now()

	c.Next() // 调用后续的处理函数
	//c.Abort() // 阻止调用后续的处理函数
	cost := time.Since(start)
	fmt.Printf("cost:%v\n", cost)
	fmt.Println("m1 out...")
}
