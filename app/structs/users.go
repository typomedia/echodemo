package structs

type Users struct {
	Users []User `json:"users"`
	Total int    `json:"total"`
	Skip  int    `json:"skip"`
	Limit int    `json:"limit"`
}

type User struct {
	ID        int     `json:"id"`
	FirstName string  `json:"firstName"`
	LastName  string  `json:"lastName"`
	Age       int     `json:"age"`
	Gender    string  `json:"gender"`
	Email     string  `json:"email"`
	BirthDate string  `json:"birthDate"`
	Image     string  `json:"image"`
	Address   Address `json:"address"`
	Company   Company `json:"company"`
}

type Address struct {
	Address    string `json:"address"`
	City       string `json:"city"`
	PostalCode string `json:"postalCode"`
	State      string `json:"state"`
}

type Company struct {
	Address    Address `json:"address"`
	Department string  `json:"department"`
	Name       string  `json:"name"`
	Title      string  `json:"title"`
}
