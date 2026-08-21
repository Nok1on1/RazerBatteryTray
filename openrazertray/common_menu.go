package openrazertray

import "fyne.io/systray"

func (t *TrayManager) addExitTrayMenu() (exitTray *systray.MenuItem) {
	exitTray = systray.AddMenuItem("Exit Razer Battery Tray", "Exit the tray")
	go t.exitTrayHandler(exitTray)
	return
}

func (t *TrayManager) exitTrayHandler(exitTray *systray.MenuItem) {
	for range exitTray.ClickedCh {
		systray.Quit()
	}
}
