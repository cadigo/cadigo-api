package modeld

import (
	"cadigo-api/app/modela"
	"time"
)

type Caddy struct {
	BaseBSONModel `bson:",inline"`
	Language      string   `bson:"language" copier:"Language"`
	Name          string   `bson:"name" copier:"Name"`
	Reference     string   `bson:"reference" copier:"Reference"`
	Location      string   `bson:"location,omitempty" copier:"Location"`
	Avialability  string   `bson:"avialability,omitempty" copier:"Avialability"`
	Skill         []string `bson:"skill,omitempty" copier:"Skill"`
	Star          int      `bson:"star,omitempty" copier:"Star"`
	Description   string   `bson:"description,omitempty" copier:"Description"`
	Time          []string `bson:"time,omitempty" copier:"Time"`
	Tags          []string `bson:"tags,omitempty" copier:"Tags"`
	Images        []string `bson:"images,omitempty" copier:"Images"`
	Cost          float64  `bson:"cost" copier:"Cost"`
	CourseGolfIDs []string `bson:"courseGolfIDs,omitempty" copier:"CourseGolfIDs"`
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

	return modela.Caddy{
		ID:           this.RawID.Hex(),
		Name:         this.Name,
		Location:     this.Location,
		Avialability: this.Avialability,
		Skill:        skill,
		Star:         this.Star,
		Description:  this.Description,
		Time:         time,
		Cost:         this.Cost,
	}
}
