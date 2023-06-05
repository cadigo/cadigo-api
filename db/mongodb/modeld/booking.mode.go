package modeld

import (
	"cadigo-api/app/modela"
	"time"

	"github.com/jinzhu/copier"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Booking struct {
	BaseBSONModel `bson:",inline"`
	TimeStart     time.Time          `bson:"timeStart"`
	TimeEnd       time.Time          `bson:"timeEnd"`
	Reference     string             `bson:"reference"`
	Caddy         primitive.ObjectID `bson:"caddyID"`
	Payment       primitive.ObjectID `bson:"paymentID"`
	Customer      primitive.ObjectID `bson:"customerID"`
	CourseGolf    primitive.ObjectID `bson:"courseGolfID"`
	TotalNet      float64            `bson:"totalNet"`
	Remark        string             `bson:"remark"`
	Status        string             `bson:"status"`
}

func (this Booking) Init() Booking {
	this.BaseBSONModel.CreatedAt = time.Now()
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this Booking) SetBooking(b modela.Booking) (*Booking, error) {
	this.BaseBSONModel.UpdatedAt = time.Now()

	err := copier.Copy(&this, &b)

	caddyID, err := primitive.ObjectIDFromHex(b.CaddyID)
	if err != nil {
		return nil, err
	}

	courseGolfID, err := primitive.ObjectIDFromHex(b.CourseGolfID)
	if err != nil {
		return nil, err
	}

	customerID, err := primitive.ObjectIDFromHex(b.CustomerID)
	if err != nil {
		return nil, err
	}

	if b.PaymentID != "" {
		paymentID, err := primitive.ObjectIDFromHex(b.PaymentID)
		if err != nil {
			return nil, err
		}

		this.Payment = paymentID
	}
	this.Caddy = caddyID
	this.CourseGolf = courseGolfID
	this.Customer = customerID

	if err != nil {
		return nil, err
	}

	return &this, nil
}

func (this Booking) Update() Booking {
	this.BaseBSONModel.UpdatedAt = time.Now()

	return this
}

func (this Booking) ToBooking() modela.Booking {
	return modela.Booking{
		ID:           this.RawID.Hex(),
		TimeStart:    this.TimeStart,
		TimeEnd:      this.TimeEnd,
		Reference:    this.Reference,
		CustomerID:   this.Customer.Hex(),
		CourseGolfID: string(this.CourseGolf.Hex()),
		CaddyID:      string(this.Caddy.Hex()),
		PaymentID:    string(this.Payment.Hex()),
		TotalNet:     &this.TotalNet,
	}
}
