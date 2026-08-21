package openrazertray

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Nok1on1/RazerBatteryTray/utils"

	"fyne.io/systray"
)

func (t *TrayManager) deviceMenu() {
	ctx, cancel := context.WithCancel(context.Background())
	t.routineCtx = ctx
	t.cancelRoute = cancel

	systray.ResetMenu()

	deviceInfoMenu := t.addDeviceInfoMenuItem()
	_ = t.backToDevicesMenuItem()
	go t.batteryLevelChangeRoutine(ctx, deviceInfoMenu)

	t.addExitTrayMenu()
}

func (t *TrayManager) addDeviceInfoMenuItem() *systray.MenuItem {
	return systray.AddMenuItem(fmt.Sprintf("%s: %d%%", t.device.DeviceName, t.device.LastBatteryLevel), "Show device information")
}

func (t *TrayManager) backToDevicesMenuItem() *systray.MenuItem {
	item := systray.AddMenuItem("Back To Devices", "Go back to the devices menu")
	go t.backToDevicesMenuItemHandler(item)
	return item
}

func (t *TrayManager) backToDevicesMenuItemHandler(deviceMenuItem *systray.MenuItem) {
	for range deviceMenuItem.ClickedCh {
		t.cancelRoute()
		t.listDevicesMenu()
	}
}

func (t *TrayManager) batteryLevelChangeRoutine(ctx context.Context, deviceInfoMenu *systray.MenuItem) {
	for {
		log.Println("batteryLevelChangeRoutine: checking battery level")

		batteryLevel, err := t.razerClient.GetBattery()
		if err != nil {
			log.Println("batteryLevelChangeRoutine: error getting battery level:", err)
			t.listDevicesMenu()
			break
		}
		isCharging, err := t.razerClient.IsCharging()
		if err != nil {
			log.Println("batteryLevelChangeRoutine: error getting charging state:", err)
			t.listDevicesMenu()
			break
		}

		if batteryLevel != t.device.LastBatteryLevel || isCharging != t.device.LastchargingState {
			if t.device.LastBatteryLevel-batteryLevel > 5 {
				log.Println("batteryLevelChangeRoutine: battery level changed by more than 5%")
				t.listDevicesMenu()
				break
			}

			icon := utils.GenerateIcon(batteryLevel, isCharging)
			systray.SetIcon(icon)

			deviceInfoMenu.SetTitle(fmt.Sprintf("%s: %d%%", t.device.DeviceName, batteryLevel))
			t.device.LastBatteryLevel = batteryLevel
			t.device.LastchargingState = isCharging

			log.Printf("batteryLevelChangeRoutine: battery level changed to: %d\n", batteryLevel)
		}
		select {
		case <-ctx.Done():
			return
		case <-time.After(updateInterval):
		}
	}
}
