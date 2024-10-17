package main

import (
	"be-golang-chapter-09-implem/model"
	"fmt"
)

// buatkan sistem keuangan untuk mengelola data keuangan seseorang
// sistem dapat menambahkan saldo
// sistem dapat mengurangi saldo
// sistem dapat membuat akun baru
// sistem dapat mencetak akun berhasil di tambahkan, akun berhasil menambahkan saldo, akun berhasil mengurangi saldo

// ketentuan baku
// sistem memiliki 2 object account dan saldo
// buatkan 2 object tersebut di file terpisah di satu package
// buatkan sistem validasi error saat manambahkan akun , menambah saldo dan mengurangi saldo

// contoh output :
// akun berhasil di tambahkan [{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 } }{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 } }]
// saldo berhasil di tambahkan [{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 } }{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 } }]
// saldo berhasil di dikurangi [{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 } }{nama: lumo , email:lumos@email.com, saldo: { saldo : 0 }]

func main() {

	// add account
	account_1 := model.NewAccount("lumoshive", "lumo@email.com", "62085385305350")
	saldo_acccount_1 := model.NewSaldo(account_1)

	account_2 := model.NewAccount("lumo", "lumo_2@email.com", "62088430430434")
	saldo_acccount_2 := model.NewSaldo(account_2)

	// show all account
	allAccountSaldo := model.NewSaldoSlice(saldo_acccount_1, saldo_acccount_2)

	fmt.Printf("Account berhasil di tambahkan: %+v\n", allAccountSaldo)

	// debit saldo account 1
	if err := allAccountSaldo[0].DebitSaldo(1000); err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Printf("Saldo akun 1 berhasil di tambahkan: %+v\n", allAccountSaldo)

	//  credit saldo account 1
	if err := allAccountSaldo[0].CreditSaldo(500); err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Printf("Saldo akun 1 berhasil di kurangkan: %+v\n", allAccountSaldo)

	// debit saldo account 2
	if err := allAccountSaldo[1].DebitSaldo(1000); err != nil {
		fmt.Println("Error :", err)
		return
	}
	fmt.Printf("Saldo akun 2 berhasil di tambahkan: %+v\n", allAccountSaldo)
}
