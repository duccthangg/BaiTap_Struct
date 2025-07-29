package main

import "fmt"

type Book struct {
	title  string
	pages  int
	author Author
}

// Constructor
func NewBook(title string, pages int, author Author) Book {
	return Book{title: title, pages: pages, author: author}
}

// Getters
func (b Book) Title() string {
	return b.title
}

func (b Book) Pages() int {
	return b.pages
}

func (b Book) Author() Author {
	return b.author
}

// Method hiển thị thông tin
func (b Book) DisplayInfo() {
	defer fmt.Println("Kết thúc hiển thị thông tin sách")

	fmt.Println("Thông tin sách:")
	fmt.Println("Tiêu đề:", b.Title())
	fmt.Println("Số trang:", b.Pages())
	fmt.Println("Tác giả:", b.Author().Name())
	fmt.Println("Email:", b.Author().Email())
}
