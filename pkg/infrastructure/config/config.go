package config

type AppConfig struct {
	Logger *Logger `mapstructure:"logger" validate:"required"`
	Server *Server `mapstructure:"server" validate:"required"`
}

func LoadConfig(configPath, secretPath string) (*AppConfig, error) {

	cfg := LoadDefaultConfig()

	return cfg, nil
}

func LoadDefaultConfig() *AppConfig {
	return &AppConfig{}
}
