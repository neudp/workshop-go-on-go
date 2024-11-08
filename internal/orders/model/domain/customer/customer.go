package customer

import "goOnGo/internal/orders/model/domain/address"

type Id uint64

type Customer struct {
	id              Id
	billingAddress  *address.Address
	shippingAddress *address.Address
}

func New(billingAddress *address.Address, shippingAddress *address.Address) *Customer {
	return &Customer{
		billingAddress:  billingAddress,
		shippingAddress: shippingAddress,
	}
}

func Restore(id Id, name string, billingAddress *address.Address, shippingAddress *address.Address) *Customer {
	return &Customer{
		id:              id,
		billingAddress:  billingAddress,
		shippingAddress: shippingAddress,
	}
}

func (customer *Customer) Id() Id {
	return customer.id
}

func (customer *Customer) BillingAddress() *address.Address {
	return customer.billingAddress
}

func (customer *Customer) ShippingAddress() *address.Address {
	return customer.shippingAddress
}

func (customer *Customer) ChangeBillingAddress(billingAddress *address.Address) *Customer {
	customer.billingAddress = billingAddress

	return customer
}

func (customer *Customer) ChangeShippingAddress(shippingAddress *address.Address) *Customer {
	customer.shippingAddress = shippingAddress

	return customer
}
