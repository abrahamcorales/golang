package main

import (
	"fmt"
	"sync"
)

type Config struct {
	AppName string
}

var (
	instance *Config
	once     sync.Once
)

func main() {
	config1 := GetConfig()
	config2 := GetConfig()
	fmt.Println("AppName from config1:", config1.AppName)
	fmt.Println("AppName from config2:", config2.AppName)
}

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}
	})
	return instance
}

func (c *Config) SetAppName(name string) {
	c.AppName = name
}
