package magazine

type Subscriber struct {
	Name   string
	Rate   float64
	Active bool
	Address
}

type Employee struct {
	Name       string
	Salary     float64
	HomeAddres Address
}

type Address struct {
	State      string
	City       string
	Street     string
	PostalCode string
}
