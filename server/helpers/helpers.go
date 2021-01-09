package helpers

import (
	"github.com/spf13/viper"
	"vitess.io/vitess/go/vt/log"
)

func ViperEnvVariable(key string) string {
	viper.SetConfigFile("../.env")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("[-] Error while reading enviroment file: %v", err)
	}

	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("[-] Invalid type assertion")
	}

	return value
}
