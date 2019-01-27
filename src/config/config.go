package config

import (
	"io/ioutil"
	"path/filepath"
)

type StConfig struct {
	BindPorts    string `json: "bind_port"`
	Days_period  int    `json: "days_period"`
	Time_reload  int    `json: "time_reload"`
	Drop_balance int    `json: "Drop_balance"`
	Version      string `json: "version"`
}

func Config_init(path string) *StConfig {
	dir, _ := filepath.Abs("./")
	raw, err := ioutil.ReadFile(dir + path)
	if err != nil {
		panic(err)
	}

	var Conf StConfig
	if err = jsoniter.ConfigFastest.Unmarshal(raw, &Conf); err != nil {
		panic(err)
	}
	return &Conf
}
