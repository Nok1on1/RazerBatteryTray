package utils

import (
	"github.com/godbus/dbus/v5"
)

type RazerMethod struct {
	Path       dbus.ObjectPath
	MethodName string
}
