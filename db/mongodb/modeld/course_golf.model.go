package modeld

import (
	"cadigo-api/app/modela"
	"time"
)

type CourseGolf struct {
	BaseBSONModel
	Language  string   `bson:"language" copier:"Language"`
	Name      string   `bson:"name"`
	Location  string   `bson:"location"`
	Images    []string `bson:"images"`
	Latitude  float64  `bson:"latitude" copier:"Latitude"`
	Longitude float64  `bson:"longitude" copier:"Longitude"`
	IsActive  bool     `bson:"isActive" copier:"IsActive"`
}

func (this CourseGolf) Init() CourseGolf {
	this.BaseBSONModel.CreatedAt = time.Now()
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this CourseGolf) Update() CourseGolf {
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this CourseGolf) ToCourseGolf() modela.CourseGolf {
	return modela.CourseGolf{
		ID:        this.RawID.Hex(),
		Name:      this.Name,
		Images:    this.Images,
		Available: 0,
		Location:  this.Location,
		Latitude:  this.Latitude,
		Longitude: this.Longitude,
		IsActive:  this.IsActive,
	}
}
