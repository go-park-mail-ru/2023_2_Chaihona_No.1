package devices

type DeviceID struct {
	DeviceId string `json:"device_id" db:"device_id"`
}

type DevicesRepository interface {
	AddNewDevice(userID int, deviceID string) (int, error)
	GetDevicesID(CreatorId int) ([]DeviceID, error)
	GetDeviceID(userId int) (string, error)
}