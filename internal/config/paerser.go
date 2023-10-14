package config

import (
	"os"

	"github.com/traefik/paerser/file"
)

const defaultConfigLoc = "./config.yaml"

type Paerser struct {
	cfg        *ConfigModel
	configFile string
}

func NewPaerser(configFile string) *Paerser {
	return &Paerser{
		configFile: configFile,
	}
}

func (p *Paerser) Instance() *ConfigModel {
	return p.cfg
}

func (p *Paerser) Load() (err error) {
	cfg := DefaultConfig

	cfgFile := defaultConfigLoc
	if p.configFile != "" {
		cfgFile = p.configFile
	}
	if err = file.Decode(cfgFile, &cfg); err != nil && !os.IsNotExist(err) {
		return
	}

	p.cfg = &cfg

	return
}
