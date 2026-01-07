package models

import "strconv"

type CallState int8

const (
	CallStateInactive CallState = iota
	CallStateActive
)

var humanReadableCallState = map[CallState]string{
	CallStateInactive: "inactive",
	CallStateActive:   "active",
}

func (t CallState) String() string {
	if str, ok := humanReadableCallState[t]; ok {
		return str
	} else {
		return strconv.FormatInt(int64(t), 10)
	}
}

type DeviceStatus struct {
	Device
	Features []DeviceFeature
}

type UserStatus struct {
	User    *User
	State   CallState
	Devices []DeviceStatus
}

type CallStatus []UserStatus
