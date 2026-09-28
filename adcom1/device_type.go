package adcom1

// DeviceType represents types of devices (AdCOM 1.0).
// https://github.com/InteractiveAdvertisingBureau/AdCOM/blob/main/AdCOM%20v1.0%20FINAL.md#list_devicetypes
type DeviceType int8

// Types of devices.
const (
	DeviceMobile    DeviceType = 1 // Mobile/Tablet - General
	DevicePC        DeviceType = 2 // Personal Computer
	DeviceTV        DeviceType = 3 // Connected TV
	DevicePhone     DeviceType = 4 // Phone
	DeviceTablet    DeviceType = 5 // Tablet
	DeviceConnected DeviceType = 6 // Connected Device
	DeviceSetTopBox DeviceType = 7 // Set Top Box
	DeviceOOH       DeviceType = 8 // OOH Device
)
