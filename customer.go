package main

type Customer struct {
	name          string
	accountNumber string
}

func NewCustomer(name string, accountNumber string) Customer {
	return Customer{name: name, accountNumber: accountNumber}
}

func (c Customer) Name() string {
	return c.name
}

func (c Customer) AccountNumber() string {
	return c.accountNumber
}
