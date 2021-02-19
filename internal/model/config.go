package model

//import "github.com/BurntSushi/toml"
import (
	"fmt"
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Writable struct {
		LogLevel string `toml:"LogLevel"`
		LogFile  string `toml:"LogFile"`
	} `toml:"Writable"`
	Service struct {
		APIHost string `toml:"ApiHost"`
		Timeout int    `toml:"Timeout"`
	} `toml:"Service"`
	Database struct {
		Type         string `toml:"Type"`
		Host         string `toml:"Host"`
		DatabaseName string `toml:"DatabaseName"`
		TableName    string `toml:"TableName"`
	} `toml:"database"`
	//Area map[string]int `toml:"Area"`

	Area struct {
		Area map[string]int `toml:"Area"`
	} `toml:"Area"`
}

func ProvideConfig() *Config {
	configFile := "../config/config.toml"
	f, err := os.Open(configFile)
	if err != nil {
		fmt.Println("打开配置文件错误：", err)

	}
	conf := &Config{}
	if err := toml.NewDecoder(f).Decode(&conf); err != nil {
		panic(err)
	}
	return conf
}
