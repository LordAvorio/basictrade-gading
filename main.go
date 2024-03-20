package main

import (
	"basictrade-gading/database"
	"basictrade-gading/routes"
	"basictrade-gading/utils"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	utils.ReadConfigEnvironment()
	utils.InitZeroLog()
	utils.InitGoValidation()
}

func main() {

	db, errDB := database.StartDB()
	if errDB != nil {
		log.Panic().Msg(errDB.Error())
	}

	if runMigration := viper.GetBool("AUTO_MIGRATE"); runMigration {
		errMigration := database.RunMigration(db)
		if errMigration != nil {
			log.Panic().Msg(errMigration.Error())
		}
	}

	portApp := viper.GetString("APP_PORT")
	app := routes.RouteSession(db)
	app.Run(portApp)

	
}
