package openrazertray

import (
	"context"

	"fyne.io/systray"
)

func (t *TrayManager) NewMenu() {
	ctx, cancel := context.WithCancel(context.Background())
	t.routineCtx = ctx
	t.cancelRoute = cancel
	systray.ResetMenu()
}

func (t *TrayManager) addExitTrayMenu() (exitTray *systray.MenuItem) {
	exitTray = systray.AddMenuItem("Exit Razer Battery Tray", "Exit the tray")

	go func() {
		for range exitTray.ClickedCh {
			systray.Quit()
		}
	}()

	return
}
