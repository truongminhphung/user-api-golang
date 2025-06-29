package initialize

import (
	"fmt"
	"os"
	"user-api/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")

	// Check if config file is provided as command line argument
	configFile := "local"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
		// Remove .yaml extension if provided
		if len(configFile) > 5 && configFile[len(configFile)-5:] == ".yaml" {
			configFile = configFile[:len(configFile)-5]
		}
		// Remove config/ prefix if provided
		if len(configFile) > 7 && configFile[:7] == "config/" {
			configFile = configFile[7:]
		}
	}
	viper.SetConfigName(configFile)
	viper.SetConfigType("yaml")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
	fmt.Println("Config file loaded successfully:", viper.ConfigFileUsed())
	fmt.Println("Server port::", viper.GetInt("server.port"))
	fmt.Println("mysql port::", viper.GetInt("mysql.port"))

	// Unmarshal the configuration into the global config struct
	if err := viper.Unmarshal(&global.Config); err != nil {
		fmt.Printf("Unable to decode configuration %v", err)
	}

}
