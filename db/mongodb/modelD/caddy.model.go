package modelD

import "cadigo-api/app/modelA"

type Caddy struct {
	BaseBSONModel
	Language     LanguageType `bson:"language" copier:"Language"`
	Name         string       `bson:"name" copier:"Name"`
	Location     string       `bson:"location,omitempty" copier:"Location"`
	Avialability string       `bson:"avialability,omitempty" copier:"Avialability"`
	Skill        []string     `bson:"skill,omitempty" copier:"Skill"`
	Start        int          `bson:"start,omitempty" copier:"Start"`
	Description  string       `bson:"description,omitempty" copier:"Description"`
	Time         []string     `bson:"time,omitempty" copier:"Time"`
	Tags         []string     `bson:"tags,omitempty" copier:"Tags"`
	Cost         float64      `bson:"cost" copier:"Cost"`
}

func (this *Caddy) CaddyDB() modelA.Caddy {
	var skill []string
	var time []string

	for _, v := range this.Skill {
		skill = append(skill, v)
	}

	for _, v := range this.Time {
		time = append(time, v)
	}

	return modelA.Caddy{
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
