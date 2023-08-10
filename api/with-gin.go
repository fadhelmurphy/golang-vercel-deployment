package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

// Bikin endpoint
func GRoute(r *gin.RouterGroup){
	r.GET("/with-gin/:name", func(ctx *gin.Context) {
		
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
	
	r.GET("/with-gin", func(ctx *gin.Context) {
		ctx.JSON(400, gin.H{
			"data": gin.H{
				"id": ctx.Query("id"),
			},
		})
	})
}

func RunRoute()  *gin.Engine {
	router := gin.Default()
	// api := router.Group("/api")
	GRoute(&gin.Default().RouterGroup)
	return router
}

func Handler(w http.ResponseWriter, r *http.Request){
	// Start REST Server on main thread
	router := RunRoute()
	router.ServeHTTP(w, r)
}