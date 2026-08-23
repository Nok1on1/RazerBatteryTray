package openrazertray

import (
	"log"
	"time"

	"fyne.io/systray"
	"github.com/Nok1on1/RazerBatteryTray/openrazer"
	"github.com/Nok1on1/RazerBatteryTray/utils"
)

func (t *TrayManager) listDevicesMenu() {
	systray.ResetMenu()
	menuCtx := CreateContext()

	systray.SetIcon(t.defaultIcon)
	deviceInfoMenuParent := t.addDeviceInfoParentMenu()
	defer t.ExitTrayMenuItem()

	go t.DeviceInfoMenuRoutine(&menuCtx, deviceInfoMenuParent)
}

func (t *TrayManager) addDeviceInfoParentMenu() *systray.MenuItem {
	return systray.AddMenuItem("Devices", "List all connected devices")
}

func (t *TrayManager) DeviceInfoMenuRoutine(menuCtx *MenuContext, parent *systray.MenuItem) {
	deviceSerials := make(map[string]*systray.MenuItem)
	config := utils.GetConfig()

	for {
		devices, err := t.razerClient.GetDevices()
		if err != nil {
			log.Fatal("DeviceInfoMenuRoutine: Error getting devices:", err)
			return
		}

		if config.AutoConnect && len(devices) == 1 && !*config.OnCooldown {
			menuCtx.cancel()
			t.razerClient.SetDevice(devices[0])
			go t.deviceMenu(devices[0])
			return
		}

		currentDevices := make(map[string]struct{})
		for _, device := range devices {
			currentDevices[device.DeviceSerial] = struct{}{}
			if _, ok := deviceSerials[device.DeviceSerial]; ok {
				continue
			}
			item := parent.AddSubMenuItem(device.DeviceName, "View battery level for "+device.DeviceName)
			deviceSerials[device.DeviceSerial] = item
			go t.deviceClickHandler(menuCtx, item, device)
		}

		for serial := range deviceSerials {
			if _, ok := currentDevices[serial]; !ok {
				log.Printf("DeviceInfoRoutine: Device %s is no longer connected\n", serial)
				deviceSerials[serial].Remove()
				delete(deviceSerials, serial)
			}
		}

		select {
		case <-menuCtx.ctx.Done():
			return
		case <-time.After(config.UpdateInterval):
		}
	}
}

func (t *TrayManager) deviceClickHandler(menuCtx *MenuContext, item *systray.MenuItem, device openrazer.Device) {
	for range item.ClickedCh {
		menuCtx.cancel()
		t.razerClient.SetDevice(device)
		go t.deviceMenu(device)
		break
	}
}
