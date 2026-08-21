package main

import (
	_ "embed"
	"log"

	"github.com/Nok1on1/RazerBatteryTray/openrazer"
	"github.com/Nok1on1/RazerBatteryTray/openrazertray"
)

//go:embed assets/snake-64x64.png
var defaultIcon []byte

func main() {
	openRazerClient, err := openrazer.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	trayManager := openrazertray.NewTrayManager(openRazerClient, defaultIcon)
	trayManager.Start()
}
