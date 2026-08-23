package openrazer

import (
	"errors"
	"log"
	"reflect"

	"github.com/Nok1on1/RazerBatteryTray/utils"

	"github.com/godbus/dbus/v5"
)

type Client struct {
	conn       *dbus.Conn
	rootDbus   DbusRootMethods
	DeviceDbus DbusDeviceMethods
}

type Device struct {
	DeviceName        string
	DeviceSerial      string
	LastBatteryLevel  int8
	LastchargingState bool
}

func NewClient() (*Client, error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return nil, err
	}
	return &Client{
		conn:       conn,
		rootDbus:   NewRoot(),
		DeviceDbus: NewDevice(""),
	}, nil
}

func (c *Client) call(method utils.RazerMethod, out any, args ...any) error {
	obj := c.conn.Object("org.razer", method.Path)
	call := obj.Call(method.MethodName, 0, args...)
	if call.Err != nil {
		return call.Err
	}

	if out != nil {
		if reflect.TypeOf(out).Kind() != reflect.Pointer {
			return errors.New("out argument must be a non-nil pointer")
		}
		return call.Store(out)
	}

	return nil
}

func (c *Client) GetDevices() (devices []Device, err error) {
	var deviceSerials []string
	err = c.call(c.rootDbus.GetDevices(), &deviceSerials)
	if err != nil {
		return
	}

	devices = make([]Device, len(deviceSerials))
	for i, v := range deviceSerials {
		devices[i].DeviceSerial = v
	}

	for i := range devices {
		err := c.call(NewDevice(devices[i].DeviceSerial).GetDeviceDisplayName(), &devices[i].DeviceName)
		if err != nil {
			log.Println("error getting device display name:", err)
			continue
		}
	}
	log.Println("devices:", devices)
	return
}

func (c *Client) SetDevice(device Device) {
	c.DeviceDbus = NewDevice(device.DeviceSerial)
}

func (c *Client) GetDeviceDisplayName() (name string, err error) {
	err = c.call(c.DeviceDbus.GetDeviceDisplayName(), &name)
	return
}

func (c *Client) GetBattery() (level int8, err error) {
	err = c.call(c.DeviceDbus.GetBattery(), &level)
	return
}

func (c *Client) IsCharging() (charging bool, err error) {
	err = c.call(c.DeviceDbus.IsCharging(), &charging)
	return
}
