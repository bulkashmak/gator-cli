package config

import (
	"encoding/json"
	"os"
)

const configFileName string = ".gatorconfig.json"

type Config struct {
	DBURL string `json:"db_url"`
	CurrUserName string `json:"current_user_name"`
}

func (cfg *Config) SetUser(userName string) error {
	cfg.CurrUserName = userName
	return write(*cfg)
}

func Read() (Config, error) {
	data, err := os.ReadFile(configFileName)
	if err != nil {
		return Config{}, err
	}

	var config Config 
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func write(cfg Config) error {
	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	err = os.WriteFile(configFileName, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
