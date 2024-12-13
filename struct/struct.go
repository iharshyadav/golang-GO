package main

import (
	"fmt"
	"time"
)

// struct in struct define
type customer struct {
	place string
	pincode string
	area string
}

type order struct {
	name      string
	id        string
	amount    float32
	status    string
	createdAt time.Time
	customer
}

func newConst (name string , id string , amount float32 , status string) *order {
 
	myOrder := order {
		id : id,
		name: name,
		amount: amount,
		status: status,
	}

	return &myOrder
}

func (o *order) myfunc (amount float32) {
	 o.amount = amount
}

func main() {

	orders := order{
		id : "sjncjsnc",
		name: "djnccjdbv",
		amount: 23.12,
		status: "deleverd",
		createdAt: time.Now(),
		customer: customer{
			place: "sjcnj",
			pincode: "sjcnjddnc",
			area: "dkncjdncjd",
		},
	}

	// instatnly invoke struct and assign value...
	profile := struct {
       name string
	   address string
	} {"harsh yadav","skardi greens"}

	fmt.Println(profile)

	orders.myfunc(12.12)
	fmt.Println(newConst("harsh","Dvdefvevfedfv23323",23.00,"successfull"))
	fmt.Println(orders.pincode)

   fmt.Println(orders.amount)

	fmt.Println(orders.pincode)
}