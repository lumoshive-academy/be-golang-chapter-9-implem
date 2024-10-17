package model

type Account struct {
	Name  string
	Email string
	Phone string
}

func NewAccount(name string, email string, phone string) Account {
	return Account{name, email, phone}
}
