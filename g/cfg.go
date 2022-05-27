package g

import (
	"encoding/json"
	"github.com/toolkits/file"
	"log"
	"sync"
)

type GlobalConfig struct {
	LogLevel string `json:"log_level"`
}

var (
	config *GlobalConfig
	configLock = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	configLock.Lock()
	defer configLock.Unlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}
	if !file.IsExist(cfg) {
		log.Fatalln("config file: ", cfg, "is not existent")
	}
	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file: ", cfg, "fail, err: ", err)
	}
	if err = json.Unmarshal([]byte(configContent), &config); err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}
	log.Println("read config file:", cfg, "successfully")
}
