package repository

import (
	"os"
	"sync"
	"window-service-watcher/internal/domain"

	"go.yaml.in/yaml/v4"
)

type YamlConfigRepository struct {
	path string
	mu   sync.RWMutex
}

func NewYamlConfigRepository(path string) *YamlConfigRepository {
	return &YamlConfigRepository{path: path}
}

func (r *YamlConfigRepository) Load() (*domain.Config, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	data, err := os.ReadFile(r.path)
	if err != nil {
		return nil, err
	}

	var cfg domain.Config
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, err
	}

	return &cfg, nil
}

func (r *YamlConfigRepository) Save(cfg *domain.Config) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}

	return os.WriteFile(r.path, data, 0644)
}
