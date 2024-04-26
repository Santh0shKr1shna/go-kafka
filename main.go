package main

import (
	"docker/kafka"
	"fmt"
)

type Booking struct {
	Id		int
	Name	string
	Email	string
	Cost	int
}

func NewBooking(id int, name, email string, cost int) *Booking {
	return &Booking{
		Id: id,
		Name: name,
		Email: email,
		Cost: cost,
	}
}

func (b Booking) Stringfy() string {
	return fmt.Sprintf("{id:%d, name:%s, email:%s, cost:%d}", b.Id, b.Name, b.Email, b.Cost)
}

func main() {
	b := NewBooking(1, "user", "user@mail.com", 500)
	for {
		var inp string
		fmt.Print("\nDo you need to send a message? (Y/N): ")
		fmt.Scan(&inp)

		if inp == "Y" {
			kafka.Produce(b.Stringfy());
		} else {
			break
		}
	}
    fmt.Println("Shutting down producer..")
}