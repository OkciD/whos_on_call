package models

type DeviceType int8

const (
	DeviceTypeLaptop DeviceType = iota
	DeviceTypeMobile
	DeviceTypePC
)

type Device struct {
	ID   int
	Name string
	Type DeviceType
	User *User
}
