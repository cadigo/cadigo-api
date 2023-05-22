package repositories

import "go.mongodb.org/mongo-driver/bson"

func ParseUpdate(v interface{}) (bson.M, error) {
	var update bson.M

	pByte, err := bson.Marshal(v)
	if err != nil {
		return nil, err
	}

	err = bson.Unmarshal(pByte, &update)
	if err != nil {
		return nil, err
	}
	return update, nil
}
