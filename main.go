package main

import "fmt"

func main() {
	author := NewAuthor("Nguyễn Văn A", "nguyenvana@example.com")
	book := NewBook("Lập trình Go cơ bản", 250, author)

	book.DisplayInfo()

	fmt.Println("================")

	customer := Customer{
		name:          "Trần Thị B",
		accountNumber: "123456789",
	}

	// Tạo tài khoản ngân hàng
	account := NewBankAccount(customer)

	// Hiển thị thông tin
	fmt.Println("Khách hàng:", account.Customer().Name())
	fmt.Println("Số tài khoản:", account.Customer().AccountNumber())
	fmt.Printf("Số dư ban đầu: %.2f\n", account.Balance())

	// Rút tiền vượt quá số dư để test panic + recover
	account.Withdraw(500.00)

}
