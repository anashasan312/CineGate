package config

type AppConfig struct {
	Logger *Logger `mapstructure:"logger" validate:"required"`
	Server *Server `mapstructure:"server" validate:"required"`
}

func LoadDefaultConfig() *AppConfig {
	return &AppConfig{}
}

func LoadConfig(configPath, secretPath string) (*AppConfig, error) {

	cfg := LoadDefaultConfig()

	return cfg, nil
}
