package entities

type Customer struct {
	Name      string `csv:"first_name"` // .csv column headers
	LastName  string `csv:"last_name"`
	Email     string `csv:"email"`
	Gender    string `csv:"gender"`
	IPAndress string `csv:"ip_address"`
	Domain    string
	NumberOfCustomers int
}


