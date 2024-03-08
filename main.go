package main

type Seller struct {
	ID   int
	Name string
}

type Store struct {
	ID       int
	Name     string
	Branches []*Branch
}

type Branch struct {
	ID      int
	Name    string
	Sellers []*Seller
}

type SellerInvitation struct {
	ID             int
	SenderID       int
	ReceiverID     int
	InvitationLink string
	InvitationDate string
}

type Supplier struct {
	ID       int
	Name     string
	Products []string
}

type PurchaseInvoice struct {
	ID           int
	SupplierID   int
	BranchID     int
	Items        []string
	Prepayment   float64
	PaymentCheck *PaymentCheck
	Debt         float64
}

type PaymentCheck struct {
	ID            int
	BankName      string
	AccountNumber string
	CheckAmount   float64
	CheckDate     string
}

type Inventory struct {
	ProductID int
	Quantity  int
	Price     float64
	Location  string
}

type Customer struct {
	ID        int
	Name      string
	Purchases []*PurchaseInvoice
	Payments  []*Payment
}

type Payment struct {
	ID       int
	Amount   float64
	Date     string
	Comments string
}
