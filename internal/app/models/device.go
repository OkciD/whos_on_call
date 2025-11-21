package models

import "github.com/google/uuid"

type DeviceType string

const (
	DeviceTypePC     DeviceType = "pc"
	DeviceTypeLaptop DeviceType = "laptop"
	DeviceTypeMobile DeviceType = "mobile"
)

type Device struct {
	UID  uuid.UUID
	Name string
	Type DeviceType
	User *User
}
