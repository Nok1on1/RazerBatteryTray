package openrazertray

import (
	"context"

	"fyne.io/systray"
)

type MenuContext struct {
	ctx    context.Context
	cancel context.CancelFunc
}

func CreateContext() MenuContext {
	ctx, cancel := context.WithCancel(context.Background())
	return MenuContext{ctx: ctx, cancel: cancel}
}

func (t *TrayManager) ExitTrayMenuItem() (exitTray *systray.MenuItem) {
	exitTray = systray.AddMenuItem("Exit Razer Battery Tray", "Exit the tray")

	go func() {
		for range exitTray.ClickedCh {
			systray.Quit()
		}
	}()

	return
}
