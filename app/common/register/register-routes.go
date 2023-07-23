package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/config"
	routes "github.com/vikashkumar2020/quigo-backend/app/routes"
    middleware "github.com/vikashkumar2020/quigo-backend/app/common/middlewares"
)

// health ckeck api
func healhCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Health check",
	})
}

// CORS middleware
func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }
        c.Next()
    }
}

func Routes(router *gin.Engine, serverConfig *config.ServerConfig) {
	router.Use(CORSMiddleware())
	webV1AuthRouterGroup := router.Group("/" + serverConfig.ServerApiPrefixV1)
    webV1UserRouterGroup := router.Group("/" + serverConfig.ServerApiPrefixV1+ "/profile")
    middleware.RegisterUserMiddleware(webV1UserRouterGroup)
	RegiterWebAuthRoutes(webV1AuthRouterGroup)
    RegiterWebUserRoutes(webV1UserRouterGroup)
	router.GET("/health", healhCheck)

}

// rest api routes 
func RegiterWebAuthRoutes(router *gin.RouterGroup) {
    routes.RegisterAuthRoutes(router)
}

func RegiterWebUserRoutes(router *gin.RouterGroup) {
    routes.RegisterUserRoutes(router)
}