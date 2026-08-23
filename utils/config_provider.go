package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	AutoConnect         bool
	AutoSelectDevice    string
	Log                 bool
	UpdateInterval      time.Duration
	AutoConnectCooldown time.Duration
	OnCooldown          *bool
}

func GetConfig() *Config {
	once.Do(parseArgs)
	return config
}

func (c *Config) SetOnCooldown() {
	log.Println("AutoConnectCooldown: ", c.AutoConnectCooldown)
	*c.OnCooldown = true
	go func() {
		time.Sleep(c.AutoConnectCooldown)
		*c.OnCooldown = false
		log.Println("AutoConnectCooldown: Cooldown ended")
	}()
}

func parseArgs() {
	autoConnectCooldown := 30 * time.Second
	onCooldown := false
	config = &Config{UpdateInterval: 5 * time.Second, AutoConnectCooldown: autoConnectCooldown, OnCooldown: &onCooldown}
	for i := 1; i < len(os.Args); i++ {
		switch os.Args[i] {
		case "--autoConnect":
			config.AutoConnect = true
			if i+1 < len(os.Args) && strings.HasPrefix(os.Args[i+1], "--") {
				config.AutoSelectDevice = os.Args[i+1]
				i++
			}
		case "--log":
			config.Log = true
		case "--updateInterval":
			if i+1 < len(os.Args) {
				sec, err := strconv.ParseInt(os.Args[i+1], 10, 8)
				if err != nil {
					log.Fatal("invalid update interval: ", err)
				}
				config.UpdateInterval = time.Duration(sec) * time.Second
				i++
			}
		case "--autoConnectCooldown":
			if i+1 < len(os.Args) {
				sec, err := strconv.ParseInt(os.Args[i+1], 10, 8)
				if err != nil {
					log.Fatal("invalid update interval: ", err)
				}
				config.AutoConnect = true
				config.AutoConnectCooldown = time.Duration(sec) * time.Second
				i++
			}
		}
	}

}
