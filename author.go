package main

type Author struct {
	name  string
	email string
}

// Constructor
func NewAuthor(name, email string) Author {
	return Author{name: name, email: email}
}

// Getter
func (a Author) Name() string {
	return a.name
}

func (a Author) Email() string {
	return a.email
}
