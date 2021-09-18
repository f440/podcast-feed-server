package config

import "github.com/BurntSushi/toml"

type Config struct {
	RSS    RSS    `toml:"rss"`
	Server Server `toml:"server"`
}

type RSS struct {
	Title       string
	Description string
	URL         string
}

type Server struct {
	Listen   string
	FeedPath string `toml:"feed_path"`
	FileRoot string `toml:"file_root"`
	PermitUA string `toml:"permit_ua"`
	Exclude  []string
}

func (c *Config) Load(file string) error {
	if _, err := toml.DecodeFile(file, c); err != nil {
		return err
	}
	return nil
}
