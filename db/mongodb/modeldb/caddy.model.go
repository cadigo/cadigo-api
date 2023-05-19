package modeldb

import "cadigo-api/app/modelapp"

type Caddy struct {
	BaseBSONModel
	Language     LanguageType `bson:"language"`
	Name         string       `bson:"name"`
	Location     string       `bson:"location,omitempty"`
	Avialability string       `bson:"avialability,omitempty"`
	Skill        []string     `bson:"skill,omitempty"`
	Start        int          `bson:"start,omitempty"`
	Desctiption  string       `bson:"desctiption,omitempty"`
	Time         []string     `bson:"time,omitempty"`
	Tags         []string     `bson:"tags,omitempty"`
	Cost         float64      `bson:"cost"`
}

func (this *Caddy) CaddyDB() modelapp.Caddy {
	var skill []*string
	var time []*string

	for _, v := range this.Skill {
		skill = append(skill, &v)
	}

	for _, v := range this.Time {
		time = append(time, &v)
	}

	return modelapp.Caddy{
		Name:         &this.Name,
		Location:     &this.Location,
		Avialability: &this.Avialability,
		Skill:        skill,
		Start:        &this.Start,
		Description:  &this.Desctiption,
		Time:         time,
		Cost:         &this.Cost,
	}
}
