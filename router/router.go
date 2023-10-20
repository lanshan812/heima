package routers

import (
	"net/http"
	"patent/router/patent"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Load(g *gin.Engine, mw ...gin.HandlerFunc) *gin.Engine {
	// swagger api doc : http://localhost:6060/swagger/index.html http://10.69.70.91:6060/swagger/index.html#/
	g.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length", "Authorization", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	g.Use(gin.Recovery())
	g.Use(mw...)

	//404 handler
	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "Incorrect API Path")
	})

	consoleRg := g.Group("api")
	patent.Load(consoleRg)
	return g
}
