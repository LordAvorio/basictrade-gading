package main

import (
	"basictrade-gading/database"
	"basictrade-gading/utils"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func init() {
	utils.ReadConfigEnvironment()
	utils.InitZeroLog()
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
	log.Info().Msg("SUCCESS RUNNING BASIC-TRADE SERVICE")
}
