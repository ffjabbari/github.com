package main

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func main() {
	viper.AddRemoteProvider("consul", "127.0.0.1:8500", "/config/helloconfig.json")
	viper.SetConfigType("json")
	err := viper.ReadRemoteConfig()

	if err != nil {
		log.Fatal(err)
	}

	msg := viper.GetString("msg")
	universalAnswer := viper.GetInt("the-answer")
	fmt.Printf("\nMsg: %s\nUniversal Answer: %d\n", msg, universalAnswer)
}
