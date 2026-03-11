package main

import (
	"fmt"

	"github.com/leafayon/musicalorg/internal/http/router"
	"github.com/leafayon/musicalorg/internal/infrastructure/config"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	conf, err := config.Parse()
	if err != nil {
		fmt.Println(err)
	}

	database, err := gorm.Open(sqlite.Open(conf.Database.Path))
	if err != nil {
		fmt.Println(err)
	}

	httpRouter := router.New(fmt.Sprintf("%s:%d", conf.Host, conf.Port))
	httpRouter.SetupRoutes(database)

	if err := httpRouter.Start(); err != nil {
		fmt.Println(err)
	}
}
