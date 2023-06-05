package modela

import (
	"cadigo-api/graph/modelgraph"

	"github.com/jinzhu/copier"
)

type CourseGolf struct {
	ID        string   `json:"id" copier:"Id"`
	Name      string   `json:"name" copier:"Name"`
	Images    []string `json:"images,omitempty" copier:"Images"`
	Available int      `json:"available" copier:"Available"`
	Location  string   `json:"location" copier:"Location"`
	Latitude  float64  `json:"latitude" copier:"Latitude"`
	Longitude float64  `json:"longitude" copier:"Longitude"`
	IsActive  bool     `json:"isActive" copier:"IsActive"`
}

func (this CourseGolf) Parse(graph modelgraph.CourseGolfInput) error {
	return copier.Copy(&this, &graph)
}

func (this CourseGolf) ToGraph() modelgraph.CourseGolf {
	g := modelgraph.CourseGolf{}
	copier.Copy(&g, &this)
	this.ID = g.ID
	return g
}
