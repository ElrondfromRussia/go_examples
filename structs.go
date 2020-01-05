package main

import "fmt"

type Person struct {
	Id int
	Name string
	Address string
}

type Account struct {
	Id int
	Cleaner func(string) string
	Owner Person
	Person
}

func (p Person) UpdateName(name string) {
	p.Name = name
}

func (p *Person) SetName(name string) {
	p.Name = name
}

func (acc *Account) SetId(ind int) {
	acc.Owner.Id = ind
}

func (acc *Account) SetFunc(newF func(string) string) {
	acc.Cleaner = newF
}

/////////////////////
type MySlice []int

func (sl *MySlice) Add(val int){
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}
////////////////////

func main() {
	var ac Account = Account{
		Id: 0,
	}
	ac.Owner = Person{
		Id:      2,
		Name:    "Name1",
		Address: "Adr1",
	}
	ac.Person.Id = 5
	ac.Address = "Adr2"
	ac.Name = "Name2"
	ac.Cleaner = func(s string) string {
		return s
	}
	fmt.Println(ac)
	fmt.Println(ac.Cleaner("string from func"))

	ac.Person.SetName("Name22")
	fmt.Println(ac)

	ac.Person.UpdateName("Name222")
	fmt.Println(ac)

	ac.Owner.SetName("Name11")
	fmt.Println(ac)
	
	(&ac).SetFunc(func(s string) string {
		return "AAAAA " + s
	})
	
	ac.SetId(6)

	fmt.Println(ac)
	fmt.Println(ac.Cleaner("string from func"))

	fmt.Println("Here i am")

	fmt.Println("################################")

	sl := MySlice([]int{1, 2})
	fmt.Println(sl.Count(), sl)
	sl.Add(5)
	fmt.Println(sl.Count(), sl)

	return
}
