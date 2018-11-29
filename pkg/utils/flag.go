package utils

import (
	"github.com/spf13/viper"
	log "github.com/sirupsen/logrus"
)

func Validate(flags []string) {

	for _, value := range flags {
		if viper.GetString(value) == "" {
			log.WithFields(log.Fields{
				"flag": value,
			}).Fatalln("Must be set and non-empty")
		}
	}
	log.Debugln("Flags are valid")
}
