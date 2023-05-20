package modelD

import "go.mongodb.org/mongo-driver/bson/primitive"

type BaseBSONModel struct {
	RawID     primitive.ObjectID `json:"_id,omitempty" bson:"_id"`
	CreatedAt int64              `json:"created_at,omitempty" bson:"created_at"`
	UpdatedAt int64              `json:"updated_at,omitempty" bson:"updated_at"`
}

type Language struct {
	Content LanguageType `bson:"content,omitempty"`
	Value   string       `bson:"value,omitempty"`
}
