package modeld

type Payment struct {
	BaseBSONModel
	TransactionID      int    `json:"TransactionId"`
	Amount             int    `json:"Amount"`
	OrderNo            string `json:"OrderNo"`
	CustomerID         string `json:"CustomerId"`
	BankCode           string `json:"BankCode"`
	PaymentDate        string `json:"PaymentDate"`
	PaymentStatus      int    `json:"PaymentStatus"`
	BankRefCode        string `json:"BankRefCode"`
	CurrentDate        string `json:"CurrentDate"`
	CurrentTime        string `json:"CurrentTime"`
	PaymentDescription string `json:"PaymentDescription"`
	CreditCardToken    string `json:"CreditCardToken"`
	Currency           string `json:"Currency"`
	CustomerName       string `json:"CustomerName"`
	CheckSum           string `json:"CheckSum"`
}
