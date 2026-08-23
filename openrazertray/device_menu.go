package openrazertray

import (
	"fmt"
	"log"
	"time"

	"github.com/Nok1on1/RazerBatteryTray/openrazer"
	"github.com/Nok1on1/RazerBatteryTray/utils"

	"fyne.io/systray"
)

func (t *TrayManager) deviceMenu(device openrazer.Device) {
	systray.ResetMenu()
	menuCtx := CreateContext()

	device.LastBatteryLevel = -1

	deviceInfoMenu := t.deviceInfoMenuItem(device)
	t.backToDevicesMenuItem(&menuCtx)
	t.ExitTrayMenuItem(&menuCtx)

	go t.watchBatteryRoutine(&menuCtx, &device, deviceInfoMenu)

	go func() {
		<-menuCtx.ctx.Done()
		t.listDevicesMenu()
	}()
}

func (t *TrayManager) deviceInfoMenuItem(device openrazer.Device) *systray.MenuItem {
	return systray.AddMenuItem(fmt.Sprintf("%s: %d%%", device.DeviceName, device.LastBatteryLevel), "Show device information")
}

func (t *TrayManager) backToDevicesMenuItem(menuCtx *MenuContext) (menuItem *systray.MenuItem) {
	menuItem = systray.AddMenuItem("Back To Devices", "Go back to the devices menu")
	go func() {
		select {
		case <-menuCtx.ctx.Done():
			return
		case <-menuItem.ClickedCh:
			menuCtx.cancel()
			config := utils.GetConfig()
			if config.AutoConnect {
				go config.SetOnCooldown()
			}
		}
	}()
	return
}

func (t *TrayManager) watchBatteryRoutine(menuCtx *MenuContext, device *openrazer.Device, deviceInfoMenu *systray.MenuItem) {
	config := utils.GetConfig()
	for {
		log.Println("BatteryLevelChangeRoutine: Checking battery level")

		batteryLevel, err := t.razerClient.GetBattery()
		if err != nil {
			menuCtx.cancel()
			return
		}
		isCharging, err := t.razerClient.IsCharging()
		if err != nil {
			menuCtx.cancel()
			return
		}

		if batteryLevel != device.LastBatteryLevel || isCharging != device.LastchargingState {
			if device.LastBatteryLevel-batteryLevel > 5 {
				log.Println("BatteryLevelChangeRoutine: Battery level changed by more than 5%%")
				menuCtx.cancel()
				return
			}

			icon := utils.GenerateIcon(batteryLevel, isCharging)
			systray.SetIcon(icon)

			deviceInfoMenu.SetTitle(fmt.Sprintf("%s: %d%%", device.DeviceName, batteryLevel))
			device.LastBatteryLevel = batteryLevel
			device.LastchargingState = isCharging

			log.Printf("BatteryLevelChangeRoutine: Battery level changed to: %d\n", batteryLevel)
		}
		select {
		case <-menuCtx.ctx.Done():
			return
		case <-time.After(config.UpdateInterval):
		}
	}
}
