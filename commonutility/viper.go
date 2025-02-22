package commonutility

import (
	"strings"

	"github.com/spf13/viper"
)

func ReadConfigFile(filepath ...string) {
	path := "./config/config.json"
	if len(filepath) != 0 && filepath[0] != "" {
		path = filepath[0]
	}
	folderStructure := strings.Split(path, "/")
	filename := strings.Split(folderStructure[len(folderStructure)-1], ".")[0]
	folderStructure = folderStructure[:len(folderStructure)-1]
	folderPath := strings.Join(folderStructure, "/")
	viper.SetConfigName(filename)
	viper.AddConfigPath(folderPath)
	viper.SetConfigType("json")
	err := viper.MergeInConfig()
	if err != nil {
		panic("viper config read from File FAILED::" + err.Error())
	}
	viper.AutomaticEnv()
}
