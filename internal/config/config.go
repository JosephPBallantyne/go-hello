package config

import (
	"path/filepath"
	"runtime"

	"github.com/josephpballantyne/hello/internal/hello"
	"github.com/spf13/viper"
)

var (
	_, b, _, _ = runtime.Caller(0)
	Root       = filepath.Join(filepath.Dir(b), "../..")
	constants  *Constants
)

type Constants struct {
	PORT string
	ENV  string
}

func NewConfig() *Constants {
	return &Constants{}
}

func InitViper() (*Constants, error) {
	const op = "config.InitViper"
	v := viper.New()
	v.AddConfigPath(Root)
	v.SetConfigFile(".env")
	if err := v.MergeInConfig(); err != nil {
		if _, ok := err.(viper.ConfigParseError); ok {
			return constants, &hello.Error{Op: op, Err: err}
		}
	}
	v.AutomaticEnv()
	constants = &Constants{}
	err := v.Unmarshal(&constants)
	if err != nil {
		return constants, &hello.Error{Op: op, Err: err}
	}
	return constants, nil
}
