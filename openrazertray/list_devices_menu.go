package openrazertray

import (
	"fmt"
	"log"
	"time"

	"fyne.io/systray"
)

func (t *TrayManager) listDevicesMenu() {
	t.NewMenu()

	systray.SetIcon(t.defaultIcon)

	deviceInfoMenuParent := t.addDeviceInfoParentMenu()
	go t.DeviceInfoMenuRoutine(deviceInfoMenuParent)

	t.addExitTrayMenu()
}

func (t *TrayManager) addDeviceInfoParentMenu() *systray.MenuItem {
	return systray.AddMenuItem("Devices", "List all connected devices")
}

func (t *TrayManager) DeviceInfoMenuRoutine(parent *systray.MenuItem) {
	deviceNames := make(map[string]*systray.MenuItem)

	for {
		devices, err := t.razerClient.GetDevices()
		if err != nil {
			log.Println("listDevices: error getting devices:", err)
			return
		}

		for _, device := range devices {
			if _, ok := deviceNames[device.DeviceName]; ok {
				continue
			}
			item := parent.AddSubMenuItem(device.DeviceName, "View battery level for "+device.DeviceName)
			deviceNames[device.DeviceName] = item
			go t.deviceClickHandler(item, device.DeviceName)
		}
		select {
		case <-t.routineCtx.Done():
			return
		case <-time.After(updateInterval):
		}
	}
}

func (t *TrayManager) deviceClickHandler(item *systray.MenuItem, device string) {
	for range item.ClickedCh {
		t.cancelRoute()

		device, err := t.razerClient.SelectDevice(device)
		if err != nil {
			log.Println("deviceClickHandler: error selecting device:", err)
			continue
		}
		t.device = &device

		systray.SetTitle(fmt.Sprintf("%s's Battery Tray", device.DeviceName))
		t.deviceMenu()
		break
	}
}
