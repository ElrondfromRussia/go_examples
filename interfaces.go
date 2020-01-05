package main

import (
	"fmt"
)

//------------------
type Payer interface {
	Pay(int) error
}

//------------------
type Wallet struct {
	Cash int
}
func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return fmt.Errorf("Не хватает денег в кошельке")
	}
	w.Cash -= amount
	return nil
}

//-----------------
type MWallet struct {
	Cash1 int
	Cash2 int
}
func (w *MWallet) Pay(amount int) error {
	if w.Cash1 + w.Cash2 < amount {
		return fmt.Errorf("Бабла маловато")
	}
	for amount > 0 {
		if w.Cash1 > 0{
			w.Cash1 -= 1
		}else{
			w.Cash2 -= 1
		}
		amount--
	}
	return nil
}
//---------------

func Buy(p Payer) {
	switch p.(type) {
	case *Wallet:
		fmt.Println(p.(*Wallet).Cash)
		fmt.Println("Оплата наличными?")
	case *MWallet:
		plasticCard, ok := p.(*MWallet)
		if !ok {
			fmt.Println("Не удалось преобразовать к типу *Card")
		}
		fmt.Println("Вставляйте карту,", plasticCard.Cash2)
	default:
		fmt.Println("Что-то новое!")
	}
	err := p.Pay(50)
	if err != nil {
		fmt.Printf("Ошибка при оплате %T: %v\n\n", p, err)
		return
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

//----------------
func main() {
	var myMoney Payer

	myMoney = &Wallet{Cash: 10}
	fmt.Println(myMoney)
	fmt.Println(myMoney.Pay(11))
	Buy(myMoney)

	myMoney = &MWallet{Cash1: 10, Cash2:11}
	fmt.Println(myMoney)
	fmt.Println(myMoney.Pay(5))
	fmt.Println(myMoney)
	Buy(myMoney)
	return
}


