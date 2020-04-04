package configs

import (
	"github.com/lbryio/lbry.go/v2/extras/errors"
	"github.com/tkanos/gonfig"
)

type DbConfig struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Database string `json:"database"`
	Password string `json:"password"`
}
type Configs struct {
	Chainquery DbConfig `json:"chainquery"`
	Reflector  DbConfig `json:"reflector"`
}

var Configuration *Configs

func Init(configPath string) error {
	if Configuration != nil {
		return nil
	}
	c := Configs{}
	err := gonfig.GetConf(configPath, &c)
	if err != nil {
		return errors.Err(err)
	}
	Configuration = &c
	return nil
}
