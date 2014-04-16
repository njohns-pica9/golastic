package main

import (
	"encoding/json"
	"os"
)

type Config struct {
	Test string `json:"test"`
}

func NewConfig(filename string) Config {
	c := Config{}
	c.loadConfig(filename)

	return c
}

func (c *Config) loadConfig(filename string) error {

	configFile, err := os.Open(filename)

	if err != nil {
		return err
	}

	defer configFile.Close()

	err = json.NewDecoder(configFile).Decode(c)

	if err != nil {
		return err
	}

	return nil
}
