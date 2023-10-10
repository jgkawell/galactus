package crud

import (
	"time"
)

const (
	CreatedKey = "history.created"
	UpdatedKey = "history.updated"
	DeletedKey = "history.deleted"
)

type ModelHistory struct {
	Created *ActionHistory `bson:"created"`
	Updated *ActionHistory `bson:"updated"`
	Deleted *ActionHistory `bson:"deleted"`
}

func NewActionHistory(at time.Time) *ActionHistory {
	return &ActionHistory{
		At:     at,
		UserId: "", // deprecated
	}
}

type ActionHistory struct {
	At     time.Time `bson:"at"`
	UserId string    `bson:"user_id"` // deprecated
}
