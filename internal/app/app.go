package app

import (
	"log"

	"github.com/bigxxby/digital-travel-test/internal/config"
	"github.com/bigxxby/digital-travel-test/internal/database/connection"
	"github.com/bigxxby/digital-travel-test/internal/database/migration"
	"github.com/bigxxby/digital-travel-test/internal/router"
)

func App() {
	//make log flags to show file name and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	config, err := config.GetCofig()
	if err != nil {
		log.Println(err)
		return
	}

	db, err := connection.GetDB(config)
	if err != nil {
		log.Println(err)
		return
	}

	// migrate db
	err = migration.Migrate(db)
	if err != nil {
		log.Println(err)
		return
	}

	// err = utils.CreateAdmin(db)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	router, err := router.NewRouter(db)
	if err != nil {
		log.Println(err)
		return
	}
	router.Run(":" + config.AppPort)
}
