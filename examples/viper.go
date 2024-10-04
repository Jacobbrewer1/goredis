// This is an example of how you would be able to use the goredis package with viper from a JSON configuration file.

package examples

import (
	"github.com/Jacobbrewer1/goredis"
	"github.com/spf13/viper"
)

func viperExample() {
	v := viper.New()
	v.SetConfigName("config.json")
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := goredis.NewPool(
		goredis.WithDefaultPool(),
		goredis.FromViper(v)...,
	); err != nil {
		panic(err)
	}
}
