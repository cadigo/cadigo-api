package modeld

import (
	"cadigo-api/app/modela"
	"fmt"
	"time"
)

type Caddy struct {
	BaseBSONModel `bson:",inline"`
	Language      string   `bson:"language" copier:"Language"`
	Name          string   `bson:"name" copier:"Name"`
	Location      string   `bson:"location,omitempty" copier:"Location"`
	Avialability  string   `bson:"avialability,omitempty" copier:"Avialability"`
	Skill         []string `bson:"skill,omitempty" copier:"Skill"`
	Start         int      `bson:"start,omitempty" copier:"Start"`
	Description   string   `bson:"description,omitempty" copier:"Description"`
	Time          []string `bson:"time,omitempty" copier:"Time"`
	Tags          []string `bson:"tags,omitempty" copier:"Tags"`
	Images        []string `bson:"images,omitempty" copier:"Images"`
	Cost          float64  `bson:"cost" copier:"Cost"`
}

func (this Caddy) Init() Caddy {
	this.BaseBSONModel.CreatedAt = time.Now()
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this Caddy) Update() Caddy {
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this Caddy) ToCaddy() modela.Caddy {
	var skill []string
	var time []string

	for _, v := range this.Skill {
		skill = append(skill, v)
	}

	for _, v := range this.Time {
		time = append(time, v)
	}

	fmt.Println(this.RawID.Hex())

	return modela.Caddy{
		ID:           this.RawID.Hex(),
		Name:         this.Name,
		Location:     this.Location,
		Avialability: this.Avialability,
		Skill:        skill,
		Start:        this.Start,
		Description:  this.Description,
		Time:         time,
		Cost:         this.Cost,
	}
}
