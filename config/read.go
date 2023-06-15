package config

import (
	"errors"
	"fmt"
	"github.com/BoyChai/Ant"
	"github.com/spf13/viper"
	"os"
)

func ReadConfig() error {
	viper.SetConfigType("yaml")

	file, err := os.Open("./Honeycomb.yaml")
	if err != nil {
		return err
	}
	err = viper.ReadConfig(file)
	if err != nil {
		return err
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		return err
	}
	validator := Ant.New(Ant.Validator{Parity: Ant.Ant})
	e := validator.Struct(Cfg)
	if e.Is {
		return errors.New(fmt.Sprint(e.Data))
	}
	return nil
}
