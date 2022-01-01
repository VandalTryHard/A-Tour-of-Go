package main

import (
	"fmt"

	"../magazine"
)

func main() {
	subscriber := magazine.Subscriber{Rate: 4.99}
	subscriber.Street = "123 Oak St"
	fmt.Printf("Name subscriber: %s; Rate subscriber %.2f; Active subscriber %v; HomeAddres: %v", subscriber.Name,
		subscriber.Rate, subscriber.Active, subscriber.Address.Street)

	employee := magazine.Employee{Name: "Val", Salary: 499}
	fmt.Printf("\nName employee: %s; Salary employee %.2f.", employee.Name, employee.Salary)

}

// func printInfo(s *subscriber) {
// 	fmt.Println("Name:", s.name)
// 	fmt.Println("Monthly rate:", s.rate)
// 	fmt.Println("Active?", s.active)
// }

// func defaultSubscriber(name string) *subscriber {
// 	var s subscriber
// 	s.name = name
// 	s.rate = 5.99
// 	s.active = true
// 	return &s
// }

// func applyDiscount(s *subscriber) {
// 	s.rate = 4.99
// }

// func main() {
// 	var s subscriber
// 	fmt.Println(&s.active)
// 	applyDiscount(&s)
// 	fmt.Println(s.rate)
// }
