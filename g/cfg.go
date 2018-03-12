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
	Wmi		[]*Wmi `json:"wmi"`
}

type Trace struct {
	Hostname	[]string	`json:"hostname"`
	Name	string	`json:"name"`
	Guid	string	`json:"guid"`
}

type Wmi struct {
	Hostname []string `json:"hostname"`
	WmiList []string `json:"wmilist"`
	Database []string `json:"database"`
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
		time.Sleep(time.Duration(60) * time.Second)
	}
}