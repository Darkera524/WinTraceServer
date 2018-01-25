package g

import (
	"sync"
	"github.com/toolkits/file"
	"fmt"
	"encoding/json"
	"time"
)

var (
	ConfigFile string
	config *Config
	lock = new(sync.RWMutex)
)

type Config struct {
	Listen 	string	`json:"listen"`
	Trace	[]*Trace	`json:"trace"`
}

type Trace struct {
	Hostname	[]string	`json:"hostname"`
	Name	string	`json:"name"`
	Guid	string	`json:"guid"`
	Tags	[]string	`json:"tags"`
}

func GetConfig() *Config {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		fmt.Println(err.Error())
	}
	var c Config
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		fmt.Println(err.Error())
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c
}

func CronParse(cfg string) {
	for {

		ParseConfig(cfg)
		fmt.Println(1)
		time.Sleep(time.Duration(60) * time.Second)
	}
}