package config

import (
	_ "embed"
	"fmt"
	"github.com/BurntSushi/toml"
)

//go:embed app.toml
var confFile string

func Initialize() *M {
	conf := &M{}
	_, err := toml.Decode(confFile, &conf)
	if err != nil {
		panic(fmt.Sprintf("加载配置文件[app.toml]失败:%+v", err))
	}
	return conf
}
