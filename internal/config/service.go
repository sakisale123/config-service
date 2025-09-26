package config

import (
	"fmt"
	"sync"
)

var (
	configs      = make(map[string]Configuration)
	configGroups = make(map[string]ConfigurationGroup)
	mu           sync.Mutex
)

type ConfigService struct{}

func NewConfigService() *ConfigService {
	return &ConfigService{}
}

func (s *ConfigService) CreateConfiguration(config Configuration) (Configuration, error) {
	mu.Lock()
	defer mu.Unlock()

	key := fmt.Sprintf("%s:%s", config.ID, config.Version)
	if _, exists := configs[key]; exists {
		return Configuration{}, fmt.Errorf("konfiguracija sa ID-jem %s i verzijom %s već postoji", config.ID, config.Version)
	}

	configs[key] = config
	return config, nil
}

func (s *ConfigService) GetConfiguration(id, version string) (Configuration, error) {
	mu.Lock()
	defer mu.Unlock()

	key := fmt.Sprintf("%s:%s", id, version)
	config, exists := configs[key]
	if !exists {
		return Configuration{}, fmt.Errorf("konfiguracija sa ID-jem %s i verzijom %s nije pronađena", id, version)
	}
	return config, nil
}

func (s *ConfigService) DeleteConfiguration(id, version string) error {
	mu.Lock()
	defer mu.Unlock()

	key := fmt.Sprintf("%s:%s", id, version)
	if _, exists := configs[key]; !exists {
		return fmt.Errorf("konfiguracija sa ID-jem %s i verzijom %s nije pronađena", id, version)
	}

	delete(configs, key)
	return nil
}

func (s *ConfigService) CreateConfigurationGroup(group ConfigurationGroup) (ConfigurationGroup, error) {
	mu.Lock()
	defer mu.Unlock()

	key := fmt.Sprintf("%s:%s", group.ID, group.Version)
	if _, exists := configGroups[key]; exists {
		return ConfigurationGroup{}, fmt.Errorf("grupa sa ID-jem %s i verzijom %s već postoji", group.ID, group.Version)
	}

	configGroups[key] = group
	return group, nil
}

func (s *ConfigService) GetConfigurationGroup(id, version string) (ConfigurationGroup, error) {
	mu.Lock()
	defer mu.Unlock()

	key := fmt.Sprintf("%s:%s", id, version)
	group, exists := configGroups[key]
	if !exists {
		return ConfigurationGroup{}, fmt.Errorf("grupa sa ID-jem %s i verzijom %s nije pronađena", id, version)
	}
	return group, nil
}

func (s *ConfigService) DeleteConfigurationGroup(id, version string) error {
	mu.Lock()
	defer mu.Unlock()

	key := fmt.Sprintf("%s:%s", id, version)
	if _, exists := configGroups[key]; !exists {
		return fmt.Errorf("grupa sa ID-jem %s i verzijom %s nije pronađena", id, version)
	}

	delete(configGroups, key)
	return nil
}
