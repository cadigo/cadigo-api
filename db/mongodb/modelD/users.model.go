package modeld

import (
	"cadigo-api/app/modela"
	"time"
)

type Users struct {
	BaseBSONModel
}

func (this Users) Init() Users {
	this.BaseBSONModel.CreatedAt = time.Now()
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this Users) Update() Users {
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this Users) ToUsers() modela.Users {
	return modela.Users{
		ID: this.RawID.Hex(),
	}
}
