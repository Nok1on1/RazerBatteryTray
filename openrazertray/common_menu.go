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

func (t *TrayManager) ExitTrayMenuItem(menuCtx *MenuContext) *systray.MenuItem {
	exitTray := systray.AddMenuItem("Exit Razer Battery Tray", "Exit the tray")

	go func() {
		select {
		case <-menuCtx.ctx.Done():
			return
		case <-exitTray.ClickedCh:
			systray.Quit()
		}
	}()

	return exitTray
}
