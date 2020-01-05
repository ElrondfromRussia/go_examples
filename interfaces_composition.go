package main

import (
	"fmt"
	"strconv"
)

type Payer interface {
	Pay(int) error
}
type Ringer interface {
	Ring(string) error
}
type NFCPhone interface {
	Payer
	Ringer
}

//------------------
type Wallet struct {
	Cash int
	Name string
}
func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}
func (w *Wallet) Ring(nam string) error {
	fmt.Println(nam)
	return nil
}

//-----------------
func (w *Wallet) String() string {
	return "Кошелёк в котором " + strconv.Itoa(w.Cash) + " денег"
}

//---------------
func Buy(in interface{}) {
	var p NFCPhone
	var ok bool
	if p, ok = in.(NFCPhone); !ok {
		fmt.Printf("%T не не является платежным средством\n\n", in)
		return
	}
	err := p.Pay(10)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

//----------------
func main() {
	myWallet := &Wallet{Cash: 100}
	//fmt.Printf("Raw payment : %#v\n", myWallet)
	//fmt.Printf("Способ оплаты: %s\n", myWallet)
	Buy(myWallet)
	Buy([]int{1, 2, 3})
	Buy(3.14)
	return
}


