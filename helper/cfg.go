package helper

import "github.com/spf13/viper"

func ReadFromConfigOrPanic(path string, name string) *viper.Viper {
	config := viper.New()
	config.AddConfigPath(path)
	config.SetConfigName(name)
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}

	return config
}
