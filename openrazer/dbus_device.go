package openrazer

import (
	"github.com/Nok1on1/RazerBatteryTray/utils"

	"github.com/godbus/dbus/v5"
)

type DbusDeviceMethods struct {
	Path dbus.ObjectPath
}

func NewDevice(serial string) DbusDeviceMethods {
	return DbusDeviceMethods{
		Path: DevicePath(serial),
	}
}

func DevicePath(serial string) dbus.ObjectPath {
	return dbus.ObjectPath("/org/razer/device/" + serial)
}

func (d DbusDeviceMethods) GetBattery() utils.RazerMethod {
	return utils.RazerMethod{Path: d.Path, MethodName: "razer.device.power.getBattery"}
}

func (d DbusDeviceMethods) IsCharging() utils.RazerMethod {
	return utils.RazerMethod{Path: d.Path, MethodName: "razer.device.power.isCharging"}
}

func (d DbusDeviceMethods) GetDeviceDisplayName() utils.RazerMethod {
	return utils.RazerMethod{Path: d.Path, MethodName: "razer.device.misc.getDeviceName"}
}
