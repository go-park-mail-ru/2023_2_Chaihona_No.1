package model

type Subscription struct {
	Id                    uint   `json:"id" db:"id"`
	Subscriber_id         uint   `json:"-" db:"subscriber_id"`
	Creator_id            uint   `json:"-" db:"creator_id"`
	Subscription_level_id uint   `json:"-" db:"subscribtion_level_id"`
	CreationDate          string `json:"-" db:"creation_date"`
}
