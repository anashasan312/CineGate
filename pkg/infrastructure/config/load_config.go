package config

type Logger struct {
	LogLevel string `yaml:"log_level" mapstructure:"log_level"`
}