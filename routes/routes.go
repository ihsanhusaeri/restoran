package routes

import (
	"net/http"
	pesanan "restoran/controllers/pesanan"

	"github.com/gin-gonic/gin"
)

//StartService function
func StartService() {
	router := gin.Default()
	api := router.Group("/api/v1")
	{
		api.GET("/pesanan/:id", pesanan.GetPesanan)
		api.POST("/pesanan", pesanan.CreatePesanan)
		//  api.GET("/users/:id", user.GetUser)
		//  api.PUT("/users/:id", user.UpdateUser)
		//  api.DELETE("/users/:id", user.DeleteUser)
	}
	router.NoRoute(func(c *gin.Context) {
		c.AbortWithStatus(http.StatusNotFound)
	})
	router.Run(":8001")
}
