package entity

type Person struct {
	FirstName string
	LastName  string
	Email     string
	StateCode *string
}

type Place struct {
	Country string
	City    string
	TelCode int
}
