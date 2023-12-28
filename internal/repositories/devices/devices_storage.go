package devices

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/v2/dbscan"
)

func SelectDeviceIdSQL(userId int) squirrel.SelectBuilder {
	return squirrel.Select("ud.device_id").
		From("public.user_device ud").
		Where(squirrel.Eq{"ud.user_id": userId}).
		PlaceholderFormat(squirrel.Dollar)
}

func SelectSubscribersDevicesByPostId(creatorId int) squirrel.SelectBuilder {
	return squirrel.Select("ud.device_id").
		From("public.user_device ud").
		LeftJoin("public.subscription s ON ud.user_id = s.subscriber_id").
	// Where(fmt.Sprintf("s.subscription_level_id >= p.min_subscription_level_id AND p.id = %d", postID)).
		Where(squirrel.Eq{"s.creator_id": creatorId}).
		PlaceholderFormat(squirrel.Dollar)
}

func InsertUserDeviceSQL(userID uint, deviceID string) squirrel.InsertBuilder {
return squirrel.Insert("public.user_device").
	Columns("user_id", "device_id").
	Values(userID, deviceID).
	Suffix("RETURNING \"id\"").
	PlaceholderFormat(squirrel.Dollar)
}

type DevicesStorage struct {
	db *sql.DB
}

func CreateDevicesStorage(db *sql.DB) DevicesRepository {
	return &DevicesStorage{
		db: db,
	}
}

func (storage *DevicesStorage) GetDevicesID(CreatorId int) ([]DeviceID, error) {
	rows, err := SelectSubscribersDevicesByPostId(CreatorId).RunWith(storage.db).Query()
	if err != nil {
		return []DeviceID{}, err
	}
	var devices []DeviceID
	err = dbscan.ScanAll(&devices, rows)
	if err != nil {
		return []DeviceID{}, err
	}

	return devices, nil
}

func (storage *DevicesStorage) GetDeviceID(userId int) (string, error) {
	var deviceId string; 
	err := SelectDeviceIdSQL(userId).RunWith(storage.db).QueryRow().Scan(&deviceId)
	if err != nil {
		return "", err
	}

	return deviceId, nil
}

func (storage *DevicesStorage) AddNewDevice(userID int, deviceID string) (int, error) {
	var rowId int
	err := InsertUserDeviceSQL(uint(userID), deviceID).RunWith(storage.db).QueryRow().Scan(&rowId)
	if err != nil {
		return 0, err
	}
	return rowId, nil
}