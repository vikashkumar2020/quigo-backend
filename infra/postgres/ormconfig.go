package postgres

import (
	"fmt"
	config "github.com/vikashkumar2020/quigo-backend/config"
)

// GetOrmConfig Get DB string
func GetOrmConfig(dbConfig *config.DBConfig) string {
	configString := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbConfig.Host, dbConfig.Username,
		dbConfig.Password, dbConfig.Dbname,
		dbConfig.Port,
	)
	return configString
}