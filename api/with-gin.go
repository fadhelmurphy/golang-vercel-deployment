package api

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

var (
	app *gin.Engine
)

func init(){
	// Start REST Server on main thread
	app = gin.New()
	// app.Group("/")
	app.GET("/api/with-gin/:name", func(ctx *gin.Context) {
		
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
	
	app.GET("/api/with-gin", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"data": gin.H{
				"id": ctx.Query("id"),
			},
		})
	})
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}