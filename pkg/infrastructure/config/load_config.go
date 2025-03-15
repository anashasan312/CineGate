package config

type Logger struct {
	LogLevel string `yaml:"log_level" mapstructure:"log_level"`
}

type Server struct {
	Name        string `yaml:"name" mapstructure:"name"`
	Host        string `yaml:"host" mapstructure:"host"`
	PublicPort  uint   `yaml:"public_port" mapstructure:"public_port"`
	PrivatePort uint   `yaml:"private_port" mapstructure:"private_port"`
}