package model

import "errors"

type Saldo struct {
	Account Account
	Saldo   int
}

func NewSaldo(account Account) Saldo {
	return Saldo{account, 0}
}

func (saldo *Saldo) DebitSaldo(nominal int) error {
	if nominal <= 0 {
		return errors.New("nominal tidak boleh kosong")
	}
	saldo.Saldo += nominal

	return nil
}

func (saldo *Saldo) CreditSaldo(nominal int) error {
	if nominal <= 0 {
		return errors.New("nominal tidak boleh kosong")
	}
	saldo.Saldo -= nominal
	return nil
}

func NewSaldoSlice(saldo ...Saldo) []Saldo {
	slice_saldo := []Saldo{}

	slice_saldo = append(slice_saldo, saldo...)
	return slice_saldo
}
