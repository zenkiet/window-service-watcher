package domain

type Config struct {
	Services []ServiceConfig `yaml:"services"`
}

type ServiceConfig struct {
	ID          string `yaml:"id" json:"id"`
	Name        string `yaml:"name" json:"name"`
	Description string `yaml:"description" json:"description"`
	ServiceName string `yaml:"service_name" json:"service_name"`
	LogPath     string `yaml:"log_path" json:"log_path"`
}
