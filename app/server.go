package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/kochcoding/golang-rest-template/app/router"
	"github.com/kochcoding/golang-rest-template/vars"
)

var (
	configFile = kingpin.Flag("config", "Path to configfile.").Required().File()
)

func main() {
	vars.LoggerOut = log.New(os.Stdout, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	vars.LoggerErr = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	kingpin.Parse()

	vars.LoggerOut.Printf("[DBG] main(): Starting")
	defer vars.LoggerOut.Print("[DBG] main(): Stopping")

	// ---

	// parse the config file
	viper.SetConfigFile((*configFile).Name())
	err := viper.ReadInConfig()
	if err != nil {
		vars.LoggerErr.Panicf("[FAT] main(): could not read config (%s)", err)
	}
	err = viper.Unmarshal(&vars.Config)
	if err != nil {
		vars.LoggerErr.Panicf("[FAT] main(): could not unmarshal config (%s)", err)
	}

	// initialize DB
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		vars.Config.DB.Host,
		vars.Config.DB.Port,
		vars.Config.DB.User,
		vars.Config.DB.Password,
		vars.Config.DB.DBName,
	)

	vars.DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		vars.LoggerErr.Panicf("[FAT] main(): failed to connect to DB (%s)", err)
	}

	e := echo.New()

	// some basic middlewares
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	mainGroup := e.Group("/api")

	_ = router.NewService(mainGroup)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", vars.Config.Port)))
}
