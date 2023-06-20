package entity

type Order struct {
	ID            string
	Asset         *Asset
	Investor      *Investor
	Shares        int
	Price         float64
	PendingShares int
	OrderType     OrderType
	Status        OrderStatus
	Transactions  []*Transaction
}

type OrderType string
type OrderStatus string

func NewOrder(id string, asset *Asset, investor *Investor, shares int, price float64, orderType OrderType) *Order {
	return &Order{
		ID:            id,
		Asset:         asset,
		Investor:      investor,
		Shares:        shares,
		Price:         price,
		PendingShares: shares,
		OrderType:     orderType,
		Status:        "OPEN",
		Transactions:  []*Transaction{},
	}
}
