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

/**
go run 12.struct.go
{1 Monicka 1996-09-18 00:00:00 +0530 IST}
 {2 hari 1996-09-18 00:00:00 +0530 IST}
---------------- Account info------------------------
{1 {2 hari 1996-09-18 00:00:00 +0530 IST} [{1 500 2023-12-03 08:54:29.7989566 +0530 IST m=+0.008266601 true} {2 1000 2023-12-03 08:54:29.7989566 +0530 IST m=+0.008266601 false} {3 1500 2023-12-03 08:54:29.7989566 +0530 IST m=+0.008266601 true}]} :
------------ json marshal data ----------------------------
{"Id":1,"Holder":{"Id":2,"Name":"hari","Dob":"1996-09-18T00:00:00+05:30"},"Transactions":[{"Id":1,"Amount":500,"Time":"2023-12-03T08:54:29.7989566+05:30","Credit":true},{"Id":2,"Amount":1000,"Time":"2023-12-03T08:54:29.7989566+05:30","Credit":false},{"Id":3,"Amount":1500,"Time":"2023-12-03T08:54:29.7989566+05:30","Credit":true}]} :
*/
