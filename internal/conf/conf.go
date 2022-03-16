package conf

import (
	"gitee.com/bethink1501/24on-library/db"
	"gitee.com/bethink1501/24on-library/lib/jsontime"
	"gitee.com/bethink1501/24on-library/zconf"
	"gitee.com/bethink1501/24on-library/zredis"
	"github.com/spf13/viper"
)

type Config struct {
	Server *Server `json:"server"`
	Data   *Data   `json:"data"`
}

type Server struct {
	Http *Http `json:"http"`
	Grpc *Grpc `json:"grpc"`
}

type Grpc struct {
	Network string            `json:"network"`
	Addr    string            `json:"addr"`
	Timeout jsontime.Duration `json:"timeout"`
}

type Http struct {
	Network string            `json:"network"`
	Addr    string            `json:"addr"`
	Timeout jsontime.Duration `json:"timeout"`
}

type Data struct {
	Mysql *db.Config     `json:"mysql"`
	Redis *zredis.Config `json:"redis"`
}

func InitConfig() *Config {
	myConf := new(Config)

	v := viper.New()
	v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(".")
	v.AddConfigPath("configs")
	v.AddConfigPath("../../configs")

	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := v.UnmarshalExact(myConf, zconf.ViperOptions()); err != nil {
		panic(err)
	}

	return myConf
}
