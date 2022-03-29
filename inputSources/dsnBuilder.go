package inputsources

import (
	"fmt"
	"log"

	"net/url"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func DSNBuilder(dbType string) string {
	//- declare data
	if dbType == "" {
		panic("dbType is empty")
	}
	driverName := fmt.Sprintf("database.%s", dbType)
	dbHost := viper.GetString(driverName + ".host")
	dbPort := viper.GetString(driverName + ".port")
	dbUser := viper.GetString(driverName + ".user")
	dbPass := viper.GetString(driverName + ".pass")
	dbName := viper.GetString(driverName + ".name")
	//- create connection string
	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Ho_Chi_Minh")
	return fmt.Sprintf("%s?%s", connection, val.Encode())
}
