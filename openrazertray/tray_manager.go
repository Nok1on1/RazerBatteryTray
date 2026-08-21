package openrazertray

import (
	"github.com/Nok1on1/RazerBatteryTray/openrazer"
	"log"
	"time"

	"fyne.io/systray"
)

const batteryLevelUpdateInterval = 5 * time.Second

type TrayManager struct {
	razerClient      *openrazer.Client
	lastBatteryLevel int8
}

func NewTrayManager(razerClient *openrazer.Client) *TrayManager {
	return &TrayManager{razerClient: razerClient, lastBatteryLevel: -1}
}

func (t *TrayManager) Start() {
	systray.Run(t.onReady, t.onExit)
}

func (t *TrayManager) onReady() {
	log.Println("onReady: starting tray manager")
	systray.SetTitle("Razer Battery Tray")
	go t.listDevicesMenu() // starting menu
	log.Println("onReady: tray manager started")
}

func (t *TrayManager) onExit() {

}
