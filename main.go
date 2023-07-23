package main

import (
	"github.com/gin-gonic/gin"
	"github.com/vikashkumar2020/quigo-backend/app/common/register"
	"github.com/vikashkumar2020/quigo-backend/config"
	pgdatabase "github.com/vikashkumar2020/quigo-backend/infra/postgres/database"
)


func main() {

	// import all config
	// Initialize the config
	config.LoadEnv()
	config.LogInfo("env loaded")

	// Initialize the server
	serverConfig := config.NewServerConfig()
	config.LogInfo("server config loaded")

	// Initialize the database
	dbConfig := config.NewDBConfig()
	config.LogInfo("db config loaded")

	config.LogInfo(serverConfig.Port)

	// initialize database
	database := pgdatabase.GetDBInstance();
	database.NewDBConnection(dbConfig);
	config.LogInfo("db connection established")

	router := gin.Default()
	register.Routes(router, serverConfig)
	router.Run(":" + serverConfig.Port)
	if err := router.Run(":" + serverConfig.Port); err != nil {
		config.LogFatal(err)
	}
	config.LogInfo("server started")

	
}