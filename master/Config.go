package master

import (
	"encoding/json"
	"io/ioutil"
)

type Config struct {
	ApiPort         int `json:"api_port"`
	ApiReadTimeout  int `json:"api_read_timeout"`
	ApiWriteTimeout int `json:"api_write_timeout"`
}

var (
	G_config *Config
)

func InitConfig(filename string) (err error) {
	var (
		content []byte
		config  Config
	)
	// 1. 把配置的文件读出来
	if content, err = ioutil.ReadFile(filename); err != nil {
		return
	}
	// 2. 做JSON反序列化
	if err = json.Unmarshal(content, &config); err != nil {
		return
	}
	// 3. 赋值单例
	G_config = &config
	return
}
