package config

import (
	"log"
	"time"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLConfig() mysql.Config {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
	}

	return mysql.Config{
		User:                 getEnvOrDefault("NS_MARIADB_USER", "root"),
		Passwd:               getEnvOrDefault("NS_MARIADB_PASSWORD", "root"),
		Net:                  "tcp",
		Addr:                 getEnvOrDefault("NS_MARIADB_HOSTNAME", "localhost") + ":" + getEnvOrDefault("NS_MARIADB_PORT", "3306"),
		DBName:               getEnvOrDefault("NS_MARIADB_DATABASE", "sodan"),
		Loc:                  jst,
		ParseTime:            true,
		AllowNativePasswords: true,
	}
}
