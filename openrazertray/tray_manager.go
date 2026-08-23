package openrazertray

import (
	"log"

	"github.com/Nok1on1/RazerBatteryTray/openrazer"

	"fyne.io/systray"
)

type TrayManager struct {
	razerClient *openrazer.Client
	defaultIcon []byte
}

func NewTrayManager(razerClient *openrazer.Client, defaultIcon []byte) *TrayManager {
	return &TrayManager{razerClient: razerClient, defaultIcon: defaultIcon}
}

func (t *TrayManager) Start() {
	systray.Run(t.onReady, t.onExit)
}

func (t *TrayManager) onReady() {
	log.Println("onReady: starting tray manager")
	systray.SetTitle("Razer Battery Tray")
	systray.SetIcon(t.defaultIcon)
	t.listDevicesMenu() // starting menu
	log.Println("onReady: tray manager started")
}

func (t *TrayManager) onExit() {

}
