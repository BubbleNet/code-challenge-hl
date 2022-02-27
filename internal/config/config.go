package config

import (
	"fmt"
	"github.com/spf13/viper"
)

const (
	TestVar = "test.var"
)

func NewConfig() {
	viper.SetDefault(TestVar, "my test")

	viper.BindEnv(TestVar, "TEST_VAR")

	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("$HOME/.sc")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("/go/bin")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error getting config file: %w", err))
	}
}
