package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	command "verifyMyTest/m/app/commands"
	"verifyMyTest/m/util"
)

func readConfig() {
	var err error

	viper.SetConfigFile(".env")
	viper.SetConfigType("props")
	err = viper.ReadInConfig()
	if err != nil {
		util.Fatal(util.FatalError, err)
		return
	}

	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		util.Fatal(util.FatalError, "WARNING: file .env not found")
	} else {
		viper.SetConfigFile(".env")
		viper.SetConfigType("props")
		err = viper.MergeInConfig()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	for _, key := range viper.AllKeys() {
		viper.BindEnv(key)
	}
}

func main() {
	readConfig()
	command.Execute()
}
