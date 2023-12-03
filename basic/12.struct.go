package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Customer struct {
	Id   int       `json:"Id"`
	Name string    `json:"Name"`
	Dob  time.Time `json:"Dob"`
}

type Transaction struct {
	Id     int       `json:"Id"`
	Amt    int       `json:"Amount"`
	Time   time.Time `json:"Time"`
	Credit bool      `json:"Credit"`
}

type Account struct {
	Id     int           `json:"Id"`
	Holder Customer      `json:"Holder"`
	Txions []Transaction `json:"Transactions"`
}

func main() {
	var c Customer
	c.Id = 1
	c.Name = "Monicka"
	c.Dob = time.Date(1996, time.September, 18, 0, 0, 0, 0, time.Local)
	fmt.Printf("%v", c)
	cust := Customer{2, "hari", time.Date(1996, time.September, 18, 0, 0, 0, 0, time.Local)}
	fmt.Printf("\n %v", cust)

	txion1 := Transaction{1, 500, time.Now(), true}
	txion2 := Transaction{2, 1000, time.Now(), false}
	txion3 := Transaction{3, 1500, time.Now(), true}

	acct := Account{1, cust, []Transaction{txion1, txion2, txion3}}
	fmt.Printf("\n---------------- Account info------------------------\n")
	fmt.Printf("%v : ", acct)
	data, err := json.Marshal(acct)
	if err != nil {
		panic(err)
	}
	fmt.Printf("\n------------ json marshal data ----------------------------\n")
	fmt.Printf("%s : ", data)
}
