package main

import "fmt"

type BankAccount struct {
	customer Customer
	balance  float64
}

func NewBankAccount(customer Customer) BankAccount {
	return BankAccount{customer: customer, balance: 500.00}
}

func (b BankAccount) Balance() float64 {
	return b.balance
}

func (b BankAccount) Custom() Customer {
	return b.customer
}

func (b *BankAccount) Withdraw(amount float64) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r) // in ra lỗi từ panic
		}
	}()

	fmt.Printf("Thực hiện rút tiền: %.2f\n", amount)
	if amount > b.balance {
		panic("Lỗi xảy ra: Số dư không đủ")
	}
	b.balance -= amount
	fmt.Printf("Rút thành công. Số dư còn lại: %.2f\n", b.balance)
}
