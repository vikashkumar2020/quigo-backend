package register

import (
	"net/http"

	"github.com/gin-gonic/gin"
	middleware "github.com/vikashkumar2020/quigo-backend/app/common/middlewares"
	routes "github.com/vikashkumar2020/quigo-backend/app/routes"
	"github.com/vikashkumar2020/quigo-backend/config"
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
        // Set the allowed origins
        c.Writer.Header().Set("Access-Control-Allow-Origin", "https://quigo-fr-9lb3.vercel.app")
        
        // Allow credentials (cookies, authentication headers) to be sent with the request
        c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        
        // Specify the allowed headers in the request
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        
        // Specify the allowed HTTP methods
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, PATCH")

        // Handle preflight requests (OPTIONS method)
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)  // Respond with a 204 No Content status
            return
        }

        // Continue to the next middleware or the route handler
        c.Next()
    }
}


func Routes(router *gin.Engine, serverConfig *config.ServerConfig) {
	router.Use(CORSMiddleware())
	webV1AuthRouterGroup := router.Group("/" + serverConfig.ServerApiPrefixV1)
	webV1UserRouterGroup := router.Group("/" + serverConfig.ServerApiPrefixV1 + "/profile")
	webV1DriverRouterGroup := router.Group("/" + serverConfig.ServerApiPrefixV1 + "/driver")
	webV1RiderRouterGroup := router.Group("/" + serverConfig.ServerApiPrefixV1 + "/rider")
	middleware.RegisterUserMiddleware(webV1UserRouterGroup)
	middleware.RegisterDriverMiddleware(webV1DriverRouterGroup)
	middleware.RegisterRiderMiddleware(webV1RiderRouterGroup)
	RegiterWebAuthRoutes(webV1AuthRouterGroup)
	RegiterWebUserRoutes(webV1UserRouterGroup)
	RegiterWebDriverRoutes(webV1DriverRouterGroup)
	RegiterWebRiderRoutes(webV1RiderRouterGroup)
	router.GET("/health", healhCheck)

}

// rest api routes
func RegiterWebAuthRoutes(router *gin.RouterGroup) {
	routes.RegisterAuthRoutes(router)
}

func RegiterWebUserRoutes(router *gin.RouterGroup) {
	routes.RegisterUserRoutes(router)
}

func RegiterWebDriverRoutes(router *gin.RouterGroup) {
	routes.RegisterDriverRoutes(router)
}

func RegiterWebRiderRoutes(router *gin.RouterGroup) {
	routes.RegisterRiderRoutes(router)
}
