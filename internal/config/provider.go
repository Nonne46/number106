package config

type ConfigProvider interface {
	Load() (err error)
	Instance() *ConfigModel
}
