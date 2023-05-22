package modeld

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseBSONModel struct {
	RawID     primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" copier:"Id"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at" copier:"CreatedAt"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" bson:"updated_at" copier:"UpdatedAt"`
}
