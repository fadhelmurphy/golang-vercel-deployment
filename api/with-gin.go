package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

func Handler(w http.ResponseWriter, r *http.Request){
	// Start REST Server on main thread
	router := gin.New()
	router.GET("/:name", func(ctx *gin.Context) {
		
		name := ctx.Param("name")
		// fmt.Println("Hello "+name+ "!")
		if name == "" {
			ctx.JSON(400, gin.H{
				"message": "name not found",
			})
		} else {
			ctx.JSON(200, gin.H{
				"data": fmt.Sprintf("Hello %s!", name),
			})
		}
	})
	
	// router.GET("/", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"data": gin.H{
	// 			"id": ctx.Query("id"),
	// 		},
	// 	})
	// })
	router.ServeHTTP(w, r)
}