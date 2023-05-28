// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package modelgraph

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type Booking struct {
	ID           string      `json:"id" copier:"Id"`
	Reference    string      `json:"reference" copier:"Reference"`
	TimeStart    time.Time   `json:"timeStart" copier:"TimeStart"`
	TimeEnd      time.Time   `json:"timeEnd" copier:"TimeEnd"`
	CustomerID   string      `json:"customerID" copier:"CustomerID"`
	Customer     *Customer   `json:"customer,omitempty" copier:"Customer"`
	CourseGolfID string      `json:"courseGolfID" copier:"CourseGolfID"`
	CourseGolf   *CourseGolf `json:"courseGolf,omitempty" copier:"CourseGolf"`
	CaddyID      string      `json:"caddyID" copier:"CaddyID"`
	Caddy        *Caddy      `json:"caddy,omitempty" copier:"Caddy"`
	TotalNet     *float64    `json:"totalNet,omitempty" copier:"TotalNet"`
}

type BookingData struct {
	Data       []*Booking      `json:"data" copier:"Data"`
	Pagination *PaginationType `json:"pagination" copier:"Pagination"`
}

type BookingInput struct {
	Reference    string       `json:"reference" copier:"Reference"`
	TimeStart    time.Time    `json:"timeStart" copier:"TimeStart"`
	TimeEnd      time.Time    `json:"timeEnd" copier:"TimeEnd"`
	CustomerID   string       `json:"customerID" copier:"CustomerID"`
	CourseGolfID string       `json:"courseGolfID" copier:"CourseGolfID"`
	CaddyID      string       `json:"caddyID" copier:"CaddyID"`
	Remark       *string      `json:"remark,omitempty" copier:"Remark"`
	Language     LanguageEnum `json:"language" copier:"Language"`
}

type BookingsInput struct {
	Language   LanguageEnum     `json:"language" copier:"Language"`
	Pagination *PaginationInput `json:"pagination" copier:"Pagination"`
	Status     *string          `json:"status,omitempty" copier:"Status"`
}

type Caddy struct {
	ID           *string   `json:"id,omitempty" copier:"Id"`
	Name         *string   `json:"name,omitempty" copier:"Name"`
	Location     *string   `json:"location,omitempty" copier:"Location"`
	Avialability *string   `json:"avialability,omitempty" copier:"Avialability"`
	Skill        []*string `json:"skill,omitempty" copier:"Skill"`
	Start        *int      `json:"start,omitempty" copier:"Start"`
	Description  *string   `json:"description,omitempty" copier:"Description"`
	Time         []*string `json:"time,omitempty" copier:"Time"`
	Cost         *float64  `json:"cost,omitempty" copier:"Cost"`
	Images       []*string `json:"images,omitempty" copier:"Images"`
}

type CaddyData struct {
	Data       []*Caddy        `json:"data" copier:"Data"`
	Pagination *PaginationType `json:"pagination" copier:"Pagination"`
}

type CaddyInput struct {
	ID           *string      `json:"id,omitempty" copier:"Id"`
	Name         string       `json:"name" copier:"Name"`
	Location     string       `json:"location" copier:"Location"`
	Avialability string       `json:"avialability" copier:"Avialability"`
	Skill        []string     `json:"skill,omitempty" copier:"Skill"`
	Start        int          `json:"start" copier:"Start"`
	Description  *string      `json:"description,omitempty" copier:"Description"`
	Time         []string     `json:"time,omitempty" copier:"Time"`
	Cost         float64      `json:"cost" copier:"Cost"`
	Images       []string     `json:"images,omitempty" copier:"Images"`
	Language     LanguageEnum `json:"language" copier:"Language"`
	IsActive     bool         `json:"isActive" copier:"IsActive"`
}

type ChatInput struct {
	RoomID        string `json:"roomId" copier:"RoomId"`
	CurrentUserID string `json:"currentUserId" copier:"CurrentUserId"`
}

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

type CourseGolfData struct {
	Data       []*CourseGolf   `json:"data" copier:"Data"`
	Pagination *PaginationType `json:"pagination" copier:"Pagination"`
}

type CourseGolfInput struct {
	ID        *string  `json:"id,omitempty" copier:"Id"`
	Name      string   `json:"name" copier:"Name"`
	Images    []string `json:"images,omitempty" copier:"Images"`
	Available int      `json:"available" copier:"Available"`
	Location  string   `json:"location" copier:"Location"`
	Latitude  float64  `json:"latitude" copier:"Latitude"`
	Longitude float64  `json:"longitude" copier:"Longitude"`
	IsActive  bool     `json:"isActive" copier:"IsActive"`
}

type Customer struct {
	ID     *string   `json:"id,omitempty" copier:"Id"`
	UserID *string   `json:"userID,omitempty" copier:"UserID"`
	Name   *string   `json:"name,omitempty" copier:"Name"`
	Images []*string `json:"images,omitempty" copier:"Images"`
}

type CustomerInput struct {
	ID     *string   `json:"id,omitempty" copier:"Id"`
	UserID string    `json:"userID" copier:"UserID"`
	Name   string    `json:"name" copier:"Name"`
	Images []*string `json:"images,omitempty" copier:"Images"`
}

type GetBookingInput struct {
	BookingReference *string      `json:"bookingReference,omitempty" copier:"BookingReference"`
	Language         LanguageEnum `json:"language" copier:"Language"`
}

type GetCaddyInput struct {
	Language LanguageEnum `json:"language" copier:"Language"`
	ID       *string      `json:"id,omitempty" copier:"Id"`
}

type GetCaddysInput struct {
	Language   LanguageEnum     `json:"language" copier:"Language"`
	Pagination *PaginationInput `json:"pagination" copier:"Pagination"`
}

type GetCourseGolfInput struct {
	Language LanguageEnum `json:"language" copier:"Language"`
	ID       *string      `json:"id,omitempty" copier:"Id"`
}

type GetCourseGolfsInput struct {
	Language   LanguageEnum     `json:"language" copier:"Language"`
	Pagination *PaginationInput `json:"pagination" copier:"Pagination"`
}

type GetCustomerInput struct {
	Language LanguageEnum `json:"language" copier:"Language"`
	ID       *string      `json:"id,omitempty" copier:"Id"`
}

type GetMessagesInput struct {
	ToUserID   string `json:"toUserId" copier:"ToUserId"`
	FromUserID string `json:"fromUserId" copier:"FromUserId"`
}

type GetOnlineInput struct {
	ToUserID []string `json:"toUserId" copier:"ToUserId"`
}

type GetPaymentInput struct {
	Language LanguageEnum `json:"language" copier:"Language"`
	ID       *string      `json:"id,omitempty" copier:"Id"`
}

type GetUserInput struct {
	BookingReference *string      `json:"bookingReference,omitempty" copier:"BookingReference"`
	Language         LanguageEnum `json:"language" copier:"Language"`
}

type Message struct {
	ToUserID   string    `json:"toUserId" copier:"ToUserId"`
	FromUserID string    `json:"fromUserId" copier:"FromUserId"`
	Message    string    `json:"message" copier:"Message"`
	CreatedAt  time.Time `json:"createdAt" copier:"CreatedAt"`
	RoomID     string    `json:"roomId" copier:"RoomId"`
}

type Online struct {
	UserID     string    `json:"userId" copier:"UserId"`
	UserName   string    `json:"userName" copier:"UserName"`
	LastOnline time.Time `json:"lastOnline" copier:"LastOnline"`
}

type OnlineInput struct {
	CurrentUserID string `json:"currentUserId" copier:"CurrentUserId"`
}

type PaginationInput struct {
	Page     int           `json:"page" copier:"Page"`
	Limit    int           `json:"limit" copier:"Limit"`
	OrderBy  *string       `json:"orderBy,omitempty" copier:"OrderBy"`
	Asc      *bool         `json:"asc,omitempty" copier:"Asc"`
	Leyword  []*string     `json:"leyword,omitempty" copier:"Leyword"`
	Language *LanguageEnum `json:"language,omitempty" copier:"Language"`
}

type PaginationType struct {
	Page  int `json:"page" copier:"Page"`
	Limit int `json:"limit" copier:"Limit"`
	Total int `json:"total" copier:"Total"`
}

type Payment struct {
	ID                 *string `json:"id,omitempty" copier:"Id"`
	TransactionID      *int    `json:"transactionId,omitempty" copier:"TransactionId"`
	Amount             *int    `json:"amount,omitempty" copier:"Amount"`
	OrderNo            *string `json:"orderNo,omitempty" copier:"OrderNo"`
	CustomerID         *string `json:"customerId,omitempty" copier:"CustomerId"`
	BankCode           *string `json:"bankCode,omitempty" copier:"BankCode"`
	PaymentDate        *string `json:"paymentDate,omitempty" copier:"PaymentDate"`
	PaymentStatus      *int    `json:"paymentStatus,omitempty" copier:"PaymentStatus"`
	BankRefCode        *string `json:"bankRefCode,omitempty" copier:"BankRefCode"`
	CurrentDate        *string `json:"currentDate,omitempty" copier:"CurrentDate"`
	CurrentTime        *string `json:"currentTime,omitempty" copier:"CurrentTime"`
	PaymentDescription *string `json:"paymentDescription,omitempty" copier:"PaymentDescription"`
	CreditCardToken    *string `json:"creditCardToken,omitempty" copier:"CreditCardToken"`
	Currency           *string `json:"currency,omitempty" copier:"Currency"`
	CustomerName       *string `json:"customerName,omitempty" copier:"CustomerName"`
	CheckSum           *string `json:"checkSum,omitempty" copier:"CheckSum"`
}

type PaymentInput struct {
	ID                 *string `json:"id,omitempty" copier:"Id"`
	TransactionID      int     `json:"transactionId" copier:"TransactionId"`
	Amount             int     `json:"amount" copier:"Amount"`
	OrderNo            string  `json:"orderNo" copier:"OrderNo"`
	CustomerID         string  `json:"customerId" copier:"CustomerId"`
	BankCode           string  `json:"bankCode" copier:"BankCode"`
	PaymentDate        string  `json:"paymentDate" copier:"PaymentDate"`
	PaymentStatus      int     `json:"paymentStatus" copier:"PaymentStatus"`
	BankRefCode        string  `json:"bankRefCode" copier:"BankRefCode"`
	CurrentDate        string  `json:"currentDate" copier:"CurrentDate"`
	CurrentTime        string  `json:"currentTime" copier:"CurrentTime"`
	PaymentDescription string  `json:"paymentDescription" copier:"PaymentDescription"`
	CreditCardToken    string  `json:"creditCardToken" copier:"CreditCardToken"`
	Currency           string  `json:"currency" copier:"Currency"`
	CustomerName       string  `json:"customerName" copier:"CustomerName"`
	CheckSum           string  `json:"checkSum" copier:"CheckSum"`
}

type PostMessageInput struct {
	ToUserID   string  `json:"toUserId" copier:"ToUserId"`
	FromUserID string  `json:"fromUserId" copier:"FromUserId"`
	Message    string  `json:"message" copier:"Message"`
	RoomID     *string `json:"roomId,omitempty" copier:"RoomId"`
}

type User struct {
	ID        string `json:"id" copier:"Id"`
	Reference string `json:"reference" copier:"Reference"`
}

type UserInput struct {
	Reference string `json:"reference" copier:"Reference"`
}

type LanguageEnum string

const (
	LanguageEnumTh LanguageEnum = "TH"
	LanguageEnumEn LanguageEnum = "EN"
)

var AllLanguageEnum = []LanguageEnum{
	LanguageEnumTh,
	LanguageEnumEn,
}

func (e LanguageEnum) IsValid() bool {
	switch e {
	case LanguageEnumTh, LanguageEnumEn:
		return true
	}
	return false
}

func (e LanguageEnum) String() string {
	return string(e)
}

func (e *LanguageEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LanguageEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LanguageEnum", str)
	}
	return nil
}

func (e LanguageEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
