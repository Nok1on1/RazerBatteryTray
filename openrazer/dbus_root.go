package openrazer

import (
	"github.com/Nok1on1/RazerBatteryTray/utils"

	"github.com/godbus/dbus/v5"
)

type DbusRootMethods struct {
	Path dbus.ObjectPath
}

func NewRoot() DbusRootMethods {
	return DbusRootMethods{Path: "/org/razer"}
}

func (m DbusRootMethods) GetDevices() utils.RazerMethod {
	return utils.RazerMethod{Path: m.Path, MethodName: "getDevices"}
}
