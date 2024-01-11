package config

import (
	"log"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Config struct {
	DBHOST     string
	DBPORT     int
	DBUSER     string
	DBPASS     string
	DBNAME     string
	SERVERPORT int
	OPENAIKEY  string
}

func InitConfig() *Config {
	var res = new(Config)
	res = loadConfig()

	if res == nil {
		log.Fatal("Config: Cannot start program, failed to load configuration")
		return nil
	}

	return res
}

func loadConfig() *Config {
	var res = new(Config)

	config := viper.New()
	config.SetConfigFile(".env")

	if val := config.GetInt("SERVERPORT"); val != 0 {
		res.SERVERPORT = int(val)
	} else {
		logrus.Error("Config: invalid server port value")
		return nil
	}

	if val := config.GetInt("DBPORT"); val != 0 {
		res.DBPORT = val
	} else {
		logrus.Error("Config: invalid db port value")
		return nil
	}

	if val := config.GetString("DBHOST"); val != "" {
		res.DBHOST = val
	}

	if val := config.GetString("DBUSER"); val != "" {
		res.DBUSER = val
	}

	if val := config.GetString("DBPASS"); val != "" {
		res.DBPASS = val
	}

	if val := config.GetString("DBNAME"); val != "" {
		res.DBNAME = val
	}

	if val := config.GetString("OPENAIKEY"); val != "" {
		res.OPENAIKEY = val
	}

	return res
}
