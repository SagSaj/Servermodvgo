package config

import (
	"io/ioutil"
	"path/filepath"

	jsoniter "github.com/json-iterator/go"
)

type StConfig struct {
	Days_period  int    `json: "days_period"`
	Time_reload  int    `json: "time_reload"`
	Drop_balance int    `json: "Drop_balance"`
	Version      string `json: "version"`
}

var Conf StConfig

func init() {
	Config_init("/modules/config/config.json")
}
func Config_init(path string) *StConfig {
	dir, _ := filepath.Abs("./")
	raw, err := ioutil.ReadFile(dir + path)
	if err != nil {
		panic(err)
	}

	if err = jsoniter.ConfigFastest.Unmarshal(raw, &Conf); err != nil {
		panic(err)
	}
	return &Conf
}
