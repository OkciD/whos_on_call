package models

import "strconv"

type DeviceType int8

const (
	DeviceTypeLaptop DeviceType = iota
	DeviceTypeMobile
	DeviceTypePC
)

var humanReadableDeviceType = map[DeviceType]string{
	DeviceTypeLaptop: "laptop",
	DeviceTypeMobile: "mobile",
	DeviceTypePC:     "pc",
}

func (t DeviceType) String() string {
	if str, ok := humanReadableDeviceType[t]; ok {
		return str
	}
	return strconv.FormatInt(int64(t), 10)
}

type Device struct {
	ID   int
	Name string
	Type DeviceType
	User *User
}

type DeviceListParams struct {
	Name *string
	Type *DeviceType
}
