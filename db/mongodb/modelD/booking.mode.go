package modelD

type Booking struct {
	BaseBSONModel
	BookingNumber string     `bson:"booking_number"`
	Caddy         Caddy      `bson:"caddy"`
	Payment       Payment    `bson:"payment"`
	Users         Users      `bson:"users"`
	CourseGolf    CourseGolf `bson:"course_golf"`
}
