package Config

import (
	"fmt"
	"github.com/spf13/viper"
)

var ConfigViper = viper.New()

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	//viper.AddConfigPath("/app/Config")
	viper.AddConfigPath("./Config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Error reading config file:", err)
		return
	}
	ConfigViper = viper.GetViper()
}
