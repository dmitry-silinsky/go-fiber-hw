package config

type AppConfig struct {
	ListenAddr string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		ListenAddr: getString("LISTEN_ADDR", "localhost:3000"),
	}
}
