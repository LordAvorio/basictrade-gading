package main

import (
	"basictrade-gading/database"
	"basictrade-gading/routes"
	"basictrade-gading/utils"
	"os"
	"strconv"
	"github.com/rs/zerolog/log"
)

func init() {
	// utils.ReadConfigEnvironment()
	utils.InitZeroLog()
	utils.InitGoValidation()
}

func main() {

	db, errDB := database.StartDB()
	if errDB != nil {
		log.Panic().Msg(errDB.Error())
	}

	doMigration := os.Getenv("AUTO_MIGRATE")
	if doMigration == "" {
		log.Panic().Msg("Cannot find environtment for database auto migration")
	} else {
		doMigrationBool, err := strconv.ParseBool(doMigration)
		if err != nil {
			log.Panic().Msg(err.Error())
		}
		if doMigrationBool {
			errMigration := database.RunMigration(db)
			if errMigration != nil {
				log.Panic().Msg(errMigration.Error())
			}
		}
	}

	portApp := os.Getenv("APP_PORT")
	app := routes.RouteSession(db)
	app.Run(portApp)

}
