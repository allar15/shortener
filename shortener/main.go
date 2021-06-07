package main

import (
	"log"
	"shortener/handlers"
	"shortener/logic"
	"shortener/repository"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/ssdb/gossdb/ssdb"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panicf("Unable to read config file: %s", err)
	}

	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}


func main() {
	e := echo.New()
	db, err := ssdb.Connect(viper.GetString("database.host"), viper.GetInt("database.port"))
	defer func() {
		err = db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	if err != nil{
		log.Fatal("Unable to connect: ", err)
	}
	shortenerRepo := repository.NewShortenerRepo(db)
	shortenerLogic := logic.NewClientLogic(*shortenerRepo)
	handlers.NewShortenerHandler(e, *shortenerLogic)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
