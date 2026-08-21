package openrazertray

import (
	"fmt"
	"log"

	"fyne.io/systray"
)

func (t *TrayManager) listDevicesMenu() {
	systray.ResetMenu()
	devices, err := t.razerClient.GetDeviceDisplayNames()
	if err != nil {
		log.Println("listDevices: error getting devices:", err)
		return
	}

	for _, device := range devices {
		item := systray.AddMenuItem(device.DeviceName, "View battery level for "+device.DeviceName)
		go t.deviceClickHandler(item, device.DeviceName)
	}

	t.addExitTrayMenu()
}

func (t *TrayManager) deviceClickHandler(item *systray.MenuItem, device string) {
	for range item.ClickedCh {
		if err := t.razerClient.SelectDevice(device); err != nil {
			log.Println("deviceClickHandler: error selecting device:", err)
			continue
		}
		systray.SetTitle(fmt.Sprintf("%s's Battery Tray", device))
		t.deviceMenu()
	}
}
