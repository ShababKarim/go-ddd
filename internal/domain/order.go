package domain

import (
	"errors"
	"github.com/google/uuid"
)

// There should be a pkg to group all domain objects,
// use cases, adapters, etc. by domain instead of leaving in the root

// Domain Objects

type Drink interface {
	GetName() string
	GetPrice() int
}

type AddOn struct {
	name  string
	price int
	drink Drink
}

func (a *AddOn) GetName() string {
	return a.name + " + " + a.drink.GetName()
}

func (a *AddOn) GetPrice() int {
	return a.price + a.drink.GetPrice()
}

type OrderStatus string

const (
	PaymentPending  OrderStatus = "payment pending"
	PaymentReceived OrderStatus = "payment received"
	PaymentDeclined OrderStatus = "payment declined"
	Pending         OrderStatus = "pending"
	Completed       OrderStatus = "completed"
)

var OrderStatuses = [...]OrderStatus{
	PaymentPending,
	PaymentReceived,
	PaymentDeclined,
	Pending,
	Completed,
}

// Commands

type PlaceOrder struct {
	customer string
	drinks   []Drink
}

type UpdateOrderStatus struct {
	NewOrderStatus OrderStatus
}

// Order is an Aggregate
type Order struct {
	ID          uuid.UUID
	customer    string
	drinks      []Drink
	orderStatus OrderStatus
}

func NewOrder(command PlaceOrder) (*Order, error) {
	if len(command.customer) == 0 {
		return nil, errors.New("cannot initialize order without customer name")
	} else if command.drinks == nil || len(command.drinks) == 0 {
		return nil, errors.New("cannot initialize order without drinks")
	}

	return &Order{
		ID:          uuid.New(),
		customer:    command.customer,
		drinks:      command.drinks,
		orderStatus: PaymentPending,
	}, nil
}

func (o *Order) UpdateOrderStatus(command UpdateOrderStatus) (*Order, error) {
	for _, status := range OrderStatuses {
		if status == command.NewOrderStatus {
			o.orderStatus = command.NewOrderStatus

			return o, nil
		}
	}

	return o, errors.New("invalid order status")
}

func (o *Order) CalculateTotalPrice() int {
	total := 0
	for _, drink := range o.drinks {
		total += drink.GetPrice()
	}

	return total
}
