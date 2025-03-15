package config

type Appconfig struct {
	Logger *Logger `mapstructure:"logger" validate:"required"`
}

func LoadConfig(configPath, secretPath string) (*Appconfig, error) {

	cfg := LoadDefaultConfig()

	return cfg, nil
}

func LoadDefaultConfig() *Appconfig {
	return &Appconfig{}
}
