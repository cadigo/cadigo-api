package modela

import (
	"cadigo-api/graph/modelgraph"
	"time"

	"github.com/jinzhu/copier"
)

type Booking struct {
	ID           string    `json:"id" copier:"Id"`
	Reference    string    `json:"reference" copier:"Reference"`
	TimeStart    time.Time `json:"timeStart" copier:"TimeStart"`
	TimeEnd      time.Time `json:"timeEnd" copier:"TimeEnd"`
	CustomerID   string    `json:"customerID" copier:"CustomerID"`
	CourseGolfID string    `json:"courseGolfID" copier:"CourseGolfID"`
	CaddyID      string    `json:"caddyID" copier:"CaddyID"`
	TotalNet     *float64  `json:"totalNet,omitempty" copier:"TotalNet"`
}

func (this Booking) Parse(graph modelgraph.BookingInput) error {
	return copier.Copy(&this, &graph)
}

func (this Booking) ToGraph() modelgraph.Booking {
	g := modelgraph.Booking{}
	copier.Copy(&g, &this)
	return g
}
