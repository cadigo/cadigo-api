package modela

import (
	"cadigo-api/graph/modelgraph"

	"github.com/jinzhu/copier"
)

type Customer struct {
	ID     string    `json:"id,omitempty" copier:"Id"`
	UserID *string   `json:"userID,omitempty" copier:"UserID"`
	Name   *string   `json:"name,omitempty" copier:"Name"`
	Images []*string `json:"images,omitempty" copier:"Images"`
}

func (this Customer) Parse(graph modelgraph.CustomerInput) error {
	return copier.Copy(&this, &graph)
}

func (this Customer) ToGraph() modelgraph.Customer {
	g := modelgraph.Customer{}
	copier.Copy(&g, &this)
	return g
}
