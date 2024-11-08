package order

import (
	"goOnGo/internal/orders/model/domain/address"
	"goOnGo/internal/orders/model/domain/customer"
	"time"
)

type Id uint64
type Number uint64

type Order struct {
	id              Id
	number          Number
	customer        *customer.Customer
	shippingAddress *address.Address
	createDate      time.Time
	updateDate      time.Time
	status          Status
}

func New(number Number, customer *customer.Customer, shippingAddress *address.Address) *Order {
	return &Order{
		number:          number,
		customer:        customer,
		shippingAddress: shippingAddress,
		createDate:      time.Now(),
		updateDate:      time.Now(),
		status:          StatusCreated,
	}
}

func Restore(
	id Id,
	number Number,
	customer *customer.Customer,
	shippingAddress *address.Address,
	createDate time.Time,
	updateDate time.Time,
	status Status,
) *Order {
	return &Order{
		id:              id,
		number:          number,
		customer:        customer,
		shippingAddress: shippingAddress,
		createDate:      createDate,
		updateDate:      updateDate,
		status:          status,
	}
}

func (order *Order) Id() Id {
	return order.id
}

func (order *Order) Number() Number {
	return order.number
}

func (order *Order) Customer() *customer.Customer {
	return order.customer
}

func (order *Order) ShippingAddress() *address.Address {
	return order.shippingAddress
}

func (order *Order) CreateDate() time.Time {
	return order.createDate
}

func (order *Order) UpdateDate() time.Time {
	return order.updateDate
}

func (order *Order) Status() Status {
	return order.status
}

func (order *Order) ChangeShippingAddress(shippingAddress *address.Address) (*Order, error) {
	if order.status.IsImmutable() {
		return nil, ErrOrderIsImmutable
	}

	order.shippingAddress = shippingAddress

	return order, nil
}

func (order *Order) TransitionTo(status Status) (*Order, error) {
	if !order.status.CanTransitionTo(status) {
		return nil, ErrTransitionNotAllowed
	}

	order.status = status

	return order, nil
}
