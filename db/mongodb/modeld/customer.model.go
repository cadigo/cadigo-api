package modeld

import (
	"cadigo-api/app/modela"
	"time"
)

type Customer struct {
	BaseBSONModel `bson:",inline"`
	UserID        string   `bson:"userID" copier:"UserID"`
	Name          string   `bson:"name" copier:"Name"`
	Images        []string `bson:"images" copier:"Images"`
}

func (this Customer) Init() Customer {
	this.BaseBSONModel.CreatedAt = time.Now()
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this Customer) Update() Customer {
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this Customer) ToCustomer() modela.Customer {

	return modela.Customer{
		ID:     this.RawID.Hex(),
		UserID: &this.UserID,
		Name:   &this.Name,
		Images: []*string{},
	}
}
