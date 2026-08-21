package openrazertray

import (
	"context"
	"log"
	"time"

	"github.com/Nok1on1/RazerBatteryTray/openrazer"

	"fyne.io/systray"
)

const updateInterval = 5 * time.Second

type TrayManager struct {
	razerClient *openrazer.Client
	device      *openrazer.Device
	defaultIcon []byte

	routineCtx  context.Context
	cancelRoute context.CancelFunc
}

func NewTrayManager(razerClient *openrazer.Client, defaultIcon []byte) *TrayManager {
	return &TrayManager{razerClient: razerClient, device: &openrazer.Device{}, defaultIcon: defaultIcon}
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
