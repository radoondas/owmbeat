// Config is put into a different package to prevent cyclic imports in case
// it is needed in several locations

package config

import "time"

type Config struct {
	Period      time.Duration `config:"period"`
	Regions     []Region      `config:"regions"`
	AppId       string        `config:"appid"`
	MaxApiCalls int           `config:"maxApiCalls"`
}

type Region struct {
	Enabled     bool   `config:"enabled"`
	Name        string `config:"name"`
	Description string `config:"description"`
	LonLeft     int    `config:"lon-left"`
	LatBottom   int    `config:"lat-bottom"`
	LonRight    int    `config:"lon-right"`
	LatTop      int    `config:"lat-top"`
	Zoom        int    `config:"zoom"`
}

var DefaultConfig = Config{
	Period:      1 * time.Hour,
	MaxApiCalls: 60,
}
