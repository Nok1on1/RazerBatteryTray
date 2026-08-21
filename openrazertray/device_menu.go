package openrazertray

import (
	"github.com/Nok1on1/RazerBatteryTray/utils"
	"log"
	"time"

	"fyne.io/systray"
)

func (t *TrayManager) deviceMenu() {
	log.Println("Menu Changed: DeviceMenu")
	systray.ResetMenu()
	go t.batteryLevelChangeRoutine()
	t.addExitTrayMenu()
	
}

func (t *TrayManager) batteryLevelChangeRoutine() {
	for {
		log.Println("batteryLevelChangeRoutine: checking battery level")
		batteryLevel, _ := t.razerClient.GetBattery()
		log.Printf("batteryLevelChangeRoutine: battery level: %d\n", batteryLevel)
		if batteryLevel != t.lastBatteryLevel {
			systray.SetIcon(utils.GenerateIcon(batteryLevel))
			t.lastBatteryLevel = batteryLevel
			log.Printf("batteryLevelChangeRoutine: battery level changed to: %d\n", batteryLevel)
		}
		time.Sleep(batteryLevelUpdateInterval)
	}
}
