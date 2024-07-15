package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Database struct {
		Port     uint16 `yaml:"port"`
		Name     string `yaml:"name"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
	} `yaml:"database"`

	Server struct {
		Addr string `yaml:"addr"`
		Port uint16 `yaml:"port"`
	} `yaml:"server"`
}

func Load(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var cfg = Config{}

	err = yaml.Unmarshal(buf, &cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}

// 設定からDBの接続情報を取得
func (cfg *Config) GetDBDNS() string {
	return fmt.Sprintf(
		"user=%s password=%s port=%d sslmode=disable host=%s",
		cfg.Database.Username,
		cfg.Database.Password,
		cfg.Database.Port,
		cfg.Database.Host,
	)
}

func (cfg *Config) GetSocketAddr() string {
	return fmt.Sprintf("%s:%d",
		cfg.Server.Addr,
		cfg.Server.Port,
	)
}
